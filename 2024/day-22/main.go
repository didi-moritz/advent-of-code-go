package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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

type entry struct {
	score, diff int
}

func part2(data []string) int {
	allEntries := make([][]entry, len(data))

	result := 0
	for i, line := range data {
		allEntries[i] = make([]entry, 2000-1)
		previousNumber, _ := strconv.Atoi(line)
		for j := range 2000 - 1 {
			nextNumber := calcNextPseudoNumber(previousNumber)

			nextDigit := nextNumber % 10
			previousDigit := previousNumber % 10

			allEntries[i][j] = entry{nextDigit, nextDigit - previousDigit}

			previousNumber = nextNumber
		}
	}

	for i1 := -9; i1 < 10; i1++ {
		for i2 := -9; i2 < 10; i2++ {
			for i3 := -9; i3 < 10; i3++ {
				for i4 := -9; i4 < 10; i4++ {
					sequence := []int{i1, i2, i3, i4}
					newResult := calcBananas(sequence, allEntries)
					if newResult > result {
						result = newResult
					}
					fmt.Println(sequence, newResult)
				}
			}
		}
	}

	return result
}

func calcBananas(sequence []int, allEntries [][]entry) int {
	allBananas := 0
	for _, entries := range allEntries {
		found, bananas := findSequence(sequence, entries)
		if found {
			allBananas += bananas
		}
	}

	return allBananas
}

func findSequence(sequence []int, entries []entry) (bool, int) {
	for i := 0; i < len(entries)-3; i++ {
		check := true
		for j, s := range sequence {
			if entries[i+j].diff != s {
				check = false
				break
			}
		}

		if check {
			return true, entries[i+len(sequence)-1].score
		}
	}

	return false, 0
}
