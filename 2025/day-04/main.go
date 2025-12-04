package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 4, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func printMap(m *[][]bool) {
	for _, line := range *m {
		for _, c := range line {
			if c {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func checkAndStore(r *[][]bool, m *[][]bool, x int, y int) {
	h := len(*m)
	w := len((*m)[0])

	if !(*m)[y][x] {
		return
	}

	found := 0
	for cx := x - 1; cx <= x+1; cx++ {
		for cy := y - 1; cy <= y+1; cy++ {
			if cx < 0 || cx == w || cy < 0 || cy == h || (cx == x && cy == y) {
				continue
			}

			if (*m)[cy][cx] {
				found++
				if found > 3 {
					return
				}
			}
		}
	}

	(*r)[y][x] = true
}

func part1(data []string) int {
	result := 0

	m := make([][]bool, len(data))
	r := make([][]bool, len(data))
	for y, line := range data {
		m[y] = make([]bool, len(line))
		r[y] = make([]bool, len(line))

		for x, c := range line {
			if c == '@' {
				m[y][x] = true
			}
		}
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			checkAndStore(&r, &m, x, y)
		}
	}

	printMap(&r)

	for y := 0; y < len(r); y++ {
		for x := 0; x < len(r[y]); x++ {
			if r[y][x] {
				result++
			}
		}
	}

	return result
}

func part2(data []string) int {
	result := 0

	return result
}
