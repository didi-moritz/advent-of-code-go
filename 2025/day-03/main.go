package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 3, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func contains(batteries string, number int) bool {
	first := strconv.Itoa(number / 10)[0]
	second := strconv.Itoa(number % 10)[0]

	for i := 0; i < len(batteries)-1; i++ {
		if batteries[i] == first {
			for j := i + 1; j < len(batteries); j++ {
				if batteries[j] == second {
					return true
				}
			}
		}
	}
	return false
}

func part1(data []string) int {
	result := 0

	for _, line := range data {
		fmt.Println(line)
		for i := 100; i > 9; i-- {
			if contains(line, i) {
				fmt.Println(i)
				result += i
				break
			}
		}
	}

	return result
}

func part2(data []string) int {
	result := 0

	return result
}
