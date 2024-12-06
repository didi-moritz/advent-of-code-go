package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	part, useRealData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2020, 4, useRealData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data []string) int {
	var lines []string

	result := 0

	for _, line := range data {
		if line == "" {
			result += processLines(lines)
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}

	result += processLines(lines)

	return result
}

func processLines(lines []string) int {
	count := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			if !strings.HasPrefix(word, "cid:") {
				count++
			}
		}
	}

	if count == 7 {
		return 1
	} else {
		return 0
	}
}

func part2(data []string) int {
	var lines []string

	result := 0

	for _, line := range data {
		if line == "" {
			result += processLines2(lines)
			lines = nil
		} else {
			lines = append(lines, line)
		}
	}

	result += processLines2(lines)

	return result
}

func processLines2(lines []string) int {
	hclPattern := regexp.MustCompile("^#[0-9a-f]{6}$")
	eclPattern := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	pidPattern := regexp.MustCompile("^\\d{9}$")

	count := 0
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			parts := strings.Split(word, ":")

			if len(parts) != 2 {
				return 0
			}

			key := parts[0]
			value := parts[1]

			if key == "cid" {
				continue
			} else if key == "byr" {
				if !checkNumber(value, 1920, 2002) {
					return 0
				}
			} else if key == "iyr" {
				if !checkNumber(value, 2010, 2020) {
					return 0
				}
			} else if key == "eyr" {
				if !checkNumber(value, 2020, 2030) {
					return 0
				}
			} else if key == "hgt" {
				if strings.HasSuffix(value, "cm") {
					var v int
					fmt.Sscanf(value, "%dcm", &v)
					if !checkNumberInt(v, 150, 193) {
						return 0
					}
				} else if strings.HasSuffix(value, "in") {
					var v int
					fmt.Sscanf(value, "%din", &v)
					if !checkNumberInt(v, 59, 76) {
						return 0
					}
				} else {
					return 0
				}
			} else if key == "hcl" {
				if !hclPattern.Match([]byte(value)) {
					return 0
				}
			} else if key == "ecl" {
				if !eclPattern.Match([]byte(value)) {
					return 0
				}
			} else if key == "pid" {
				if !pidPattern.Match([]byte(value)) {
					return 0
				}
			}

			count++
		}
	}

	if count == 7 {
		return 1
	} else {
		return 0
	}
}

func checkNumber(value string, min int, max int) bool {
	number := int(utils.ParseInt(value))
	return checkNumberInt(number, min, max)
}

func checkNumberInt(number int, min int, max int) bool {
	return number >= min && number <= max
}
