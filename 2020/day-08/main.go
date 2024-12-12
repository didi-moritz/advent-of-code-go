package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
}

func part2(data []string) int {
	is := make([]Instruction, len(data))

	for i, line := range data {
		parts := strings.Split(line, " ")
		command := commandMap[parts[0]]
		value := int(utils.ParseInt(parts[1]))
		is[i] = Instruction{command, value}
	}

	for i, instruction := range is {
		if instruction.command == ACC {
			continue
		}

		if instruction.command == NOP && (instruction.value == 0 || instruction.value == 1) {
			continue
		}

		if instruction.command == JMP && instruction.value == 1 {
			continue
		}

		newIs := make([]Instruction, len(is))
		copy(newIs, is)

		var newCommand CommandType
		if instruction.command == NOP {
			newCommand = JMP
		} else {
			newCommand = NOP
		}

		newIs[i].command = newCommand

		success, value := checkInstructions(newIs)
		if success {
			return value
		}
	}

	return 0
}

func checkInstructions(is []Instruction) (bool, int) {
	visited := make([]bool, len(is))

	accValue := 0
	i := 0

	for {
		if i == len(is) {
			return true, accValue
		}

		if i > len(is) {
			return false, 0
		}

		if visited[i] == true {
			return false, 0
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
}
