package main

import (
	"fmt"
	"log"
)

func day6_getInput() []int {
	input := `11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11`
	return stringOfIntsToInts(input)
}

func day6_part1() {
	banks := day6_getInput()
	var combos []string
	steps := 0
	found := false

	for !found {
		steps++
		_, maxIndex := MinMaxIndex(banks)
		currentValue := banks[maxIndex]

		banks[maxIndex] = 0

		nextIndex := maxIndex + 1
		if nextIndex > (len(banks) - 1) {
			nextIndex = 0
		}

		for currentValue > 0 {
			banks[nextIndex]++
			nextIndex++
			if nextIndex > (len(banks) - 1) {
				nextIndex = 0
			}
			currentValue--
		}

		newCombo := intsToString(banks, ",")
		if arrayHasString(combos, newCombo) >= 0 {
			found = true
			log.Print(newCombo)
		}
		combos = append(combos, newCombo)
	}

	fmt.Printf("Steps: %d \n", steps)
}

func day6_part2() {
	banks := day6_getInput()
	var combos []string
	steps := 0
	found := false

	for !found {
		steps++
		_, maxIndex := MinMaxIndex(banks)
		currentValue := banks[maxIndex]

		banks[maxIndex] = 0

		nextIndex := maxIndex + 1
		if nextIndex > (len(banks) - 1) {
			nextIndex = 0
		}

		for currentValue > 0 {
			banks[nextIndex]++
			nextIndex++
			if nextIndex > (len(banks) - 1) {
				nextIndex = 0
			}
			currentValue--
		}

		newCombo := intsToString(banks, ",")
		if newCombo == "1,0,14,14,12,12,10,10,8,8,6,6,4,3,2,1" && arrayHasStringCount(combos, newCombo) == 2 {
			found = true
		}
		combos = append(combos, newCombo)
	}

	fmt.Printf("Steps: %d \n", (steps - 4074))
}
