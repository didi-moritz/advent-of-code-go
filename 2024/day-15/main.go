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
	part, realData := utils.GetRunConfig(2, false)

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
		result = part2(m, robot, commands, !realData)
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
			time.Sleep(100 * time.Millisecond)
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

func part2(m [][]byte, robot v, commands []v, showMap bool) int {
	for y, line := range m {
		var newLine []byte
		for _, c := range line {
			if c == 'O' {
				newLine = append(newLine, '[')
				newLine = append(newLine, ']')
			} else {
				newLine = append(newLine, c)
				newLine = append(newLine, c)
			}
		}
		m[y] = newLine
	}

	robot.x = robot.x * 2

	printMap(m, robot)

	for _, c := range commands {
		x := robot.x + c.x
		y := robot.y + c.y

		newPos := v{x, y}

		stillOk := !isBlocked(m, newPos)

		if stillOk {
			if isBlock(m, newPos) {
				stillOk = moveBoxes(m, newPos, c, false)
				if stillOk {
					moveBoxes(m, newPos, c, true)
				}
			}
		}

		if stillOk {
			robot.x += c.x
			robot.y += c.y
		}

		if showMap {
			printMap(m, robot)
			fmt.Println(c)
			time.Sleep(10 * time.Millisecond)
		}
	}

	printMap(m, robot)

	result := 0
	for y, line := range m {
		for x, c := range line {
			if c == '[' {
				result += y*100 + x
			}
		}
	}

	return result
}

func moveBoxes(m [][]byte, from v, move v, doPush bool) bool {
	if isFree(m, from) {
		return true
	}

	verticalMove := move.x == 0

	part1 := from
	part2 := findOtherPartOfBlock(m, part1)

	to1 := v{part1.x + move.x, part1.y + move.y}
	to2 := v{part2.x + move.x, part2.y + move.y}

	if isBlocked(m, to1) || isBlocked(m, to2) {
		return false
	}

	if verticalMove {
		if isBlock(m, to1) {
			if !moveBoxes(m, to1, move, doPush) {
				return false
			}
		}

		if isBlock(m, to2) {
			if !moveBoxes(m, to2, move, doPush) {
				return false
			}
		}
	} else {
		relevantTo := to1
		if to1 == part2 {
			relevantTo = to2
		}

		if isBlocked(m, relevantTo) {
			return false
		}

		if isBlock(m, relevantTo) {
			if !moveBoxes(m, relevantTo, move, doPush) {
				return false
			}
		}
	}

	if doPush {
		moveBox(m, part1, part2, move)
	}
	return true
}

func moveBox(m [][]byte, part1 v, part2 v, move v) {
	c1 := m[part1.y][part1.x]
	c2 := m[part2.y][part2.x]
	m[part1.y][part1.x] = '.'
	m[part2.y][part2.x] = '.'
	m[part1.y+move.y][part1.x+move.x] = c1
	m[part2.y+move.y][part2.x+move.x] = c2
}

func isBlock(m [][]byte, p v) bool {
	return m[p.y][p.x] == '[' || m[p.y][p.x] == ']'
}

func findOtherPartOfBlock(m [][]byte, p v) v {
	if m[p.y][p.x] == '[' {
		return v{p.x + 1, p.y}
	} else {
		return v{p.x - 1, p.y}
	}
}

func isBlocked(m [][]byte, p v) bool {
	return m[p.y][p.x] == '#'
}

func isFree(m [][]byte, p v) bool {
	return m[p.y][p.x] == '.'
}
