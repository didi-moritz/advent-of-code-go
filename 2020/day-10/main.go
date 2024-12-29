package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 10, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	numbers := make([]int, len(data))
	for i, line := range data {
		numbers[i], _ = strconv.Atoi(line)
	}

	sort.Ints(numbers)

	oneDiffCounts := 0
	threeDiffCounts := 0

	current := 0
	for _, number := range numbers {
		diff := number - current
		if diff == 1 {
			oneDiffCounts++
		}

		if diff == 3 {
			threeDiffCounts++
		}

		current = number
	}

	threeDiffCounts++

	fmt.Println(oneDiffCounts, threeDiffCounts)

	return oneDiffCounts * threeDiffCounts
}

func part2(data []string) int {
	return 0
}
