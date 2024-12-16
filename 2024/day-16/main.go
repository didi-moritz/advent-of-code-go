package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 16, realData))

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

var (
	moves = []v{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
)

func part1(data [][]byte) int {
	var start, end v

	found := make([][]int, len(data))

	for y, line := range data {
		found[y] = make([]int, len(line))
		for x, c := range line {
			found[y][x] = math.MaxInt
			if c == 'S' {
				start = v{x, y}
				data[y][x] = '.'
			} else if c == 'E' {
				end = v{x, y}
				data[y][x] = '.'
			}
		}
	}

	return move(data, end, start.x, start.y, 0, true, 0, found, make([]v, 0))
}

func move(data [][]byte, end v, x int, y int, dir int, canRotate bool, score int, found [][]int, visited []v) int {
	minimum := math.MaxInt

	p := v{x, y}

	if data[y][x] == '#' {
		return minimum
	}

	if found[y][x] > 0 {
		if score > found[y][x] {
			return minimum
		}
	}

	if p == end {
		return score
	}

	if slices.Contains(visited, p) {
		return minimum
	}

	visited = append(visited, p)

	found[y][x] = score

	for i, newDir := range []int{dir, dir - 1, dir + 1} {
		addScore := 1

		//if !canRotate && i > 1 {
		//	continue
		//}

		newCanRotate := true
		if i > 0 {
			addScore = 1001
			newCanRotate = false
		}

		if newDir < 0 {
			newDir += 4
		}

		newDir = newDir % 4

		newScore := move(data, end, x+moves[newDir].x, y+moves[newDir].y, newDir, newCanRotate, score+addScore, found, visited)
		if newScore < minimum {
			minimum = newScore
		}
	}

	return minimum
}

func part2(data [][]byte) int {
	return 0
}
