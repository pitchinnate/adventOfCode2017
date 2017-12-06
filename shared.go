package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func stringOfIntsToInts(input string) []int {
	strings := strings.Fields(input)
	var numbers []int
	for _, r := range strings {
		currentInt, _ := strconv.Atoi(r)
		numbers = append(numbers, currentInt)
	}
	return numbers
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MinMaxIndex(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	var maxIndex int = 0
	var minIndex int = 0
	for i, value := range array {
		if max < value {
			max = value
			maxIndex = i
		}
		if min > value {
			min = value
			minIndex = i
		}
	}
	return minIndex, maxIndex
}

func intsToString(ints []int, separator string) string {
	var stringsInts []string
	for _, v := range ints {
		stringsInts = append(stringsInts, strconv.Itoa(v))
	}
	return strings.Join(stringsInts, separator)
}

func arrayHasString(strings []string, search string) int {
	for i, v := range strings {
		if v == search {
			return i
		}
	}
	return -1
}

func arrayHasStringCount(strings []string, search string) int {
	count := 0
	for _, v := range strings {
		if v == search {
			count++
		}
	}
	return count
}
