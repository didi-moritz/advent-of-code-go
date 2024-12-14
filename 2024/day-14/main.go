package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

type v struct {
	x, y int
}

type robot struct {
	p, v v
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 14, realData))

	var robots []robot

	for _, line := range data {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, robot{v{px, py}, v{vx, vy}})
	}

	var result int
	if part == 1 {
		width := 11
		height := 7
		if realData {
			width = 101
			height = 103
		}
		result = part1(robots, width, height)
	} else {
		result = part2(robots)
	}

	fmt.Println(result)
}

func part1(robots []robot, width int, height int) int {
	mx := width / 2
	my := height / 2

	score := make([]int, 4)

	for _, r := range robots {
		x := (r.p.x + r.v.x*100) % width
		y := (r.p.y + r.v.y*100) % height

		if x < 0 {
			x += width
		}

		if y < 0 {
			y += height
		}

		if x == mx || y == my {
			continue
		}

		i := 0
		if x > mx {
			i++
		}
		if y > my {
			i += 2
		}

		score[i]++
	}

	result := 1
	for _, s := range score {
		result *= s
	}

	return result
}

func part2(robots []robot) int {
	return 0
}
