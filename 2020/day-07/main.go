package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"regexp"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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

type bags struct {
	name  string
	count int
}

func part2(data []string) int {
	r1 := regexp.MustCompile("^(.*) bags contain (.*)\\.$")
	r2 := regexp.MustCompile("(\\d+) ([^,]+) bags?")

	colorContains := make(map[string][]bags)

	for _, line := range data {
		find := r1.FindStringSubmatch(line)
		firstColor := find[1]

		find2 := r2.FindAllStringSubmatch(find[2], -1)
		if len(find2) > 0 {
			for _, f2 := range find2 {
				secondColorCount := int(utils.ParseInt(f2[1]))
				secondColorName := f2[2]
				colorContains[firstColor] = append(colorContains[firstColor], bags{secondColorName, secondColorCount})
			}
		}
	}

	result := calc("shiny gold", colorContains)

	return result
}

func calc(color string, colorContains map[string][]bags) int {
	containingBags := colorContains[color]
	if len(containingBags) == 0 {
		return 0
	}

	sum := 0
	for _, bag := range containingBags {
		sum += bag.count * (calc(bag.name, colorContains) + 1)
	}

	return sum
}
