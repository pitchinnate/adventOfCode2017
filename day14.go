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
	var lines []string
	for _, row := range rows {
		tempString := ""
		letters := strings.Split(row, "")
		for _, letter := range letters {
			value, _ := strconv.ParseUint(letter, 16, 32)
			tempString += fmt.Sprintf("%04s", strconv.FormatInt(int64(value), 2))
		}
		lines = append(lines, tempString)
	}
	var grid [128][128]int
	for y := 0; y < 128; y++ {
		letters := strings.Split(lines[y], "")
		for x := 0; x < 128; x++ {
			tempInt, _ := strconv.Atoi(letters[x])
			grid[y][x] = tempInt
		}
	}
	//http://gregtrowbridge.com/a-basic-pathfinding-algorithm/
	group_count := 0
	//displayGrid(grid)
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if grid[y][x] == 1 {
				group_count++
				grid = findPath(x, y, grid)
			}
		}
	}
	log.Print(group_count)
}

func findPath(xCord int, yCord int, grid [128][128]int) [128][128]int {
	var queue []Point
	queue = append(queue, Point{xCord: xCord, yCord: yCord})
	search := true
	for search {
		nextCord := queue[0]
		queue = append(queue[:0], queue[1:]...)
		grid[nextCord.yCord][nextCord.xCord] = 0

		if nextCord.yCord > 0 && grid[(nextCord.yCord - 1)][nextCord.xCord] == 1 {
			queue = append(queue, Point{xCord: nextCord.xCord, yCord: (nextCord.yCord - 1)})
		}
		if nextCord.yCord < 127 && grid[(nextCord.yCord + 1)][nextCord.xCord] == 1 {
			queue = append(queue, Point{xCord: nextCord.xCord, yCord: (nextCord.yCord + 1)})
		}
		if nextCord.xCord > 0 && grid[nextCord.yCord][(nextCord.xCord-1)] == 1 {
			queue = append(queue, Point{xCord: (nextCord.xCord - 1), yCord: nextCord.yCord})
		}
		if nextCord.xCord < 127 && grid[nextCord.yCord][(nextCord.xCord+1)] == 1 {
			queue = append(queue, Point{xCord: (nextCord.xCord + 1), yCord: nextCord.yCord})
		}
		if len(queue) == 0 {
			search = false
		}
	}
	return grid
}

func displayGrid(grid [128][128]int) {
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			fmt.Printf("%d", grid[y][x])
		}
		fmt.Print("\n")
	}
}

type Point struct {
	xCord int
	yCord int
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
