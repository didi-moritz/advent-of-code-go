package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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

	return move(data, end, start.x, start.y, 0, 0, found)
}

func move(data [][]byte, end v, x int, y int, dir int, score int, found [][]int) int {
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

	found[y][x] = score

	for i, newDir := range []int{dir, dir - 1, dir + 1} {
		addScore := 1

		if i > 0 {
			addScore = 1001
		}

		if newDir < 0 {
			newDir += 4
		}

		newDir = newDir % 4

		newScore := move(data, end, x+moves[newDir].x, y+moves[newDir].y, newDir, score+addScore, found)
		if newScore < minimum {
			minimum = newScore
		}
	}

	return minimum
}

func part2(data [][]byte) int {
	var start, end v

	found := make([][][]int, len(data))

	for y, line := range data {
		found[y] = make([][]int, len(line))
		for x, c := range line {
			found[y][x] = make([]int, 4)
			for d := range 4 {
				found[y][x][d] = math.MaxInt
			}
			if c == 'S' {
				start = v{x, y}
				data[y][x] = '.'
			} else if c == 'E' {
				end = v{x, y}
				data[y][x] = '.'
			}
		}
	}

	score, visited := move2(data, end, start.x, start.y, 0, 0, found, make([]v, 0))
	fmt.Println(score)
	printMap(data, visited)
	return len(visited)
}

func move2(data [][]byte, end v, x int, y int, dir int, score int, found [][][]int, visited []v) (int, []v) {
	minimum := math.MaxInt

	p := v{x, y}

	if data[y][x] == '#' {
		return minimum, nil
	}

	if found[y][x][dir] > 0 {
		if score > found[y][x][dir] {
			return minimum, visited
		}
	}

	if p == end {
		newVisited := make([]v, len(visited))
		copy(newVisited, visited)
		return score, append(newVisited, p)
	}

	if slices.Contains(visited, p) {
		return minimum, nil
	}

	newVisited := make([]v, len(visited))
	copy(newVisited, visited)
	visited = append(newVisited, p)

	found[y][x][dir] = score

	var returnVisited []v

	for i, newDir := range []int{dir, dir - 1, dir + 1} {
		addScore := 1

		if i > 0 {
			addScore = 1001
		}

		if newDir < 0 {
			newDir += 4
		}

		newDir = newDir % 4

		newScore, newVisited := move2(data, end, x+moves[newDir].x, y+moves[newDir].y, newDir, score+addScore, found, visited)

		if newScore < minimum && newScore < math.MaxInt {
			minimum = newScore
			returnVisited = newVisited
		} else if newScore <= minimum && newScore < math.MaxInt {
			minimum = newScore
			for _, v := range newVisited {
				if !slices.Contains(returnVisited, v) {
					returnVisited = append(returnVisited, v)
				}
			}
		}
	}

	return minimum, returnVisited
}

func printMap(data [][]byte, visited []v) {
	for y, line := range data {
		for x, c := range line {
			p := v{x, y}
			if slices.Contains(visited, p) {
				fmt.Print("O")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}
