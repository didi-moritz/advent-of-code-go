package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 10, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	width := len(data[0])
	height := len(data)

	m := make([][]int, width)
	for y, line := range data {
		m[y] = make([]int, width)

		for x, c := range line {
			m[y][x] = int(c - '0')
		}
	}

	result := 0
	for y := range height {
		for x := range width {
			var found []v
			getTrailHeadScore(m, x, y, 0, width, height, &found)
			result += len(found)
		}
	}
	return result
}

type v struct {
	x, y int
}

func getTrailHeadScore(m [][]int, x int, y int, step int, width int, height int, found *[]v) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return
	}

	if m[y][x] != step {
		return
	}

	if step == 9 {
		f := v{x, y}
		if !slices.Contains(*found, f) {
			*found = append(*found, f)
		}
		return

	}

	moves := []v{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	for _, move := range moves {
		getTrailHeadScore(m, x+move.x, y+move.y, step+1, width, height, found)
	}

	return
}

func part2(data [][]byte) int {
	width := len(data[0])
	height := len(data)

	m := make([][]int, width)
	for y, line := range data {
		m[y] = make([]int, width)

		for x, c := range line {
			m[y][x] = int(c - '0')
		}
	}

	result := 0
	for y := range height {
		for x := range width {
			result += getTrailHeadScore2(m, x, y, 0, width, height)
		}
	}
	return result
}

func getTrailHeadScore2(m [][]int, x int, y int, step int, width int, height int) int {
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0
	}

	if m[y][x] != step {
		return 0
	}

	if step == 9 {
		return 1

	}

	moves := []v{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	sum := 0
	for _, move := range moves {
		sum += getTrailHeadScore2(m, x+move.x, y+move.y, step+1, width, height)
	}

	return sum
}
