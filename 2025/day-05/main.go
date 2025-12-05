package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
	var ranges []Range

	for _, line := range data {
		if line == "" {
			break
		}
		var from, to int
		fmt.Sscanf(line, "%d-%d", &from, &to)
		ranges = append(ranges, Range{from, to})

	}

	result := 0

	for {
		merged := false
		for i, r1 := range ranges {
			for j, r2 := range ranges {
				if i == j {
					continue
				}

				if r1.from >= r2.from && r1.from <= r2.to {
					if r1.to > r2.to {
						r2.to = r1.to
						ranges[j] = r2
					}
					merged = true
				} else if r1.to >= r2.from && r1.to <= r2.to {
					if r1.from < r2.from {
						r2.from = r1.from
						ranges[j] = r2
					}
					merged = true
				}

				if merged {
					ranges = append(ranges[:i], ranges[i+1:]...)
					break
				}
			}

			if merged {
				break
			}
		}

		if !merged {
			break
		}
	}

	for _, r := range ranges {
		result += r.to - r.from + 1
	}

	return result
}
