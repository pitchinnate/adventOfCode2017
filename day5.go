package main

import (
	"fmt"
	"strconv"
)

func day5_input() []int {
	input := readFile("./day5_input.txt")
	//input := readFile("./day5_input_test.txt")
	var numbers []int
	for _, v := range input {
		currentInt, _ := strconv.Atoi(v)
		numbers = append(numbers,currentInt)
	}
	return numbers
}

func day5_part1() {
	numbers := day5_input()
	position := 0
	steps := 0

	for position < len(numbers) {
		//fmt.Printf("at positin %d \n",position)
		steps++
		changePosition := numbers[position]
		numbers[position]++
		position += changePosition
	}

	fmt.Printf("took %d steps",steps)
}

func day5_part2() {
	numbers := day5_input()
	position := 0
	steps := 0

	for position < len(numbers) {
		//fmt.Printf("at positin %d \n",position)
		steps++
		changePosition := numbers[position]
		if changePosition >= 3 {
			numbers[position]--
		} else {
			numbers[position]++
		}
		position += changePosition
	}

	fmt.Printf("took %d steps",steps)
}