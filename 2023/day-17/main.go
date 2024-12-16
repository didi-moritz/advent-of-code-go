package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2023, 17, realData))

	numbers := make([][]int, len(data[0]))
	for x := range len(data[0]) {
		numbers[x] = make([]int, len(data))
		for y := range len(data) {
			numbers[x][y] = int(data[y][x] - '0')
		}
	}

	var result int
	if part == 1 {
		result = part1(numbers)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type v struct {
	x, y int
}

type cacheKey struct {
	dir, score int
}

var (
	cache [][]map[cacheKey]int
	moves = []v{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

func part1(numbers [][]int) int {
	cache = make([][]map[cacheKey]int, len(numbers))
	for x := range len(numbers) {
		cache[x] = make([]map[cacheKey]int, len(numbers[0]))
		for y := range len(numbers[0]) {
			cache[x][y] = make(map[cacheKey]int)
		}
	}

	minHeat := math.MaxInt
	move(0, 0, -1, 0, 0, numbers, make([]v, 0), &minHeat)
	return minHeat
}

func move(x int, y int, dir int, steps int, heat int, numbers [][]int, visited []v, minHeat *int) int {
	width := len(numbers)
	height := len(numbers[0])
	if x < 0 || x >= width || y < 0 || y >= height {
		return -1
	}

	if dir >= 0 {
		heat += numbers[x][y]
	}

	if heat > *minHeat {
		return -1
	}

	if x == width-1 && y == height-1 {
		if heat < *minHeat {
			*minHeat = heat
			printMap(numbers, visited)
			fmt.Println(heat)
		}
		return heat
	}

	p := v{x, y}

	if slices.Contains(visited, p) {
		return -1
	}

	newVisited := make([]v, len(visited))
	copy(newVisited, visited)
	visited = append(newVisited, p)

	k := cacheKey{dir, steps}
	{
		val, ok := cache[x][y][k]
		if ok {
			if heat >= val {
				return -1
			}
		}

		cache[x][y][k] = heat
	}

	minimum := math.MaxInt
	for i, m := range moves {
		if dir > -1 && i+2%4 == dir {
			continue
		}

		newSteps := 1

		if i == dir {
			newSteps = steps + 1
			if newSteps > 3 {
				continue
			}
		}

		fullHeat := move(x+m.x, y+m.y, i, newSteps, heat, numbers, visited, minHeat)
		if fullHeat > -1 && fullHeat < minimum {
			minimum = fullHeat
		}

		val, ok := cache[x][y][k]
		if ok {
			if heat > val {
				return -1
			}
		}
	}

	return minimum
}

func printMap(data [][]int, visited []v) {
	for y, line := range data {
		for x, c := range line {
			p := v{x, y}
			if slices.Contains(visited, p) {
				fmt.Print(" ")
			} else {
				fmt.Print(c)
			}
		}
		fmt.Println()
	}
}

func part2(data [][]byte) int {
	return 0
}
