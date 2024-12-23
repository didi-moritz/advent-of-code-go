package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(2, false)

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
		if realData {
			if result >= 373249646668954 {
				// 373249646668954
				fmt.Println("must be lower than 373249646668954")
			}
		}
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
		for i := 0; i < 100; i++ {
			newResult := calcCode(code, 2)
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

func calcCode(code []byte, directionalRobots int) int {
	var newCode []byte
	{
		var newCode2 []byte
		currentPos := A

		for _, c := range code {
			var to int
			if c == 'A' {
				to = A
			} else {
				to, _ = strconv.Atoi(string(c))
			}
			newCode2 = calcNumericFromTo(currentPos, to)
			newCode = append(newCode, newCode2...)
			currentPos = to
		}
	}

	codeToNewCode := make(map[string][]byte)
	codeToScore := make(map[string]int)

	score := calcCodeStep(newCode, 0, directionalRobots, codeToNewCode, codeToScore)
	number, _ := strconv.Atoi(string(code[:len(code)-1]))

	return number * score
}

func calcCodeStep(code []byte, robot int, directionalRobots int, codeToNewCode map[string][]byte, codeToScore map[string]int) int {
	if robot == directionalRobots {
		return len(code)
	}

	var newCode []byte
	var found bool

	newCode, found = codeToNewCode[string(code)]

	if !found {
		var newCode2 []byte
		{
			currentPos := byte('A')
			lastVerticalFirst := true

			for _, to := range code {
				lastVerticalFirst = rand.Intn(2) < 1
				newCode2, lastVerticalFirst = calcDirectionalFromTo(currentPos, to, lastVerticalFirst)
				newCode = append(newCode, newCode2...)
				currentPos = to
			}
		}

		codeToNewCode[string(code)] = newCode
	}

	result := 0

	i := 0
	finished := false
	for !finished {
		j := i
		for ; j < len(newCode) && newCode[j] != 'A'; j++ {
		}

		if j == len(newCode) {
			finished = true
			break
		}

		newCode3 := newCode[i : j+1]
		scoreKey := string(newCode3) + strconv.Itoa(robot+1)

		var newScore int
		var found bool

		newScore, found = codeToScore[scoreKey]
		if !found {
			newScore = calcCodeStep(newCode3, robot+1, directionalRobots, codeToNewCode, codeToScore)
			codeToScore[scoreKey] = newScore
		}

		result += newScore

		i = j + 1

		if j == len(newCode)-1 {
			finished = true
		}
	}

	return result
}

func calcDirectionalFromTo(from byte, to byte, lastVerticalFirst bool) ([]byte, bool) {
	var code []byte
	fromPos := directionalKeys[from]
	toPos := directionalKeys[to]

	fromX := fromPos.x
	fromY := fromPos.y

	verticalFirst := lastVerticalFirst

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

	return append(code, 'A'), verticalFirst
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
	directionalKeys[UP] = v{1, 0}
	directionalKeys[DOWN] = v{1, 1}
	directionalKeys[LEFT] = v{0, 1}
	directionalKeys[RIGHT] = v{2, 1}
	directionalKeys['A'] = v{2, 0}

	result := math.MaxInt

	minimumScores := make([]int, len(data))

	for i := range minimumScores {
		minimumScores[i] = math.MaxInt
	}

	for range 50 {
		for codeNumber, code := range data {
			for i := 0; i < 1000; i++ {
				newResult := calcCode(code, 25)
				if newResult < minimumScores[codeNumber] {
					minimumScores[codeNumber] = newResult
				}
			}
		}

		sum := 0
		for _, scores := range minimumScores {
			sum += scores
		}

		if sum < result {
			result = sum
		}

		fmt.Println(result, sum)
	}

	return result
}
