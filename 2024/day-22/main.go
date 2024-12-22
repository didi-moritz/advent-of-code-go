package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 22, realData))

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
		n, _ := strconv.Atoi(line)
		for range 2000 {
			n = calcNextPseudoNumber(n)
		}
		result += n
	}

	return result
}

func calcNextPseudoNumber(number int) int {
	newNumber := prune(mix(number*64, number))
	newNumber = prune(mix(newNumber, newNumber/32))
	newNumber = prune(mix(newNumber, newNumber*2048))

	return newNumber
}

func mix(a int, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}
func part2(data []string) int {
	return 0
}
