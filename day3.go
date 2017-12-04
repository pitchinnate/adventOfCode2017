package main

import (
	"fmt"
	"log"
	"math"
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
	numberOfCircles := 5
	maxSideLength := (numberOfCircles * 2) - 1
	numbers := makeGrid(maxSideLength)
	xPos := maxSideLength / 2
	xCenterPos := xPos
	numbers[xPos][xPos] = 1

	numbers[xPos][(xPos + 1)] = 1
	numbers[(xPos - 1)][(xPos + 1)] = 2
	numbers[(xPos - 1)][xPos] = 4
	numbers[(xPos - 1)][(xPos - 1)] = 5
	numbers[xPos][(xPos - 1)] = 10
	numbers[(xPos + 1)][(xPos - 1)] = 11
	numbers[(xPos + 1)][xPos] = 23
	numbers[(xPos + 1)][(xPos + 1)] = 25

	nextY := xPos + 1
	nextX := xPos + 2

	for circle := 3; circle <= numberOfCircles; circle++ {
		currentMin := xCenterPos - circle + 1
		currentMax := xCenterPos + circle - 1
		direction := "new"

		for direction != "done" {
			newVal := getValues(numbers, nextX, nextY)
			numbers[nextY][nextX] = newVal
			if direction == "new" {
				direction = "up"
				nextY--
			} else if direction == "up" {
				if (nextY - 1) >= currentMin {
					nextY--
				} else {
					direction = "left"
					nextX--
				}
			} else if direction == "left" {
				if (nextX - 1) >= currentMin {
					nextX--
				} else {
					direction = "down"
					nextY++
				}
			} else if direction == "down" {
				if (nextY + 1) <= currentMax {
					nextY++
				} else {
					direction = "right"
					nextX++
				}
			} else if direction == "right" {
				if (nextX + 1) <= currentMax {
					nextX++
				} else {
					direction = "done"
					nextX++
				}
			}
		}
	}
	printGrid(numbers)
}

func getValues(numbers [][]int, xCord int, yCord int) int {
	max := len(numbers) - 1
	total := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			newYCord := yCord + y
			newXCord := xCord + x
			if newYCord >= 0 && newXCord >= 0 && newYCord <= max && newXCord <= max {
				total += numbers[newYCord][newXCord]
			}
		}
	}
	return total
}

func makeGrid(sideLength int) [][]int {
	numbers := make([][]int, sideLength)
	for y := 0; y < sideLength; y++ {
		numbers[y] = make([]int, sideLength)
		for x := 0; x < sideLength; x++ {
			numbers[y][x] = 0
		}
	}
	return numbers
}

func printGrid(numbers [][]int) {
	for y := 0; y < len(numbers); y++ {
		for x := 0; x < len(numbers); x++ {
			fmt.Printf("%7d", numbers[y][x])
		}
		fmt.Print("\n")
	}
}
