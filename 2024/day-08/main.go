package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 8, realData))

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
	antennasMap := make(map[byte][]v)

	for y, line := range data {
		for x, c := range line {
			if c != '.' {
				antennasMap[c] = append(antennasMap[c], v{x, y})
			}
		}
	}

	width := len(data[0])
	height := len(data)

	var signals []v

	for _, vs := range antennasMap {
		for _, a := range vs {
			for _, b := range vs {
				if a == b {
					continue
				}

				x := 2*a.x - b.x
				y := 2*a.y - b.y
				if x < 0 || x >= width || y < 0 || y >= height {
					continue
				}

				s := v{x, y}

				if !slices.Contains(signals, s) {
					signals = append(signals, s)
				}
			}
		}
	}

	return len(signals)
}

func part2(data [][]byte) int {
	return 0
}
