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
	part, realData := utils.GetRunConfig(2, false)

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
		result = part2(obstacles, start, len(data[0]), len(data))
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

func part2(obstacles []v, start v, width int, height int) int {
	result := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			no := v{x, y}
			if no == start || slices.Contains(obstacles, no) {
				continue
			}

			newObstacles := make([]v, len(obstacles)+1)
			copy(newObstacles, obstacles)
			newObstacles[len(obstacles)] = no

			fmt.Println("checking", no)

			if isALoop(newObstacles, start, 0, width, height) {
				result++
			}
		}
	}

	return result
}

func isALoop(obstacles []v, pos v, dir int, width int, height int) bool {
	moves := []v{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	visits := make([][]int, width)

	for x := range width {
		visits[x] = make([]int, height)
	}

	for {
		nx := pos.x + moves[dir].x
		ny := pos.y + moves[dir].y

		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			return false
		}

		if visits[nx][ny] == (dir + 1) {
			return true
		}

		np := v{nx, ny}
		if slices.Contains(obstacles, np) {
			dir++
			dir %= len(moves)
		} else {
			visits[nx][ny] = dir + 1
			pos = np
		}
	}
}