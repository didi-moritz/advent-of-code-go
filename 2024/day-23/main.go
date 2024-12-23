package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 23, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

const T = 't'

func part1(data [][]byte) int {
	map1 := make(map[string][]string)
	map2 := make(map[string][]string)
	fullMap := make(map[string][]string)

	var ts []string

	for _, line := range data {
		n1 := string(line[0:2])
		n2 := string(line[3:5])

		map1[n1] = append(map1[n1], n2)
		map2[n2] = append(map1[n2], n1)

		fullMap[n1] = append(fullMap[n1], n2)
		fullMap[n2] = append(fullMap[n2], n1)

		if n1[0] == T {
			if !slices.Contains(ts, n1) {
				ts = append(ts, n1)
			}
		}
		if n2[0] == T {
			if !slices.Contains(ts, n2) {
				ts = append(ts, n2)
			}
		}
	}

	var foundNetworks []string

	for _, t := range ts {
		otherNodes := fullMap[t]
		for n1 := 0; n1 < len(otherNodes)-1; n1++ {
			for n2 := 1; n2 < len(otherNodes); n2++ {
				node1 := otherNodes[n1]
				node2 := otherNodes[n2]
				if slices.Contains(fullMap[node1], node2) {
					networkName := buildNetworkName(t, node1, node2)
					if !slices.Contains(foundNetworks, networkName) {
						foundNetworks = append(foundNetworks, networkName)
					}
				}
			}
		}
	}

	return len(foundNetworks)
}

func buildNetworkName(nodes ...string) string {
	sort.Strings(nodes)
	return strings.Join(nodes, "")
}

func part2(data [][]byte) int {
	return 0
}
