package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 10, realData))

	numbers := make([]int, len(data))
	for i, line := range data {
		numbers[i], _ = strconv.Atoi(line)
	}

	sort.Ints(numbers)

	var result int
	if part == 1 {
		result = part1(numbers)
	} else {
		result = part2(numbers)
	}

	fmt.Println(result)
}

func part1(numbers []int) int {
	oneDiffCounts := 0
	threeDiffCounts := 0

	current := 0
	for _, number := range numbers {
		diff := number - current
		if diff == 1 {
			oneDiffCounts++
		}

		if diff == 3 {
			threeDiffCounts++
		}

		current = number
	}

	threeDiffCounts++

	fmt.Println(oneDiffCounts, threeDiffCounts)

	return oneDiffCounts * threeDiffCounts
}

func part2(numbers []int) int {
	allNumbers := make([]int, len(numbers)+1)
	allNumbers[0] = 0
	for i, number := range numbers {
		allNumbers[i+1] = number
	}

	cache := make(map[int]int)

	return findPermutations(allNumbers, 0, cache)
}

func findPermutations(numbers []int, i int, cache map[int]int) int {
	result := 0

	if i == len(numbers)-1 {
		return 1
	}

	cachedResult, found := cache[i]
	if found {
		return cachedResult
	}

	maxNextNumber := numbers[i] + 3
	for j := i + 1; j < len(numbers) && numbers[j] <= maxNextNumber; j++ {
		result += findPermutations(numbers, j, cache)
	}

	cache[i] = result

	return result
}
