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

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
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

	result := 0
	for _, m := range machines {
		result += calcTokens(m)
	}
	return result
}

func calcTokens(m machine) int {
	for ia := range 101 {
		xa := ia * m.a.x
		ya := ia * m.a.y

		if xa > m.prize.x || ya > m.prize.y {
			return 0
		}

		for ib := range 101 {
			xb := ib * m.b.x
			yb := ib * m.b.y

			if (xa+xb) == m.prize.x && (ya+yb) == m.prize.y {
				fmt.Println(ia, ib)
				tokens := ia*3 + ib
				return tokens
			}

			if (xa+xb) >= m.prize.x || (ya+yb) >= m.prize.y {
				break
			}

		}
	}

	return 0
}

type button struct {
	v      v
	tokens int
}

type machine2 struct {
	a, b  button
	prize v
}

func part2(data []string) int {
	var machines []machine2

	for i := range (len(data) + 1) / 4 {
		m := machine2{}
		var x, y int
		fmt.Sscanf(data[i*4], "Button A: X+%d, Y+%d", &x, &y)
		m.a = button{v{x, y}, 3}
		fmt.Sscanf(data[i*4+1], "Button B: X+%d, Y+%d", &x, &y)
		m.b = button{v{x, y}, 1}
		fmt.Sscanf(data[i*4+2], "Prize: X=%d, Y=%d", &x, &y)
		m.prize = v{x, y}

		machines = append(machines, m)
	}

	result := 0
	for _, m := range machines {
		m.prize.x += 10000000000000
		m.prize.y += 10000000000000
		tokens := calcTokens2(m)
		result += tokens
	}

	return result
}

func calcTokens2(m machine2) int {
	a := float64(m.a.v.x)
	b := float64(m.b.v.x)
	c := float64(m.prize.x)

	d := float64(m.a.v.y)
	e := float64(m.b.v.y)
	f := float64(m.prize.y)

	ib := int((f - (c * d / a)) / (e - (b * d / a)))

	for ibc := ib - 10; ibc < ib+10; ibc++ {

		x := ib * m.b.v.x
		y := ib * m.b.v.y

		rx := m.prize.x - x
		ry := m.prize.y - y

		var ia int
		if rx < ry {
			ia = rx / m.a.v.x
		} else {
			ia = ry / m.a.v.y
		}

		for iac := ia - 10; iac < ia+10; iac++ {
			if iac*m.a.v.x+ibc*m.b.v.x == m.prize.x && iac*m.a.v.y+ibc*m.b.v.y == m.prize.y {
				return iac*m.a.tokens + ibc*m.b.tokens
			}
		}
	}

	return 0
}
