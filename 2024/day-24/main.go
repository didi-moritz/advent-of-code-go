package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
	"sort"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2024, 24, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type gate struct {
	in1, in2 string
	out      string
	operator string
	executed bool
}

func part1(data []string) int {
	register := make(map[string]bool)
	var zs []string

	i := 0
	line := ""
	for i, line = range data {
		if line == "" {
			break
		}
		var name string
		var value int
		fmt.Sscanf(line, "%s %d", &name, &value)
		name = name[:len(name)-1]
		register[name] = value == 1
	}

	gates := make([]gate, 0)

	for j := i + 1; j < len(data); j++ {
		var in1, in2, out, operator string
		fmt.Sscanf(data[j], "%s %s %s -> %s", &in1, &operator, &in2, &out)
		gates = append(gates, gate{in1, in2, out, operator, false})

		if out[0] == 'z' && !slices.Contains(zs, out) {
			zs = append(zs, out)
		}
	}

	sort.Strings(zs)

	done := false
	for !done {
		done = true

		for _, gate := range gates {
			if gate.executed {
				continue
			}

			v1, found1 := register[gate.in1]
			v2, found2 := register[gate.in2]

			if !found1 || !found2 {
				//done = false
				continue
			}

			var value bool
			switch gate.operator {
			case "AND":
				value = v1 && v2
			case "OR":
				value = v1 || v2
			case "XOR":
				value = v1 != v2
			}

			register[gate.out] = value
			gate.executed = true
		}

		for _, z := range zs {
			_, found := register[z]
			if !found {
				done = false
				break
			}
		}
	}

	result := 0
	for i = len(zs) - 1; i >= 0; i-- {
		result = result << 1
		if register[zs[i]] {
			result++
		}
	}

	return result
}

func part2(data []string) int {
	return 0
}
