package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"sort"
)

func main() {
	part, realData := utils.GetRunConfig(1, true)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 1, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(data []string) int {
	var lefts []int
	var rights []int
	for _, line := range data {
		var left, right int
		fmt.Sscanf(line, "%d  %d", &left, &right)
		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	result := 0

	for i := 0; i < len(lefts); i++ {
		result += abs(lefts[i] - rights[i])
	}

	return result
}

func part2(data []string) int {
	return 0
}
