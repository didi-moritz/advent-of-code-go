package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, true)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2020, 3, realData))

	var result int
	if part == 1 {
		result = part1(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	result := 0
	x := 0
	width := len(data[0])
	for y := 0; y < len(data); y++ {
		if data[y][x] == '#' {
			result++
		}

		x += 3
		if x >= width {
			x -= width
		}
	}

	return result
}
