package main

import (
	"log"
	"strconv"
	"strings"
)

func day12_input() []Pipe {
	lines := readFile("./day12_input.txt")
	var pipes []Pipe
	for _, line := range lines {
		pieces := strings.Fields(line)
		pipeName, _ := strconv.Atoi(pieces[0])
		var childInts []int
		if len(pieces) > 2 {
			for i := 2; i < len(pieces); i++ {
				childInt, _ := strconv.Atoi(strings.Trim(pieces[i], ","))
				childInts = append(childInts, childInt)
			}
		}
		pipes = append(pipes, Pipe{name: pipeName, childrenInts: childInts})
	}
	return pipes
}

func day12_part1() {
	pipes := day12_input()
	var connections []int
	allConnections := getConnections(0, pipes, connections)
	log.Print(len(allConnections))
}

func day12_part2() {
	pipes := day12_input()
	var connections []int
	groups := 0
	for _, pipe := range pipes {
		if arrayIntsFind(pipe.name, connections) == -1 {
			connections = getConnections(pipe.name, pipes, connections)
			groups++
		}
	}
	log.Print(groups)
}

func getConnections(name int, pipes []Pipe, connections []int) []int {
	pipe := pipes[findPipeIndex(name, pipes)]
	if len(pipe.childrenInts) > 0 {
		for _, childInt := range pipe.childrenInts {
			index := arrayIntsFind(childInt, connections)
			if index == -1 {
				connections = append(connections, childInt)
				connections = getConnections(childInt, pipes, connections)
			}
		}
	}
	return connections
}

func findPipeIndex(name int, pipes []Pipe) int {
	for i, pipe := range pipes {
		if pipe.name == name {
			return i
		}
	}
	return -1
}

type Pipe struct {
	name         int
	childrenInts []int
	connections  []int
}
