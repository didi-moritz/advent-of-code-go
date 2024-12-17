package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 17, realData))

	var a, b, c int
	fmt.Sscanf(data[0], "Register A: %d", &a)
	fmt.Sscanf(data[1], "Register B: %d", &b)
	fmt.Sscanf(data[2], "Register C: %d", &c)

	var instructions string
	fmt.Sscanf(data[4], "Program: %s", &instructions)

	fmt.Println(a, b, c, instructions)

	var result string
	if part == 1 {
		result = part1(a, b, c, instructions)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(a int, b int, c int, instructionsString string) string {
	parts := strings.Split(instructionsString, ",")
	instructions := make([]int, len(parts))
	for i, p := range parts {
		instructions[i], _ = strconv.Atoi(p)
	}

	result := ""

	p := 0
	for p < len(instructions) {
		opCode := instructions[p]
		operand := instructions[p+1]

		value := getValue(operand, a, b, c)

		switch opCode {
		case 0:
			a = a >> value
		case 1:
			b = b ^ operand
		case 2:
			b = value & 7
		case 3:
			if a > 0 {
				p = operand
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			if len(result) > 0 {
				result += ","
			}
			result += strconv.Itoa(value & 7)
		case 6:
			b = a >> value
		case 7:
			c = a >> value
		}

		p += 2
	}

	fmt.Println(a, b, c)
	return result
}

func getValue(operand int, a int, b int, c int) int {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}

	return -1
}

func part2(data []string) string {
	return "-"
}
