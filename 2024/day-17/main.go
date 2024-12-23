package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
		result = part2()
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

		value := getValue(operand, a, b)

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

func getValue(operand int, a int, b int) int {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return a
	case 5:
		return b
	case 6:
		panic("c")
		//return c
	}

	return -1
}

func part2() string {
	target := "2,4,1,1,7,5,1,5,4,5,0,3,5,5,3,0"

	parts := strings.Split(target, ",")

	targetBytes := make([]int, len(parts))

	for i, part := range parts {
		targetBytes[i], _ = strconv.Atoi(part)
	}

	bytes := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	consider := 3

	var number int

	for i := 0; i < len(bytes)-consider; i++ {

		for j := 0; j < consider; j++ {
			bytes[i+j] = 0
		}

		for {
			bytes[i]++

			for j := i; j < len(bytes)-1; j++ {
				if bytes[j] == 8 {
					bytes[j] = 0
					bytes[j+1]++
				}
			}

			number = calcNumber(bytes)

			result := calc2(number)

			ok := true
			for j := 0; j < i+consider; j++ {
				if result[j] != targetBytes[j] {
					ok = false
					break
				}
			}
			if ok {
				fmt.Println(number)
				fmt.Println(targetBytes)
				fmt.Println(result)
				break
			}
		}
	}

	return strconv.Itoa(number)
}

func calcNumber(bytes []int) int {
	number := 0
	for i := range len(bytes) {
		number = number << 3
		number += bytes[len(bytes)-1-i]
	}
	return number
}

func calc2(a int) []int {
	b := 0
	c := 0

	//2,4,  b = a & 7
	//1,1,  b = b ^ 1  - 001
	//7,5,  c = a >> b
	//1,5,  b = b ^ 5  - 101
	//4,5,  b = b & c
	//0,3,  a = a >> 3
	//5,5,  out b
	//3,0

	var result []int

	for a > 0 {
		b = a & 7
		b = b ^ 1
		c = a >> b
		b = b ^ 5
		b = b ^ c
		a = a >> 3

		r := b & 7
		result = append(result, r)
	}

	return result
}
