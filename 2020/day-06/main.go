package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2020, 6, realData))

	var blocks [][][]byte
	lastI := 0
	for i, line := range data {
		if string(line) == "" {
			blocks = append(blocks, data[lastI:i])
			lastI = i + 1
		}
	}
	blocks = append(blocks, data[lastI:])

	var result int
	if part == 1 {
		result = part1(blocks)
	} else {
		result = part2(blocks)
	}

	fmt.Println(result)
}

func part1(blocks [][][]byte) int {
	result := 0
	for _, block := range blocks {
		var found []byte
		for _, line := range block {
			for _, c := range line {
				if !slices.Contains(found, c) {
					found = append(found, c)
				}
			}
		}

		result += len(found)
	}

	return result
}

func part2(blocks [][][]byte) int {
	result := 0
	for _, block := range blocks {
		var found []byte

		for _, c := range block[0] {
			found = append(found, c)
		}

		for _, line := range block[1:] {
			for i := len(found) - 1; i >= 0; i-- {
				c := found[i]
				if !slices.Contains(line, c) {
					found = append(found[:i], found[i+1:]...)
				}
			}
		}

		result += len(found)
	}

	return result
}
