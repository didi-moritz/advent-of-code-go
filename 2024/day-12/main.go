package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 12, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type v struct {
	x, y int
}

func part1(data [][]byte) int {
	result := 0

	found := make([][]bool, len(data))
	for y, line := range data {
		found[y] = make([]bool, len(line))
	}

	moves := []v{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	for y, line := range data {
		for x := range line {
			area, fence := calcAreaAndFence(x, y, data, found, moves, 0, 0)
			result += area * fence
		}
	}

	return result
}

func calcAreaAndFence(x int, y int, data [][]byte, found [][]bool, moves []v, area int, fence int) (int, int) {
	if found[y][x] {
		return area, fence
	}

	found[y][x] = true

	area++
	fence += calcFence(x, y, data, moves)

	c := data[y][x]

	width := len(data[0])
	height := len(data)

	for _, m := range moves {
		nx := x + m.x
		ny := y + m.y

		if nx < 0 || nx >= width || ny < 0 || ny >= height || data[ny][nx] != c {
			continue
		}

		if found[ny][nx] {
			continue
		}

		area, fence = calcAreaAndFence(nx, ny, data, found, moves, area, fence)
	}

	return area, fence
}

func calcFence(x int, y int, data [][]byte, moves []v) int {
	c := data[y][x]

	width := len(data[0])
	height := len(data)

	fence := 0
	for _, m := range moves {
		nx := x + m.x
		ny := y + m.y

		if nx < 0 || nx >= width || ny < 0 || ny >= height || data[ny][nx] != c {
			fence++
		}
	}

	return fence
}

func part2(data [][]byte) int {
	return 0
}
