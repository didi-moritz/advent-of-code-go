package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 7, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {

	result := 0

	ts := make([]bool, len(data[0]))

	for i, line := range data {
		if i == 0 {
			for j, c := range line {
				if c == 'S' {
					ts[j] = true
				}
			}
			continue
		}

		for j, c := range line {
			if c == '^' && ts[j] {
				ts[j] = false
				ts[j-1] = true
				ts[j+1] = true
				result++
			}
		}
	}

	return result
}

func part2(data []string) int {

	result := 0

	return result
}
