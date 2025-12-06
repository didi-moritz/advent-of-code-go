package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 6, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {

	result := 0

	var numbers [][]int
	var operations []string

	for i, line := range data {
		fields := strings.Fields(line)

		if i == 0 {
			numbers = make([][]int, len(fields))
		}

		if i == len(data)-1 {
			operations = fields
			break
		}

		for j, field := range fields {
			var n int
			n, _ = strconv.Atoi(field)
			numbers[j] = append(numbers[j], n)
		}
	}

	for i, ns := range numbers {
		sum := 0
		add := operations[i] == "+"
		if !add {
			sum = 1
		}

		for _, n := range ns {
			if add {
				sum += n
			} else {
				sum *= n
			}

		}

		result += sum
	}

	return result
}

func part2(data []string) int {

	result := 0

	return result
}
