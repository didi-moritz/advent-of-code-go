package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 9, realData))

	var result int
	if part == 1 {
		result = part1(data, realData)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string, realData bool) int {
	considerCount := 25
	if !realData {
		considerCount = 5
	}

	numbers := make([]int, len(data))
	for i, line := range data {
		numbers[i], _ = strconv.Atoi(line)
	}

	for i := considerCount; i < len(numbers); i++ {
		consideredNumbers := numbers[i-considerCount : i]

		if !check(numbers[i], consideredNumbers) {
			return numbers[i]
		}
	}

	return 0
}

func check(number int, numbers []int) bool {
	for i1, n1 := range numbers[:len(numbers)-1] {
		for _, n2 := range numbers[i1+1:] {
			if n1+n2 == number {
				return true
			}
		}
	}

	return false
}

func part2(data []string) int {
	return 0
}
