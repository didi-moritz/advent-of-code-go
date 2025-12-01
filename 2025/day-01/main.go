package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 1, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	pos := 50

	result := 0

	for _, line := range data {
		var command string
		var steps int
		fmt.Sscanf(line, "%1s%d", &command, &steps)
		if command == "L" {
			pos += steps
		} else {
			pos -= steps
		}

		for pos < 0 {
			pos += 100
		}

		for pos > 99 {
			pos -= 100
		}

		if pos == 0 {
			result++
		}
	}

	return result
}

func part2(data []string) int {
	return 0
}
