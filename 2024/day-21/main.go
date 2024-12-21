package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsByteArray(utils.GetFileName(2024, 21, realData))

	var result int
	if part == 1 {
		result = part1(data)
		if realData {
			if result >= 171492 || result == 170508 {
				// 171492
				// 172476
				// 179200
				// 170508
				fmt.Println("must be lower than 171492 and not 170508")
			}
		}
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

func part1(data [][]byte) int {
	directionalKeys[UP] = v{1, 0}
	directionalKeys[DOWN] = v{1, 1}
	directionalKeys[LEFT] = v{0, 1}
	directionalKeys[RIGHT] = v{2, 1}
	directionalKeys['A'] = v{2, 0}

	result := 0
	for _, code := range data {
		minimum := math.MaxInt
		for i := 0; i < 10000; i++ {
			newResult := calcCode(code)
			if newResult < minimum {
				minimum = newResult
			}
		}
		result += minimum
	}
	return result
}

const A int = -1
const UP byte = '^'
const DOWN byte = 'v'
const LEFT byte = '<'
const RIGHT byte = '>'

type v struct {
	x, y int
}

var (
	directionalKeys = make(map[byte]v)
)

func calcCode(code []byte) int {
	var newCode []byte
	{
		currentPos := A

		for _, c := range code {
			var to int
			if c == 'A' {
				to = A
			} else {
				to, _ = strconv.Atoi(string(c))
			}
			newCode = append(newCode, calcNumericFromTo(currentPos, to)...)
			currentPos = to
		}
	}

	for i := 0; i < 2; i++ {
		var newCode2 []byte
		{
			currentPos := byte('A')

			for _, to := range newCode {
				newCode2 = append(newCode2, calcDirectionalFromTo(currentPos, to)...)
				currentPos = to
			}
		}

		newCode = newCode2
	}

	number, _ := strconv.Atoi(string(code[:len(code)-1]))

	return number * len(newCode)
}

func calcDirectionalFromTo(from byte, to byte) []byte {
	var code []byte
	fromPos := directionalKeys[from]
	toPos := directionalKeys[to]

	fromX := fromPos.x
	fromY := fromPos.y

	verticalFirst := rand.Intn(2) < 1

	if (fromPos == v{0, 1}) && fromY != toPos.y {
		verticalFirst = false
	} else if (toPos == v{0, 1}) && fromY != toPos.y {
		verticalFirst = true
	}

	if verticalFirst {
		if fromY > toPos.y {
			code = append(code, UP)
		} else if fromY < toPos.y {
			code = append(code, DOWN)
		}

		for fromX != toPos.x {
			if fromX > toPos.x {
				code = append(code, LEFT)
				fromX--
			} else if fromX < toPos.x {
				code = append(code, RIGHT)
				fromX++
			}
		}
	} else {
		for fromX != toPos.x {
			if fromX > toPos.x {
				code = append(code, LEFT)
				fromX--
			} else if fromX < toPos.x {
				code = append(code, RIGHT)
				fromX++
			}
		}

		if fromY > toPos.y {
			code = append(code, UP)
		} else if fromY < toPos.y {
			code = append(code, DOWN)
		}
	}

	return append(code, 'A')
}

func calcNumericFromTo(from int, to int) []byte {
	var code []byte

	verticalFirst := rand.Intn(2) < 1

	current := from

	if rowOfNumeric(current) == 3 && colOfNumeric(to) == 0 {
		verticalFirst = true
	} else if (rowOfNumeric(to) == 3) && colOfNumeric(from) == 0 {
		verticalFirst = false
	}

	if verticalFirst {
		rFrom := rowOfNumeric(current)
		rTo := rowOfNumeric(to)

		if rTo > rFrom {
			for rTo != rFrom {
				code = append(code, DOWN)
				rFrom++
			}
		} else {
			for rTo != rFrom {
				code = append(code, UP)
				rFrom--
			}
		}

		cFrom := colOfNumeric(current)
		cTo := colOfNumeric(to)

		if cTo > cFrom {
			for cTo != cFrom {
				code = append(code, RIGHT)
				cFrom++
			}
		} else {
			for cTo != cFrom {
				code = append(code, LEFT)
				cFrom--
			}
		}
	} else {
		cFrom := colOfNumeric(current)
		cTo := colOfNumeric(to)

		if cTo > cFrom {
			for cTo != cFrom {
				code = append(code, RIGHT)
				cFrom++
			}
		} else {
			for cTo != cFrom {
				code = append(code, LEFT)
				cFrom--
			}
		}

		rFrom := rowOfNumeric(current)
		rTo := rowOfNumeric(to)

		if rTo > rFrom {
			for rTo != rFrom {
				code = append(code, DOWN)
				rFrom++
			}
		} else {
			for rTo != rFrom {
				code = append(code, UP)
				rFrom--
			}
		}
	}

	return append(code, 'A')
}

func rowOfNumeric(a int) int {
	if a == A || a == 0 {
		return 3
	}

	if a >= 7 {
		return 0
	}

	if a >= 4 {
		return 1
	}

	return 2
}

func colOfNumeric(a int) int {
	if a == A {
		return 2
	}

	if a == 0 {
		return 1
	}

	return (a - 1) % 3
}

func part2(data [][]byte) int {
	return 0
}
