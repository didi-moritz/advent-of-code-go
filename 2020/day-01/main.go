package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 1, true))

	// Convert string array to int array
	var intData []int64
	for _, v := range data {
		intData = append(intData, utils.ParseInt(v))
	}

	part2(intData)

}

func part1(intData []int64) {
	for i := 0; i < len(intData); i++ {
		for j := i + 1; j < len(intData); j++ {
			if intData[i]+intData[j] == 2020 {
				fmt.Print(intData[i] * intData[j])
			}
		}
	}
}

func part2(intData []int64) {
	for i := 0; i < len(intData); i++ {
		for j := i + 1; j < len(intData); j++ {
			for k := j + 1; k < len(intData); k++ {
				if intData[i]+intData[j]+intData[k] == 2020 {
					fmt.Print(intData[i] * intData[j] * intData[k])
				}
			}
		}
	}
}
