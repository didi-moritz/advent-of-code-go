package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"regexp"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 7, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	r1 := regexp.MustCompile("^(.*) bags contain (.*)\\.$")
	r2 := regexp.MustCompile("\\d+ ([^,]+) bags?")

	colorContains := make(map[string][]string)

	for _, line := range data {
		find := r1.FindStringSubmatch(line)
		firstColor := find[1]

		find2 := r2.FindAllStringSubmatch(find[2], -1)
		if len(find2) > 0 {
			for _, f2 := range find2 {
				secondColor := f2[1]
				colorContains[secondColor] = append(colorContains[secondColor], firstColor)
			}
		}
	}

	var found []string
	found = find("shiny gold", found, colorContains)

	return len(found) - 1
}

func find(color string, found []string, colorContains map[string][]string) []string {
	if slices.Contains(found, color) {
		return found
	}

	found = append(found, color)

	isContainedIn := colorContains[color]
	for _, newColor := range isContainedIn {
		found = find(newColor, found, colorContains)
	}

	return found
}

func part2(data []string) int {
	return 0
}
