package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 11, realData))

	parts := strings.Split(data[0], " ")
	numbers := make([]int, len(parts))
	for i, p := range parts {
		numbers[i] = int(utils.ParseInt(p))
	}

	var result int
	if part == 1 {
		result = part1(numbers)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += calcNumber(n, 0)
	}

	return result
}

func calcNumber(n int, step int) int {
	if step == 25 {
		return 1
	}

	if n == 0 {
		return calcNumber(1, step+1)
	}

	nString := strconv.Itoa(n)

	if len(nString)%2 == 0 {
		return calcNumber(int(utils.ParseInt(nString[:len(nString)/2])), step+1) +
			calcNumber(int(utils.ParseInt(nString[len(nString)/2:])), step+1)
	}

	return calcNumber(n*2024, step+1)
}

func part2(data []string) int {
	return 0
}
