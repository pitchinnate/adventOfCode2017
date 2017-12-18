package main

import (
	"fmt"
	"log"
	"strconv"
)

func day15_part1() {
	genA := 116
	genB := 299
	// genA := 65
	// genB := 8921
	runtimes := 40000000
	matches := 0

	for i := 0; i < runtimes; i++ {
		genA = (genA * 16807) % 2147483647
		genB = (genB * 48271) % 2147483647
		// log.Printf("A: %10d B: %10d", genA, genB)
		genABinary := fmt.Sprintf("%032s", strconv.FormatInt(int64(genA), 2))
		genBBinary := fmt.Sprintf("%032s", strconv.FormatInt(int64(genB), 2))
		// fmt.Printf("A: %s \nB: %s\n", genABinary, genBBinary)
		// fmt.Printf("A: %s \nB: %s\n", genABinary[16:], genBBinary[16:])
		if genABinary[16:] == genBBinary[16:] {
			matches++
		}
	}
	log.Print(matches)
}

func day15_part2() {
	genA := 116
	genB := 299
	// genA := 65
	// genB := 8921
	runtimes := 5000000
	matches := 0

	for i := 0; i < runtimes; i++ {
		newGenA := true
		for newGenA {
			genA = (genA * 16807) % 2147483647
			if genA%4 == 0 {
				newGenA = false
			}
		}
		newGenB := true
		for newGenB {
			genB = (genB * 48271) % 2147483647
			if genB%8 == 0 {
				newGenB = false
			}
		}
		// log.Printf("A: %10d B: %10d", genA, genB)
		genABinary := fmt.Sprintf("%032s", strconv.FormatInt(int64(genA), 2))
		genBBinary := fmt.Sprintf("%032s", strconv.FormatInt(int64(genB), 2))
		// fmt.Printf("A: %s \nB: %s\n", genABinary, genBBinary)
		// fmt.Printf("A: %s \nB: %s\n", genABinary[16:], genBBinary[16:])
		if genABinary[16:] == genBBinary[16:] {
			matches++
		}
	}
	log.Print(matches)
}
