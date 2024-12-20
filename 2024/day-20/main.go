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

type cacheValue struct {
	unvisited []v
	scores    map[v]int
	score     int
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

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
		result = part1(unvisited, bricks, start, end, width, height, realData)
	} else {
		result = part2()
	}

	fmt.Println(result)
}

var (
	moves = []v{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

func part1(way []v, bricks []v, start v, end v, width int, height int, realData bool) int {

	baseScore, cache := calcInit(way, start, end)

	fmt.Println(baseScore)

	scores := cache[len(cache)-1].scores

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

func calcInit(way []v, start v, end v) (int, []cacheValue) {
	var cache []cacheValue

	unvisited := make([]v, len(way))
	copy(unvisited, way)

	scores := make(map[v]int)
	for _, p := range unvisited {
		scores[p] = math.MaxInt
	}
	scores[start] = 0

	finished := false
	for !finished {
		finished, unvisited, cache = calcDijkstra(unvisited, scores, end, math.MaxInt, true, cache)
	}

	return scores[end], cache
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

func calcDijkstra(unvisited []v, scores map[v]int, end v, maxScore int, init bool, cache []cacheValue) (bool, []v, []cacheValue) {
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
		return true, nil, cache
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
			return true, nil, cache
		}
	}

	if score > maxScore {
		return true, nil, cache
	}

	if init {
		cachedUnvisited := make([]v, len(unvisited))
		copy(cachedUnvisited, unvisited)

		cachedScores := make(map[v]int)
		for key, value := range scores {
			//if value >= score {
			cachedScores[key] = value
			//}
		}
		cache = append(cache, cacheValue{cachedUnvisited, cachedScores, score})
	}

	return false, append(unvisited[:i], unvisited[i+1:]...), cache
}

func part2() int {
	return 0
}
