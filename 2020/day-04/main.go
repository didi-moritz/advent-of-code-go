package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strings"
)

func main() {
	part, useRealData := utils.GetRunConfig(1, true)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 4, useRealData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	var lines []string

	result := 0

	for _, line := range data {
		if line == "" {
			result += processLines(lines)
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}

	result += processLines(lines)

	return result
}

func processLines(lines []string) int {
	count := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			if !strings.HasPrefix(word, "cid:") {
				count++
			}
		}
	}

	if count == 7 {
		return 1
	} else {
		return 0
	}
}

func part2(data []string) int {
	return 0
}
