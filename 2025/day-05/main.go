package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 5, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type Range struct {
	from, to int
}

func part1(data []string) int {
	var ingredients []int
	var ranges []Range

	rangesDone := false
	for _, line := range data {
		if rangesDone {
			var n int
			n, _ = strconv.Atoi(line)
			ingredients = append(ingredients, n)
		} else {
			if line == "" {
				rangesDone = true
			} else {
				var from, to int
				fmt.Sscanf(line, "%d-%d", &from, &to)
				ranges = append(ranges, Range{from, to})
			}
		}
	}

	result := 0

	for _, ingredient := range ingredients {
		found := false
		for _, r := range ranges {
			if ingredient >= r.from && ingredient <= r.to {
				found = true
				break
			}
		}

		fmt.Println(ingredient, found)

		if found {
			result++
		}
	}

	return result
}

func part2(data []string) int {
	return 0
}
