package main

import (
	"log"
	"strings"
)

func day9_input() []string {
	chars := strings.Split(readFile("./day9_input.txt")[0], "")
	//chars := strings.Split("<<<<>", "")
	return chars
}

func day9_part1() {
	chars := day9_input()
	numberOpened := 0
	escaped := false
	total := 0
	inGarbage := false
	for _, char := range chars {
		if escaped {
			escaped = false
			continue
		} else if char == "!" {
			escaped = true
		} else if char == ">" {
			inGarbage = false
		} else if char == "<" {
			inGarbage = true
		} else if inGarbage {
			continue
		} else if char == "{" {
			numberOpened++
			total += numberOpened
		} else if char == "}" {
			numberOpened--
		}
	}
	log.Print(total)
}

func day9_part2() {
	chars := day9_input()
	numberOpened := 0
	escaped := false
	total := 0
	totalGarbage := 0
	inGarbage := false
	for _, char := range chars {
		if escaped {
			escaped = false
			continue
		} else if char == "!" {
			escaped = true
		} else if char == ">" {
			inGarbage = false
		} else if char == "<" && inGarbage == false {
			inGarbage = true
		} else if inGarbage {
			totalGarbage++
			continue
		} else if char == "{" {
			numberOpened++
			total += numberOpened
		} else if char == "}" {
			numberOpened--
		}
	}
	log.Print(totalGarbage)
}
