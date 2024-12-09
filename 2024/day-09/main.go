package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 9, realData))

	var result int64
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int64 {
	var n []int64

	fileIndex := 0
	isFile := true
	for _, line := range data {
		for _, c := range line {
			for i := 0; i < int(c-'0'); i++ {
				var value int
				if isFile {
					value = fileIndex
				} else {
					value = -1
				}
				n = append(n, int64(value))
			}

			if isFile {
				fileIndex++
			}

			isFile = !isFile
		}
	}

	b := int64(len(n))

	sum := int64(0)
	for a := int64(0); a < b; a++ {
		if n[a] == -1 {
			nb := b - 1
			for {
				if n[nb] > -1 || nb == a {
					break
				}
				nb--
			}

			if nb > a {
				n[a] = n[nb]
				n[nb] = -1
			}

			b = nb
		}

		if b <= a {
			break
		}

		if n[a] > -1 {
			sum += n[a] * a
		}
	}

	return sum
}

func part2(data [][]byte) int64 {
	return 0
}
