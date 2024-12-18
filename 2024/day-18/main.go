package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"slices"
)

type v struct {
	x, y int
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 18, realData))

	bs := make([]v, len(data))

	for i, line := range data {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		bs[i] = v{x, y}
	}

	var result int
	if part == 1 {
		result = part1(bs, realData)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

var (
	moves = []v{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

func part1(bs []v, realData bool) int {
	width := 71
	height := 71
	if !realData {
		width = 7
		height = 7
	}

	bytesToConsiderCount := 1024
	if !realData {
		bytesToConsiderCount = 12
	}

	bytesToConsider := bs[0:bytesToConsiderCount]

	unvisited := make([]v, 0)
	scores := make(map[v]int)
	for x := range width {
		for y := range height {
			p := v{x, y}

			if slices.Contains(bytesToConsider, p) {
				continue
			}

			unvisited = append(unvisited, p)

			if x == 0 && y == 0 {
				scores[p] = 0
			} else {
				scores[p] = math.MaxInt
			}
		}
	}

	for len(unvisited) > 0 {
		unvisited = calcDijkstra(unvisited, scores)
	}

	return scores[v{width - 1, height - 1}]
}

func calcDijkstra(unvisited []v, scores map[v]int) []v {
	// find next candidate
	i := 0
	score := math.MaxInt
	for j, u := range unvisited {
		checkScore := scores[u]
		if checkScore < score {
			i = j
			score = checkScore
		}
	}

	p := unvisited[i]

	for _, m := range moves {
		p2 := v{p.x + m.x, p.y + m.y}
		if slices.Contains(unvisited, p2) {
			currentScore := scores[p2]
			if currentScore > score+1 {
				scores[p2] = score + 1
			}
		}
	}

	return append(unvisited[:i], unvisited[i+1:]...)
}

func part2(data []string) int {
	return 0
}
