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
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 20, realData))

	unvisited := make([]v, 0)
	bricks := make([]v, 0)

	width := len(data[0])
	height := len(data)

	var start, end v

	for y, line := range data {
		for x, c := range line {
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				continue
			}

			p := v{x, y}

			if c == '#' {
				bricks = append(bricks, p)
			} else {
				if c == 'S' {
					start = p
				} else if c == 'E' {
					end = p
				}

				unvisited = append(unvisited, p)
			}
		}
	}

	var result int
	if part == 1 {
		result = part1(unvisited, bricks, start, end, realData)
	} else {
		result = part2(unvisited, start, end, realData)
	}

	fmt.Println(result)
}

var (
	moves = []v{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

func part1(way []v, bricks []v, start v, end v, realData bool) int {
	scores := calcInit(way, start, end)

	steps := make([]v, len(scores))
	for i := range len(scores) {
		for key, value := range scores {
			if value == i {
				steps[i] = key
			}
		}
	}

	targetDiff := 100
	if !realData {
		targetDiff = 12
	}

	result := 0
	for step := range steps {
		fmt.Println(step)
		result += calcWithCheat(step, targetDiff, scores, steps, bricks)
	}

	return result
}

func calcInit(way []v, start v, end v) map[v]int {
	unvisited := make([]v, len(way))
	copy(unvisited, way)

	scores := make(map[v]int)
	for _, p := range unvisited {
		scores[p] = math.MaxInt
	}
	scores[start] = 0

	finished := false
	for !finished {
		finished, unvisited = calcDijkstra(unvisited, scores, end, math.MaxInt, true)
	}

	return scores
}

func calcWithCheat(step int, diff int, scores map[v]int, steps []v, bricks []v) int {
	p := steps[step]

	result := 0

	// find all surrounding bricks
	for _, m := range moves {
		p2 := v{p.x + m.x, p.y + m.y}
		if slices.Contains(bricks, p2) {
			for _, m2 := range moves {
				p3 := v{p2.x + m2.x, p2.y + m2.y}

				stepsTo, found := scores[p3]

				if found {
					d := stepsTo - step - 2

					if d >= diff {
						result++
					}
				}
			}
		}
	}

	return result
}

func calcDijkstra(unvisited []v, scores map[v]int, end v, maxScore int, init bool) (bool, []v) {
	// find next candidate
	i := -1
	score := math.MaxInt
	for j, u := range unvisited {
		checkScore := scores[u]
		if checkScore < score {
			i = j
			score = checkScore
		}
	}

	if i == -1 {
		return true, nil
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

	if !init {
		if p == end {
			return true, nil
		}
	}

	if score > maxScore {
		return true, nil
	}

	return false, append(unvisited[:i], unvisited[i+1:]...)
}

func part2(way []v, start v, end v, realData bool) int {
	scores := calcInit(way, start, end)

	steps := make([]v, len(scores))
	for i := range len(scores) {
		for key, value := range scores {
			if value == i {
				steps[i] = key
			}
		}
	}

	targetDiff := 100
	maxCheat := 20
	if !realData {
		targetDiff = 74
		maxCheat = 20
	}

	result := 0
	for step := range steps {
		fmt.Println(step)
		result += calcWithCheat2(step, targetDiff, maxCheat, scores, steps)
	}

	return result
}

func calcWithCheat2(step int, targetDiff int, maxCheat int, scores map[v]int, steps []v) int {
	p := steps[step]

	result := 0

	// find all surrounding bricks
	for dx := -maxCheat; dx <= maxCheat; dx++ {
		for dy := -maxCheat; dy <= maxCheat; dy++ {
			cheats := abs(dx) + abs(dy)
			if cheats > maxCheat {
				continue
			}

			p2 := v{p.x + dx, p.y + dy}

			stepsTo, found := scores[p2]

			if found {
				d := stepsTo - step - cheats

				if d >= targetDiff {
					result++
				}
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
