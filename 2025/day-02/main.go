package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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

func isInvalid2(number string) bool {
	length := len(number)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}

		word := number[:i]

		success := true
		for j := i; j < length; j += i {
			if number[j:j+i] != word {
				success = false
				break
			}
		}

		if success {
			return true
		}
	}

	return false
}

func part2(data []string) int {
	fmt.Println(isInvalid2("446446"))
	result := 0

	for _, line := range data {
		words := strings.Split(line, ",")

		for _, word := range words {
			var start int
			var end int
			fmt.Sscanf(word, "%d-%d", &start, &end)
			fmt.Println(word)

			for n := start; n <= end; n++ {
				number := strconv.Itoa(n)
				if isInvalid2(number) {
					fmt.Println(number)
					result += n
				}

			}
		}
	}

	return result
}
