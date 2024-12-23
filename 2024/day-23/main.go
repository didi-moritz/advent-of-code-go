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
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 23, realData))

	var result string
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

const T = 't'

func part1(data [][]byte) string {
	fullMap := make(map[string][]string)

	var ts []string

	for _, line := range data {
		n1 := string(line[0:2])
		n2 := string(line[3:5])

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

	return strconv.Itoa(len(foundNetworks))
}

func buildNetworkName(nodes ...string) string {
	sort.Strings(nodes)
	return strings.Join(nodes, ",")
}

func part2(data [][]byte) string {
	fullMap := make(map[string][]string)

	for _, line := range data {
		n1 := string(line[0:2])
		n2 := string(line[3:5])

		fullMap[n1] = append(fullMap[n1], n2)
		fullMap[n2] = append(fullMap[n2], n1)
	}

	// sort map
	for _, nodes := range fullMap {
		sort.Strings(nodes)
	}

	biggestNetwork := ""
	biggestNetworkSize := 1

	var checkedNetworks []string

	for node, otherNodes := range fullMap {
		if len(otherNodes)+1 <= biggestNetworkSize {
			continue
		}

		scores := make(map[string][]string)

		for _, otherNode := range otherNodes {
			for _, otherNode2 := range otherNodes {
				if otherNode == otherNode2 {
					continue
				}

				if slices.Contains(fullMap[otherNode], otherNode2) {
					scores[otherNode] = append(scores[otherNode], otherNode2)
				}
			}
		}

		for i := biggestNetworkSize + 1; i < len(otherNodes)+2; i++ {

			if i <= biggestNetworkSize {
				continue
			}

			networks := getPermutations(otherNodes, i-1)

			for _, network := range networks {
				if (len(network) + 1) <= biggestNetworkSize {
					continue
				}

				check := true

				fullNetwork := make([]string, len(network))
				copy(fullNetwork, network)
				fullNetwork = append(fullNetwork, node)
				networkName := buildNetworkName(fullNetwork...)
				if slices.Contains(checkedNetworks, networkName) {
					continue
				}
				checkedNetworks = append(checkedNetworks, networkName)

				for _, checkNode := range network {
					if !check {
						break
					}

					checkOtherNodes := fullMap[checkNode]
					for _, checkOtherNode := range network {
						if checkOtherNode == checkNode {
							continue
						}

						if !slices.Contains(checkOtherNodes, checkOtherNode) {
							check = false
							break
						}
					}
				}

				if check {
					if len(fullNetwork) > biggestNetworkSize {
						biggestNetworkSize = len(fullNetwork)
						biggestNetwork = networkName
						fmt.Println(networkName)
					}
				}
			}
		}
	}

	return biggestNetwork
}

func getPermutations(otherNodes []string, size int) [][]string {
	var result [][]string
	getNextPermutation(otherNodes, make([]string, 0), size, &result)
	return result
}

func getNextPermutation(remainingOtherNodes []string, selectedNodes []string, size int, result *[][]string) {
	if len(remainingOtherNodes) < 1 {
		return
	}

	newSelectedNodes := make([]string, len(selectedNodes))
	copy(newSelectedNodes, selectedNodes)
	newSelectedNodes = append(newSelectedNodes, remainingOtherNodes[0])

	if len(newSelectedNodes) == size {
		*result = append(*result, newSelectedNodes)
	}

	if len(remainingOtherNodes) > 0 {
		getNextPermutation(remainingOtherNodes[1:], newSelectedNodes, size, result)
		getNextPermutation(remainingOtherNodes[1:], selectedNodes, size, result)
	}
}
