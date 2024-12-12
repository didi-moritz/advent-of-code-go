package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
	result := 0

	found := make([][]bool, len(data))
	foundFences := make([][][]bool, len(data))
	for y, line := range data {
		found[y] = make([]bool, len(line))
		foundFences[y] = make([][]bool, len(data))
	}

	moves := []v{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for y, line := range data {
		for x := range line {
			foundFences[y][x] = calcFence2(x, y, data, moves)
		}
	}

	for y, line := range data {
		for x := range line {
			area, fence := calcAreaAndFence2(x, y, data, found, moves, 0, 0, foundFences)
			if area > 0 {
				result += area * fence
			}
		}
	}

	return result
}

func calcAreaAndFence2(x int, y int, data [][]byte, found [][]bool, moves []v, area int, fence int, foundFences [][][]bool) (int, int) {
	if found[y][x] {
		return area, fence
	}

	found[y][x] = true

	c := data[y][x]

	area++
	for i, ff := range foundFences[y][x] {
		if ff {
			fence++
			// uniquify found fences
			clearFoundFences(x, y, i, c, moves[(i+1)%4], data, foundFences)
			clearFoundFences(x, y, i, c, moves[(i+3)%4], data, foundFences)
		}
	}

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

		area, fence = calcAreaAndFence2(nx, ny, data, found, moves, area, fence, foundFences)
	}

	return area, fence
}

func clearFoundFences(x int, y int, i int, c byte, m v, data [][]byte, foundFences [][][]bool) {
	width := len(data[0])
	height := len(data)

	x += m.x
	y += m.y

	if x < 0 || x >= width || y < 0 || y >= height || data[y][x] != c {
		return
	}

	if foundFences[y][x][i] {
		foundFences[y][x][i] = false
		clearFoundFences(x, y, i, c, m, data, foundFences)
	}
}

func calcFence2(x int, y int, data [][]byte, moves []v) []bool {
	c := data[y][x]

	width := len(data[0])
	height := len(data)

	foundFences := make([]bool, len(moves))

	for i, m := range moves {
		nx := x + m.x
		ny := y + m.y

		if nx < 0 || nx >= width || ny < 0 || ny >= height || data[ny][nx] != c {
			foundFences[i] = true
		}
	}

	return foundFences
}
