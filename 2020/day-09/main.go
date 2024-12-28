package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 9, realData))

	var result int
	if part == 1 {
		result = part1(data, realData)
	} else {
		result = part2(data, realData)
	}

	fmt.Println(result)
}

func part1(data []string, realData bool) int {
	considerCount := 25
	if !realData {
		considerCount = 5
	}

	numbers := make([]int, len(data))
	for i, line := range data {
		numbers[i], _ = strconv.Atoi(line)
	}

	for i := considerCount; i < len(numbers); i++ {
		consideredNumbers := numbers[i-considerCount : i]

		if !check(numbers[i], consideredNumbers) {
			return numbers[i]
		}
	}

	return 0
}

func check(number int, numbers []int) bool {
	for i1, n1 := range numbers[:len(numbers)-1] {
		for _, n2 := range numbers[i1+1:] {
			if n1+n2 == number {
				return true
			}
		}
	}

	return false
}

func part2(data []string, realData bool) int {
	numbers := make([]int, len(data))
	for i, line := range data {
		numbers[i], _ = strconv.Atoi(line)
	}

	invalidNumber := part1(data, realData)

	currentSum := 0

	currentLow := 0
	currentHigh := 0

	for i, n := range numbers {
		if currentSum == invalidNumber && currentLow != currentHigh {
			break
		}

		if currentSum < invalidNumber {
			currentHigh = i
			currentSum += n
		}

		if currentSum > invalidNumber {
			for j := currentLow; j < currentHigh; j++ {
				if currentSum <= invalidNumber {
					break
				}

				currentSum -= numbers[j]
				currentLow = j + 1
			}
		}
	}

	fmt.Println(currentLow, currentHigh)

	minimum := math.MaxInt
	maximum := 0

	for _, n := range numbers[currentLow : currentHigh+1] {
		if n < minimum {
			minimum = n
		}
		if n > maximum {
			maximum = n
		}
	}

	return minimum + maximum
}
