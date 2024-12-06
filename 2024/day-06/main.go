package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
)

type v struct {
	x, y int
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 6, realData))

	var obstacles []v

	var start v

	for y := range data {
		for x := range data[0] {
			if data[y][x] == '#' {
				obstacles = append(obstacles, v{x, y})
			} else if data[y][x] == '^' {
				start = v{x, y}
			}
		}
	}

	var result int
	if part == 1 {
		result = part1(obstacles, start, len(data[0]), len(data))
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(obstacles []v, start v, width int, height int) int {
	visits := move(obstacles, start, 0, width, height)

	return len(visits)
}

func move(obstacles []v, pos v, dir int, width int, height int) []v {
	moves := []v{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	var visits []v

	for {
		nx := pos.x + moves[dir].x
		ny := pos.y + moves[dir].y

		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			return visits
		}

		np := v{nx, ny}
		if slices.Contains(obstacles, np) {
			dir++
			dir %= len(moves)
		} else {
			if !slices.Contains(visits, np) {
				visits = append(visits, np)
			}
			pos = np
		}
	}
}

func part2(data [][]byte) int {
	return 0
}
