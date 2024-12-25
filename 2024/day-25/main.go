package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 25, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	var keys, locks [][]int

	for i := 0; i < len(data); i += 8 {
		isLock := data[i][0] == '#'

		combination := make([]int, 5)
		for col := range 5 {
			c := 0
			for row := i + 1; row < i+6; row++ {
				if data[row][col] == '#' {
					c++
				}
			}
			combination[col] = c
		}

		if isLock {
			locks = append(locks, combination)
		} else {
			keys = append(keys, combination)
		}
	}

	result := 0
	for _, key := range keys {
		for _, lock := range locks {
			check := true
			for col := range 5 {
				if key[col]+lock[col] > 5 {
					check = false
					break
				}
			}

			if check {
				result++
			}
		}
	}

	return result
}

func part2(data [][]byte) int {
	return 0
}
