package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(1, true)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 4, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type vector struct {
	x, y int
}

func part1(data [][]byte) int {
	word := "XMAS"
	found := make([][]bool, len(data))
	for i := range data {
		found[i] = make([]bool, len(data[i]))

	}

	ms := []vector{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	height := len(data)
	width := len(data[0])

	result := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for _, m := range ms {
				cx := x
				cy := y
				match := false
				for _, c := range word {
					if cx < 0 || cx >= width || cy < 0 || cy >= height {
						break
					}

					if byte(c) != data[cy][cx] {
						break
					}

					if byte(c) == word[len(word)-1] {
						match = true
						break
					}
					cx += m.x
					cy += m.y
				}

				if match {
					result++
					cx = x
					cy = y
					for i := 0; i < len(word); i++ {
						found[cy][cx] = true
						cx += m.x
						cy += m.y
					}
				}
			}

		}
	}

	print_data(data, found)

	return result
}

func part2(data [][]byte) int {
	return 0
}

func print_data(data [][]byte, found [][]bool) {
	colorRed := "\033[0;31m"
	colorNone := "\033[0m"

	for y := range data {
		for x := range data[y] {
			if found[y][x] {
				fmt.Print(colorRed)
			}
			fmt.Printf("%c", data[y][x])
			fmt.Print(colorNone)
		}
		fmt.Println()
	}
}
