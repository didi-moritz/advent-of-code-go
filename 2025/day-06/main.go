package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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

	operations := strings.Fields(data[len(data)-1])
	fmt.Println(operations)

	sum := 0
	var add bool
	for i, c := range data[len(data)-1] {
		if string(c) != " " {
			result += sum

			fmt.Println("new sum", sum, result)

			add = string(c) == "+"

			if add {
				sum = 0
			} else {
				sum = 1
			}
		}

		n := ""
		for j := range len(data) - 1 {
			nc := string(data[j][i])
			if nc != " " {
				n += nc
			}
		}

		if n != "" {
			nn, _ := strconv.Atoi(n)

			if add {
				sum += nn
			} else {
				sum *= nn
			}
			fmt.Println(nn, sum)
		}

	}

	result += sum

	return result
}
