package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
		result = part2(numbers)
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

func part2(numbers []int) int {

	cache := make(map[int]map[int]int)
	result := 0
	for _, n := range numbers {
		result += calcNumber2(n, 0, cache)
	}

	return result
}

func calcNumber2(n int, step int, cache map[int]map[int]int) int {
	if step == 75 {
		return 1
	}

	if cache[n] != nil {
		if cache[n][step] > 0 {
			return cache[n][step]
		}
	}

	if n == 0 {
		result := calcNumber2(1, step+1, cache)
		initCacheForNumber(cache, n)
		cache[n][step] = result
		return result
	}

	nString := strconv.Itoa(n)

	if len(nString)%2 == 0 {
		result := calcNumber2(int(utils.ParseInt(nString[:len(nString)/2])), step+1, cache) +
			calcNumber2(int(utils.ParseInt(nString[len(nString)/2:])), step+1, cache)
		initCacheForNumber(cache, n)
		cache[n][step] = result
		return result
	}

	return calcNumber2(n*2024, step+1, cache)
}

func initCacheForNumber(cache map[int]map[int]int, n int) {
	if cache[n] == nil {
		cache[n] = make(map[int]int)
	}
}
