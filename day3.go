package main

import (
	"math"
	"log"
)

func day3_part1() {
	input := 325489

	//find circles to find how many moves inward
	circles := 1
	currentValue := 1
	sideLength := 1
	for currentValue < input {
		currentValue = currentValue + (circles * 8)
		sideLength += 2
		circles++
	}

	//find distance from middle of side
	notFound := true
	previousCorner := 0
	for notFound {
		previousCorner = currentValue - sideLength + 1
		if previousCorner < input {
			notFound = false
		} else {
			currentValue = previousCorner
		}
	}
	middleOfSide := previousCorner + ((sideLength - 1) / 2)
	diffFromMiddle := int(math.Abs(float64(middleOfSide - input)))

	//total distance from middle
	totalDistance := (circles - 1) + diffFromMiddle

	log.Print("total moves")
	log.Print(totalDistance)
}

func day3_part2() {
	// input := 325489
	//input := 747
	maxSideLength := 5
	totalNumbers := maxSideLength * maxSideLength
	numbers := makeGrid(maxSideLength)
	xPos := (maxSideLength / 2) + 1
	yPos := xPos
	log.Print(xPos)
	log.Print(yPos)
	log.Print(numbers)
	log.Print(totalNumbers)
	numbers[(yPos - 1)][(xPos - 1)] = 1
	log.Print(numbers)

	changeX := 1
	changeY := 0

	for i:=1; i<=totalNumbers; i++ {
		xPos += changeX
		yPos += changeY
	}

}

func makeGrid(sideLength int) [][]int {
	numbers := make([][]int, sideLength)
	for i:=0; i<sideLength; i++ {
		numbers[i] = make([]int, sideLength)
	}
	return numbers
}