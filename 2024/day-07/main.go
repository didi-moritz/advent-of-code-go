package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 7, realData))

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
		parts := strings.Split(line, ": ")

		check := int(utils.ParseInt(parts[0]))

		words := strings.Split(parts[1], " ")
		numbers := make([]int, len(words))
		for i, word := range words {
			numbers[i] = int(utils.ParseInt(word))
		}

		sum := checkRecursively(check, numbers, 1, numbers[0])

		if sum == check {
			result += check
		}
	}

	return result
}

func checkRecursively(check int, numbers []int, pos int, sum int) int {
	if sum > check {
		return -1
	}

	if pos == len(numbers) {
		return sum
	}

	sumAdd := checkRecursively(check, numbers, pos+1, sum+numbers[pos])
	if sumAdd == check {
		return check
	}

	sumMultiplication := checkRecursively(check, numbers, pos+1, sum*numbers[pos])
	if sumMultiplication == check {
		return check
	}

	return -1
}

func part2(data []string) int {
	result := 0
	for _, line := range data {
		parts := strings.Split(line, ": ")

		check := int(utils.ParseInt(parts[0]))

		words := strings.Split(parts[1], " ")
		numbers := make([]int, len(words))
		for i, word := range words {
			numbers[i] = int(utils.ParseInt(word))
		}

		sum := checkRecursively2(check, numbers, 1, numbers[0])

		if sum == check {
			result += check
		}
	}

	return result
}

func checkRecursively2(check int, numbers []int, pos int, sum int) int {
	if sum > check {
		return -1
	}

	if pos == len(numbers) {
		return sum
	}

	sumAdd := checkRecursively2(check, numbers, pos+1, sum+numbers[pos])
	if sumAdd == check {
		return check
	}

	sumMultiplication := checkRecursively2(check, numbers, pos+1, sum*numbers[pos])
	if sumMultiplication == check {
		return check
	}

	concat := int(utils.ParseInt(strconv.Itoa(sum) + strconv.Itoa(numbers[pos])))
	sumConcat := checkRecursively2(check, numbers, pos+1, concat)
	if sumConcat == check {
		return check
	}

	return -1
}
