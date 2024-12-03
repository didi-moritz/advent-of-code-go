package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"regexp"
)

func main() {
	part, realData := utils.GetRunConfig(2, true)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 3, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	result := 0
	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")

	for _, line := range data {

		finds := r.FindAll(line, -1)

		for _, find := range finds {
			exp := string(find)
			var a, b int
			fmt.Sscanf(exp, "mul(%d,%d)", &a, &b)
			result += a * b
		}
	}

	return result
}

func part2(data [][]byte) int {
	result := 0
	r := regexp.MustCompile("(mul\\(\\d{1,3},\\d{1,3}\\)|do(n't)?\\(\\))")
	do := true

	for _, line := range data {

		finds := r.FindAll(line, -1)

		for _, find := range finds {
			exp := string(find)

			if exp == "do()" {
				do = true
				continue
			} else if exp == "don't()" {
				do = false
				continue
			}

			if !do {
				continue
			}

			var a, b int
			fmt.Sscanf(exp, "mul(%d,%d)", &a, &b)
			result += a * b
		}
	}

	return result
}
