package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func day4_part2() {
	lines := getDay4Input()
	goodLines := 0
	for _, v := range lines {
		pieces := strings.Fields(v)
		found := false
		for x := 0; x < (len(pieces) - 1); x++ {
			found = containsAnagram(pieces[(x+1):len(pieces)], pieces[x])
			if found {
				break
			}
		}
		if !found {
			goodLines++
		}
	}
	fmt.Printf("Good Lines: %d", goodLines)
}

func day4_part1() {
	lines := getDay4Input()
	goodLines := 0
	for _, v := range lines {
		pieces := strings.Fields(v)
		found := false
		for x := 0; x < (len(pieces) - 1); x++ {
			found = contains(pieces[(x+1):len(pieces)], pieces[x])
			if found {
				break
			}
		}
		if !found {
			goodLines++
		}
	}
	fmt.Printf("Good Lines: %d", goodLines)
}

func getDay4Input() []string {
	var lines []string
	file, err := os.Open("./day4_input.txt")
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

func containsAnagram(s []string, e string) bool {
	newE := SortString(e)
	for _, a := range s {
		newA := SortString(a)
		if newE == newA {
			return true
		}
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
