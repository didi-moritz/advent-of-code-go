package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func foo(defaultPart int, defaultRealData bool) (int, bool) {
	var part int
	fmt.Print("Part 1 or 2: ")
	_, err := fmt.Scan(&part)
	if err != nil {
		panic(err)
	}

	var realDataString string
	fmt.Print("Real data: ")
	_, err = fmt.Scan(&realDataString)
	if err != nil {
		panic(err)
	}

	var realData bool

	if realDataString == "" {
		realData = defaultRealData
	} else {
		realData = realDataString == "1" || realDataString == "true" || realDataString == "yes" || realDataString == "y"
	}

	return part, realData
}

func main() {
	part, realData := foo(1, true)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 2, realData))

	result := 0
	for _, line := range data {
		if processLine(line, part) {
			result++
		}
	}

	fmt.Println(result)
}

func processLine(line string, part int) bool {
	if part == 1 {
		return processLinePart1(line)
	} else {
		return processLinePart2(line)
	}
}

func processLinePart1(line string) bool {
	var min, max int
	var letter, password string
	fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &letter, &password)

	count := 0
	for _, c := range password {
		if string(c) == letter {
			count++
		}
	}

	return count >= min && count <= max
}

func processLinePart2(line string) bool {
	var pos1, pos2 int
	var letter, password string
	fmt.Sscanf(line, "%d-%d %1s: %s", &pos1, &pos2, &letter, &password)

	match1 := string(password[pos1-1]) == letter
	match2 := string(password[pos2-1]) == letter
	return match1 != match2
}
