package main

import (
	"advent-of-code-go/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func main() {
	part, realData := utils.GetRunConfig(1, false)

	data := utils.ReadFileAsStringArray(utils.GetFileName(2016, 11, realData))

	var result int
	if part == 1 {
		result = part1(data)
	} else {
		result = part2(data)
	}

	fmt.Println(result)
}

type Kind int

const (
	MICROCHIP Kind = iota
	GENERATOR
)

type Thing struct {
	kind    Kind
	floor   int
	element string
}

type CacheValue struct {
	steps  int
	result int
}

func part1(data []string) int {

	var things []Thing

	microchipRegex := regexp.MustCompile("([^ ]+)-compatible microchip")
	generatorRegex := regexp.MustCompile("([^ ]+) generator")

	for i, line := range data {
		microchipFinds := microchipRegex.FindAllStringSubmatch(line, -1)
		for _, microchipFind := range microchipFinds {
			things = append(things, Thing{MICROCHIP, i, microchipFind[1]})
		}
		generatorFinds := generatorRegex.FindAllStringSubmatch(line, -1)
		for _, generatorFind := range generatorFinds {
			things = append(things, Thing{GENERATOR, i, generatorFind[1]})
		}
	}

	cache := make(map[string]CacheValue)
	finalFloor := len(data) - 1

	currentMinimum := math.MaxInt

	return move(things, 0, 0, cache, finalFloor, &currentMinimum)
}

func move(things []Thing, floor int, steps int, cache map[string]CacheValue, finalFloor int, currentMinimum *int) int {
	if steps >= *currentMinimum {
		return math.MaxInt
	}

	cacheKey := calcCacheKey(things, floor)

	cacheValue, found := cache[cacheKey]

	if found && cacheValue.steps <= steps {
		return math.MaxInt
	}

	if !found {
		cacheValue = CacheValue{steps, math.MaxInt}
		cache[cacheKey] = cacheValue
	} else {
		if cacheValue.result < math.MaxInt {
			return cacheValue.result - cacheValue.steps + steps
		}

		cacheValue.steps = steps
		cache[cacheKey] = cacheValue
	}

	if !checkThings(things) {
		return math.MaxInt
	}

	// Check if end is found
	if floor == finalFloor {
		checkFinalFloor := true
		for _, thing := range things {
			if thing.floor != finalFloor {
				checkFinalFloor = false
				break
			}
		}

		if checkFinalFloor {
			if steps < *currentMinimum {
				*currentMinimum = steps
			}

			fmt.Println(steps)

			return steps
		}
	}

	diffs := []int{1, -1}

	minimum := math.MaxInt

	for firstThingIndex, firstThing := range things {
		if firstThing.floor != floor {
			continue
		}

		for _, diff := range diffs {
			newFloor := floor + diff
			if newFloor < 0 || newFloor > finalFloor {
				continue
			}

			{
				newThings := make([]Thing, len(things))
				copy(newThings, things)
				newThings[firstThingIndex].floor += diff
				newSteps := move(newThings, newFloor, steps+1, cache, finalFloor, currentMinimum)
				if newSteps < minimum {
					minimum = newSteps
				}
			}

			for secondThingIndex, secondThing := range things {
				if secondThing.floor != floor || secondThingIndex == firstThingIndex {
					continue
				}

				newThings := make([]Thing, len(things))
				copy(newThings, things)
				newThings[firstThingIndex].floor += diff
				newThings[secondThingIndex].floor += diff
				newSteps := move(newThings, newFloor, steps+1, cache, finalFloor, currentMinimum)
				if newSteps < minimum {
					minimum = newSteps
				}
			}
		}
	}

	cacheValue = cache[cacheKey]
	if minimum < cacheValue.result {
		cacheValue.result = minimum
		cache[cacheKey] = cacheValue
	}

	return minimum
}

func checkThings(things []Thing) bool {
	for _, microchip := range things {
		if microchip.kind == GENERATOR {
			continue
		}

		check := true
		for _, generator := range things {
			if generator.kind == MICROCHIP || generator.floor != microchip.floor {
				continue
			}

			if generator.element == microchip.element {
				check = true
				break
			} else {
				check = false
			}
		}

		if !check {
			return false
		}
	}

	return true
}

func calcCacheKey(things []Thing, floor int) string {
	key := strconv.Itoa(floor)
	for _, thing := range things {
		key += strconv.Itoa(thing.floor)
	}

	return key
}

func part2(data []string) int {
	return 0
}
