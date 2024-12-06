package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2020, 5, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	result := 0
	for _, line := range data {
		row := make([]byte, 7)
		col := make([]byte, 3)
		for i, c := range line {
			if i < 7 {
				if c == 'B' {
					row[i] = '1'
				} else {
					row[i] = '0'
				}
			} else {
				if c == 'R' {
					col[i-7] = '1'
				} else {
					col[i-7] = '0'
				}
			}
		}
		rowN, _ := strconv.ParseInt(string(row), 2, 8)
		colN, _ := strconv.ParseInt(string(col), 2, 8)
		sum := int(rowN)*8 + int(colN)

		if sum > result {
			result = sum
		}
	}

	return result
}

func part2(data [][]byte) int {
	return 0
}
