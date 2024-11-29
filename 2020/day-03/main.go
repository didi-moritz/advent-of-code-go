package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2020, 3, realData))

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

type move struct {
	dx int
	dy int
}

func part2(data [][]byte) int {
	moves := []move{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, m := range moves {
		result *= processForPart2(data, m)
	}

	return result
}

func processForPart2(data [][]byte, m move) int {
	result := 0
	x := 0
	width := len(data[0])
	for y := 0; y < len(data); y += m.dy {
		if data[y][x] == '#' {
			result++
		}

		x += m.dx
		if x >= width {
			x -= width
		}
	}

	return result
}
