package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 5, realData))

	emptyLine := 0
	for i, line := range data {
		if line == "" {
			emptyLine = i
			break
		}
	}

	rules := data[:emptyLine]
	updates := data[emptyLine+1:]
	fmt.Println(updates)

	var result int
	if part == 1 {
		result = part1(rules, updates)
	} else {
		result = part2(rules, updates)
	}

	fmt.Println(result)
}

func part1(ruleLines []string, updateLines []string) int {
	rules := make(map[int][]int)
	for _, ruleLine := range ruleLines {
		var a, b int
		fmt.Sscanf(ruleLine, "%d|%d", &a, &b)
		rules[a] = append(rules[a], b)
	}

	result := 0

	for _, updateLine := range updateLines {
		check := true
		pages := strings.Split(updateLine, ",")
		pageNumbers := make([]int, len(pages))
		for i, page := range pages {
			pageNumbers[i] = int(utils.ParseInt(page))
		}

		for i, a := range pageNumbers {
			for _, b := range pageNumbers[i+1:] {
				if slices.Contains(rules[b], a) {
					check = false
					break
				}
			}
		}

		if check {
			fmt.Println(updateLine, "foo")
			result += pageNumbers[len(pageNumbers)/2]
		}
	}

	return result
}

func part2(rules []string, updates []string) int {
	return 0
}
