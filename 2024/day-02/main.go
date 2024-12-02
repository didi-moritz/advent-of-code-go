package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, true)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 2, realData))

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

	for _, line := range data {
		if processLinePart1(line) {
			result++
		}
	}

	return result
}

func processLinePart1(line string) bool {
	words := strings.Split(line, " ")
	var increase bool
	var numbers []int
	for _, word := range words {
		number, _ := strconv.Atoi(word)
		numbers = append(numbers, number)
	}

	if len(numbers) < 2 {
		return false
	}

	increase = numbers[0] < numbers[1]

	for i := 0; i < len(numbers)-1; i++ {
		var diff int
		if increase {
			diff = numbers[i+1] - numbers[i]
		} else {
			diff = numbers[i] - numbers[i+1]
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func part2(data []string) int {
	result := 0

	for _, line := range data {
		if processLinePart2(line) {
			result++
		}
	}

	return result
}

func processLinePart2(line string) bool {
	words := strings.Split(line, " ")
	numbers := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		number, _ := strconv.Atoi(words[i])
		numbers[i] = number
	}

	if len(numbers) < 2 {
		return false
	}

	for i := 0; i < len(numbers); i++ {
		numbersCopy := make([]int, len(numbers))
		copy(numbersCopy, numbers)
		numbersCopy = append(numbersCopy[:i], numbersCopy[i+1:]...)
		if processNumbers(numbersCopy) {
			return true
		}
	}

	return false
}

func processNumbers(numbers []int) bool {
	increase := numbers[0] < numbers[1]

	for i := 0; i < len(numbers)-1; i++ {
		var diff int
		if increase {
			diff = numbers[i+1] - numbers[i]
		} else {
			diff = numbers[i] - numbers[i+1]
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
