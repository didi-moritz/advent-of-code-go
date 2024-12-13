package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

type v struct {
	x, y int
}

type machine struct {
	a, b, prize v
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 13, realData))

	var machines []machine

	for i := range (len(data) + 1) / 4 {
		m := machine{}
		var x, y int
		fmt.Sscanf(data[i*4], "Button A: X+%d, Y+%d", &x, &y)
		m.a = v{x, y}
		fmt.Sscanf(data[i*4+1], "Button B: X+%d, Y+%d", &x, &y)
		m.b = v{x, y}
		fmt.Sscanf(data[i*4+2], "Prize: X=%d, Y=%d", &x, &y)
		m.prize = v{x, y}

		machines = append(machines, m)
	}

	var result int
	if part == 1 {
		result = part1(machines)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(machines []machine) int {
	result := 0
	for _, m := range machines {
		result += calcTokens(m)
	}
	return result
}

func calcTokens(m machine) int {
	minimum := 0
	for ia := range 101 {
		xa := ia * m.a.x
		ya := ia * m.a.y

		if xa > m.prize.x || ya > m.prize.y {
			return minimum
		}

		for ib := range 101 {
			xb := ib * m.b.x
			yb := ib * m.b.y

			if (xa+xb) == m.prize.x && (ya+yb) == m.prize.y {
				tokens := ia*3 + ib
				if tokens < minimum || minimum == 0 {
					minimum = tokens
				}
			}

			if (xa+xb) >= m.prize.x || (ya+yb) >= m.prize.y {
				break
			}

		}
	}

	return minimum
}

func part2(data []string) int {
	return 0
}
