package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2025, 8, realData))

	var result int
	if part == 1 {
		result = part1(data, realData)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type P struct {
	x, y, z int
	name    string
}

type Pair struct {
	p1, p2   P
	distance float64
}

func distance(p1 P, p2 P) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)))
}

func part1(data []string, realData bool) int {

	ps := make([]P, len(data))

	for i, line := range data {
		parts := strings.Split(line, ",")
		numbers := make([]int, len(parts))
		for j, p := range parts {
			numbers[j] = int(utils.ParseInt(p))
		}

		ps[i] = P{numbers[0], numbers[1], numbers[2], strconv.Itoa(i)}
	}

	var pairs []Pair

	for i, p1 := range ps {
		for j, p2 := range ps {
			if j <= i {
				continue
			}

			pairs = append(pairs, Pair{p1, p2, distance(p1, p2)})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	var gs [][]P

	count := 10
	if realData {
		count = 1000
	}

	for i := 0; i < count; i++ {
		pair := pairs[i]

		found := false
		for j, g := range gs {
			if slices.Contains(g, pair.p1) && slices.Contains(g, pair.p2) {
				found = true
			} else if slices.Contains(g, pair.p1) {
				merged := false

				for k, g2 := range gs {
					if slices.Contains(g2, pair.p2) {
						gs[j] = append(g, g2...)
						gs[k] = make([]P, 0)
						merged = true
					}
				}

				if !merged {
					gs[j] = append(g, pair.p2)
				}
				found = true
			} else if slices.Contains(g, pair.p2) {
				merged := false

				for k, g2 := range gs {
					if slices.Contains(g2, pair.p1) {
						gs[j] = append(g, g2...)
						gs[k] = make([]P, 0)
						merged = true
					}
				}

				if !merged {
					gs[j] = append(g, pair.p1)
				}
				found = true

			}

			if found {
				break
			}
		}

		if found {
			continue
		}

		var newG []P
		newG = append(newG, pair.p1)
		newG = append(newG, pair.p2)

		gs = append(gs, newG)
	}

	sort.Slice(gs, func(i, j int) bool {
		return len(gs[i]) > len(gs[j])
	})

	for _, g := range gs {
		fmt.Println(len(g), g)
	}

	result := 1
	for i := 0; i < 3; i++ {
		result *= len(gs[i])
	}

	return result
}

func part2(data []string) int {

	result := 0

	return result
}
