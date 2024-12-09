package main

import (
	"advent-of-code-go/utils"
	"fmt"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 9, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	var n []int

	fileIndex := 0
	isFile := true
	for _, line := range data {
		for _, c := range line {
			for i := 0; i < int(c-'0'); i++ {
				var value int
				if isFile {
					value = fileIndex
				} else {
					value = -1
				}
				n = append(n, value)
			}

			if isFile {
				fileIndex++
			}

			isFile = !isFile
		}
	}

	b := len(n)

	sum := 0
	for a := 0; a < b; a++ {
		if n[a] == -1 {
			nb := b - 1
			for {
				if n[nb] > -1 || nb == a {
					break
				}
				nb--
			}

			if nb > a {
				n[a] = n[nb]
				n[nb] = -1
			}

			b = nb
		}

		if b <= a {
			break
		}

		if n[a] > -1 {
			sum += n[a] * a
		}
	}

	return sum
}

func part2(data [][]byte) int {
	var n []int

	fileIndex := -1
	isFile := true
	for _, line := range data {
		for _, c := range line {
			if isFile {
				fileIndex++
			}

			for i := 0; i < int(c-'0'); i++ {
				var value int
				if isFile {
					value = fileIndex
				} else {
					value = -1
				}
				n = append(n, value)
			}

			isFile = !isFile
		}
	}

	fmt.Println(fileIndex)

	for fi := fileIndex; fi >= 0; fi-- {
		pos, length := findFile(n, fi)

		sa := findNextEmptySpace(n, length, pos)

		if sa > -1 {
			for i := 0; i < length; i++ {
				n[pos+i] = -1
			}

			for i := 0; i < length; i++ {
				n[sa+i] = fi
			}
		}
	}

	sum := 0
	for i, c := range n {
		if c > -1 {
			sum += i * c
			fmt.Print(c % 10)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")

	return sum
}

func findFile(n []int, fileIndex int) (int, int) {
	start := -1
	end := -1
	for i, c := range n {
		if c == fileIndex {
			if start == -1 {
				start = i
			}
		} else {
			if start > -1 {
				end = i
				break
			}
		}
	}

	// special case for last file
	if end == -1 {
		end = len(n)
	}

	return start, end - start
}

func isEmptySpace(n []int, pos int, length int) bool {
	for p := pos; p < pos+length; p++ {
		if n[p] > -1 {
			return false
		}
	}

	return true
}

func findNextEmptySpace(n []int, length int, maxPos int) int {
	for i := 0; i < maxPos; i++ {
		if isEmptySpace(n, i, length) {
			return i
		}
	}

	return -1
}
