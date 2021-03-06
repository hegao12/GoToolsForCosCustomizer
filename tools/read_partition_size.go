package tools

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

//ReadPartitionSize reads the size of a partition (unit:sectors of 512 Bytes)
func ReadPartitionSize(disk, partNum string) (int, error) {
	//dump partition table and grep the line
	partName := disk + partNum
	cmd := string("sfdisk --dump ") + disk + " |grep " + partName
	line, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if Check(err, cmd) {
		return -1, err
	}
	if len(line) < 4 { //not find a valid info line
		return -1, errors.New("cannot find partition " + partName)
	}
	size := -1
	ls := strings.Split(string(line), " ")
	mode := 0
	for _, word := range ls {
		switch mode {
		case 0: //looking for size
			if word == "size=" {
				mode = 1
			}
		case 1:
			if len(word) > 0 { //a valid size number has at least 1 digits
				mode = 2
				size, err = strconv.Atoi(word[:len(word)-1]) //a comma at the end
				if Check(err, "cannot covert size sector to int") {
					return 0, err
				}
			}
		default:
			return -1, errors.New("Error: error in looking for partition")
		}
		if mode == 2 {
			break
		}
	}
	if size == -1 {
		return -1, errors.New("Error: error in looking for partition")
	}
	return size, nil
}
