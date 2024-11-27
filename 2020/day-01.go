package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	data := utils.ReadFileAsStringArray("2020/day-01.data")

	// Convert string array to int array
	var intData []int64
	for _, v := range data {
		intData = append(intData, utils.ParseInt(v))
	}

	for i := 0; i < len(intData); i++ {
		for j := i + 1; j < len(intData); j++ {
			if intData[i]+intData[j] == 2020 {
				fmt.Print(intData[i] * intData[j])
			}
		}
	}
}
