package main

import (
	"fmt"
	"log"
)

// func day10_input() ([5]int, []int) {
// 	var numbers [5]int
// 	for i, _ := range numbers {
// 		numbers[i] = i
// 	}
// 	lengths := []int{3, 4, 1, 5}
// 	return numbers, lengths
// }

func day10_input() ([256]int, []int) {
	var numbers [256]int
	for i, _ := range numbers {
		numbers[i] = i
	}
	lengths := []int{18, 1, 0, 161, 255, 137, 254, 252, 14, 95, 165, 33, 181, 168, 2, 188}
	return numbers, lengths
}

func day10_input2() ([256]int, []int) {
	var numbers [256]int
	for i, _ := range numbers {
		numbers[i] = i
	}
	var lengths []int
	characters := []rune("18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188")
	for _, character := range characters {
		lengths = append(lengths, int(character))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
	return numbers, lengths
}

func day10_part1() {
	numbers, lengths := day10_input()
	skipSize := 0
	pivotIndex := 0
	for _, length := range lengths {
		if length > 1 {
			subset := make([]int, length)
			for index := 0; index < length; index++ {
				newIndex := (pivotIndex + index) % len(numbers)
				subset = append(subset, numbers[newIndex])
			}
			subset = reverseSlice(subset)
			for index := 0; index < length; index++ {
				newIndex := (pivotIndex + index) % len(numbers)
				numbers[newIndex] = subset[index]
			}
		}
		pivotIndex += length + skipSize
		pivotIndex = pivotIndex % len(numbers)
		skipSize++
	}
	log.Print(numbers)
}

func day10_part2() {
	numbers, lengths := day10_input2()
	skipSize := 0
	pivotIndex := 0
	rounds := 0
	for rounds < 64 {
		for _, length := range lengths {
			if length > 1 {
				subset := make([]int, length)
				for index := 0; index < length; index++ {
					newIndex := (pivotIndex + index) % len(numbers)
					subset = append(subset, numbers[newIndex])
				}
				subset = reverseSlice(subset)
				for index := 0; index < length; index++ {
					newIndex := (pivotIndex + index) % len(numbers)
					numbers[newIndex] = subset[index]
				}
			}
			pivotIndex += length + skipSize
			pivotIndex = pivotIndex % len(numbers)
			skipSize++
		}
		rounds++
	}
	var hashValues []int
	for x := 0; x < 16; x++ {
		startIndex := x * 16
		total := numbers[startIndex]
		for y := 1; y < 16; y++ {
			index := startIndex + y
			total = total ^ numbers[index]
		}
		hashValues = append(hashValues, total)
	}
	for _, v := range hashValues {
		fmt.Printf("%x \n", v)
	}
	fmt.Printf("\n")
	log.Print(hashValues)
}
