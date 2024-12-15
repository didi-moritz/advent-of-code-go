package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"time"
)

type v struct {
	x, y int
}

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 15, realData))

	i := 0
	for ; len(data[i]) > 0; i++ {
	}

	var robot v

	m := data[0:i]

	for y, line := range data {
		for x, c := range line {
			if c == '@' {
				m[y][x] = '.'
				robot = v{x, y}
			}
		}
	}

	commandMap := map[byte]v{
		'<': {-1, 0},
		'>': {1, 0},
		'v': {0, 1},
		'^': {0, -1},
	}

	var commands []v

	for _, line := range data[i+1:] {
		for _, c := range line {
			commands = append(commands, commandMap[c])
		}
	}

	var result int
	if part == 1 {
		result = part1(m, robot, commands, !realData)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(m [][]byte, robot v, commands []v, showMap bool) int {
	for _, c := range commands {
		x := robot.x
		y := robot.y
		for {
			x += c.x
			y += c.y

			if m[y][x] == '.' || m[y][x] == '#' {
				break
			}
		}

		if m[y][x] == '.' {
			m[y][x] = 'O'

			robot.x += c.x
			robot.y += c.y

			m[robot.y][robot.x] = '.'
		}

		if showMap {
			printMap(m, robot)
			time.Sleep(10 * time.Millisecond)
		}
	}

	printMap(m, robot)

	result := 0
	for y, line := range m {
		for x, c := range line {
			if c == 'O' {
				result += y*100 + x
			}
		}
	}

	return result
}

func printMap(m [][]byte, robot v) {
	for y, line := range m {
		for x, c := range line {
			if robot.x == x && robot.y == y {
				fmt.Print("@")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func part2(data [][]byte) int {
	return 0
}
