package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day14_input() [128]string {
	//input := "flqrgnkx"
	input := "hwlqcszp"
	var rows [128]string
	for i := 0; i < 128; i++ {
		rows[i] = input
		rows[i] += "-"
		rows[i] += strconv.Itoa(i)
	}
	for i := 0; i < 128; i++ {
		rows[i] = day14_hash(rows[i])
	}
	return rows
}

func day14_part1() {
	rows := day14_input()
	fullString := ""
	for _, row := range rows {
		letters := strings.Split(row, "")
		for _, letter := range letters {
			value, _ := strconv.ParseUint(letter, 16, 32)
			fullString += fmt.Sprintf("%4s", strconv.FormatInt(int64(value), 2))
		}
	}
	numberOfOnes := 0
	allLetters := strings.Split(fullString, "")
	for _, letter := range allLetters {
		if letter == "1" {
			numberOfOnes++
		}
	}
	log.Print(numberOfOnes)
}

func day14_part2() {
	rows := day14_input()
	fullString := ""
	for _, row := range rows {
		letters := strings.Split(row, "")
		for _, letter := range letters {
			value, _ := strconv.ParseUint(letter, 16, 32)
			fullString += fmt.Sprintf("%04s", strconv.FormatInt(int64(value), 2))
		}
		fullString += "\n"
	}
	fmt.Print(fullString)
}

func day14_input2(input string) ([256]int, []int) {
	var numbers [256]int
	for i, _ := range numbers {
		numbers[i] = i
	}
	var lengths []int
	characters := []rune(input)
	for _, character := range characters {
		lengths = append(lengths, int(character))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
	return numbers, lengths
}

func day14_hash(input string) string {
	numbers, lengths := day14_input2(input)
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
	returnString := ""
	for _, v := range hashValues {
		tempString := fmt.Sprintf("%02x", v)
		returnString += tempString
	}
	return returnString
}
