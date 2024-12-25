package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(2, true)

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
		ins := []string{in1, in2}
		sort.Strings(ins)
		gates = append(gates, gate{ins[0], ins[1], out, operator, false})

		if out[0] == 'z' && !slices.Contains(zs, out) {
			zs = append(zs, out)
		}
	}

	var out []string

	for _, gate := range gates {
		if gate.in1[0] == 'x' || gate.in1[0] == 'y' {
			out = append(out, gateToString(gate))
		}
	}

	sort.Strings(zs)

	// found wrong gates
	gates = swapGateOuts("bpf", "z05", gates)
	gates = swapGateOuts("hcc", "z11", gates)
	gates = swapGateOuts("qcw", "hqc", gates)
	gates = swapGateOuts("fdw", "z35", gates)

	// print final result
	result := []string{"bpf", "z05", "hcc", "z11", "qcw", "hqc", "fdw", "z35"}
	sort.Strings(result)
	fmt.Println(strings.Join(result, ","))
	panic("ok!")

	overflows := make(map[int]string)
	for j, z := range zs {

		fmt.Printf("### %s ###\n", z)
		fmt.Println(overflows)
		found := make([]gate, 0)

		neededOverflow := ""
		if j > 0 {
			neededOverflow = overflows[j-1]
		}

		findParents(z, gates, &found, neededOverflow)
		fmt.Println("found parents, should be two XOR")
		for _, gate := range found {
			fmt.Println(gateToString(gate))
		}

		if j > 0 {
			if len(found) != 2 {
				panic("foo")
			}

			for _, gate := range found {
				if gate.operator != "XOR" {
					panic("bar")
				}
			}
		}

		overflowGate := findOverflowGate(z, gates, neededOverflow)
		overflows[j] = overflowGate.out

		fmt.Println()
	}

	return 0
}

func gateToString(gate gate) string {
	return fmt.Sprintf("%s %s %s -> %s", gate.in1, gate.operator, gate.in2, gate.out)
}

func findParents(name string, gates []gate, found *[]gate, overflow string) {
	if name == overflow {
		return
	}

	for _, gate := range gates {
		if gate.out == name {
			if !slices.Contains(*found, gate) {
				*found = append(*found, gate)
				findParents(gate.in1, gates, found, overflow)
				findParents(gate.in2, gates, found, overflow)
			}
		}
	}
}

func findOverflowGate(z string, gates []gate, previousOverflow string) gate {
	zNumber, _ := strconv.Atoi(z[1:])

	xyNumber := twoDigitNumber(zNumber)

	x := "x" + xyNumber
	y := "y" + xyNumber

	if z == "z00" {
		for _, gate := range gates {
			if gate.in1 == x && gate.in2 == y && gate.operator == "AND" {
				return gate
			}
		}

		panic("first overflow gate not found")
	}

	xorGateName := ""
	for _, gate := range gates {
		if gate.in1 == x && gate.in2 == y && gate.operator == "XOR" {
			xorGateName = gate.out
			break
		}
	}
	if xorGateName == "" {
		panic("xorGateName gate not found")
	}

	andGateName := ""
	for _, gate := range gates {
		if gate.in1 == x && gate.in2 == y && gate.operator == "AND" {
			andGateName = gate.out
			break
		}
	}
	if andGateName == "" {
		panic("andGateName gate not found")
	}

	andGateNameWithPreviousOverflow := ""
	for _, gate := range gates {
		if ((gate.in1 == previousOverflow && gate.in2 == xorGateName) || (gate.in1 == xorGateName && gate.in2 == previousOverflow)) && gate.operator == "AND" {
			andGateNameWithPreviousOverflow = gate.out
			break
		}
	}
	if andGateNameWithPreviousOverflow == "" {
		fmt.Println("looking for", previousOverflow, xorGateName, "AND")
		panic("andGateNameWithPreviousOverflow gate not found")
	}

	for _, gate := range gates {
		if ((gate.in1 == andGateName && gate.in2 == andGateNameWithPreviousOverflow) || (gate.in2 == andGateName && gate.in1 == andGateNameWithPreviousOverflow)) && gate.operator == "OR" {
			return gate
		}
	}

	// find alternative
	fmt.Println("andGateName", andGateName)
	fmt.Println("xorGateName", xorGateName)
	fmt.Println("andGateNameWithPreviousOverflow", andGateNameWithPreviousOverflow)

	for _, gate := range gates {
		if ((gate.in1 == andGateName) || (gate.in2 == andGateName)) && gate.operator == "OR" {
			fmt.Println("maybe use?")
			fmt.Println(gateToString(gate))
			fmt.Println("andGateName should be there:", andGateName)
			fmt.Println("not ", previousOverflow)
		}
	}

	panic("overflow gate not found")
}

func twoDigitNumber(n int) string {
	s := strconv.Itoa(n)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}

func swapGateOuts(out1 string, out2 string, gates []gate) []gate {
	var newGates []gate
	for _, g := range gates {
		out := g.out
		if out == out1 {
			out = out2
		} else if out == out2 {
			out = out1
		}

		newGates = append(newGates, gate{g.in1, g.in2, out, g.operator, g.executed})
	}

	return newGates
}
