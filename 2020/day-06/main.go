package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

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
		result = part2(data)
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

func part2(data [][]byte) int {
	return 0
}
