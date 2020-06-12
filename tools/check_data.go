package tools

import (
	"fmt"
	"os"
)

//check whether 2 partition have the same data
func CheckData(part1, part2 string) {
	part1File, err := os.Open(part1)
	if Check(err) {
		return
	}
	part2File, err := os.Open(part2)
	if Check(err) {
		return
	}
	defer part1File.Close()
	defer part2File.Close()

	_, err = part1File.Seek(7516192768, 0)
	if Check(err) {
		return
	}
	_, err = part2File.Seek(0, 0)
	if Check(err) {
		return
	}
	size := (1024 * 1024)
	buf1 := make([]byte, size)
	buf2 := make([]byte, size)
	cnt := 0
	for {
		_, err = part1File.Read(buf1)
		if Check(err) {
			return
		}
		n2, err := part2File.Read(buf2)
		if Check(err) {
			return
		}
		cnt += n2
		for i := 0; i < n2; i++ {
			if buf1[i] != buf2[i] {
				fmt.Println("wrong data!!!!!  at", cnt)
				return
			}
			if n2 < size {
				fmt.Println("ALL data same!!!!")
				return
			}
		}
	}

}