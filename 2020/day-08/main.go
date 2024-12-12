package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 8, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type CommandType int

const (
	ACC CommandType = iota
	JMP
	NOP
)

var commandMap = map[string]CommandType{
	"acc": ACC,
	"jmp": JMP,
	"nop": NOP,
}

type Instruction struct {
	command CommandType
	value   int
}

func part1(data []string) int {
	visited := make([]bool, len(data))

	is := make([]Instruction, len(data))

	for i, line := range data {
		parts := strings.Split(line, " ")
		command := commandMap[parts[0]]
		value := int(utils.ParseInt(parts[1]))
		is[i] = Instruction{command, value}
	}

	accValue := 0
	i := 0

	for {
		if visited[i] == true {
			return accValue
		}

		visited[i] = true

		if is[i].command == JMP {
			i += is[i].value
		} else {
			if is[i].command == ACC {
				accValue += is[i].value
			}

			i++
		}
	}

	return 0
}

func part2(data []string) int {
	return 0
}
