package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 19, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	patterns := strings.Split(data[0], ", ")

	towels := data[2:]

	result := 0
	for i, towel := range towels {
		check := checkTower(towel, patterns)
		if check {
			result++
		}
		fmt.Println(i, check)
	}

	return result
}

func checkTower(towelPart string, patterns []string) bool {
	for _, pattern := range patterns {
		if !strings.HasPrefix(towelPart, pattern) {
			continue
		}

		if len(pattern) == len(towelPart) {
			return true
		}

		if checkTower(towelPart[len(pattern):], patterns) {
			return true
		}
	}

	return false
}

func part2(data []string) int {
	return 0
}
