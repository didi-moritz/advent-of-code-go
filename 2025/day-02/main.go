package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 2, realData))

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
		words := strings.Split(line, ",")

		for _, word := range words {
			var start int
			var end int
			fmt.Sscanf(word, "%d-%d", &start, &end)

			fmt.Println(word, start, end)

			for n := start; n <= end; n++ {
				number := strconv.Itoa(n)
				l := len(number)
				if l%2 == 1 {
					continue
				}

				if number[l/2:] == number[:l/2] {
					fmt.Println(number)
					result += n
				}

			}
		}
	}

	return result
}

func part2(data []string) int {
	return 0
}
