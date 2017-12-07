package main

import "strings"

import "strconv"
import "log"
import "fmt"

func day7_input() []Program {
	lines := readFile("./day7_input.txt")
	var programs []Program
	for _, v := range lines {
		pieces := strings.Fields(v)
		weight, _ := strconv.Atoi(strings.Trim(pieces[1], "()"))
		var children []string
		if len(pieces) > 3 {
			children = pieces[3:]
			for i := 0; i < len(children); i++ {
				children[i] = strings.Trim(children[i], ",")
			}
		}
		programs = append(programs, Program{name: pieces[0], weight: weight, children: children, isChild: false, totalWeight: weight})
	}
	return programs
}

func day7_part1() {
	programs := day7_input()
	for _, program := range programs {
		if len(program.children) > 0 && !program.isChild {
			log.Print(program)
			break
		}
	}
}

func day7_part2() {
	programs := day7_input()
	masterProgramIndex := findProgramIndex("vtzay", programs)
	programs[masterProgramIndex].level = 1
	markChildrenLevels(programs[masterProgramIndex], programs, "vtzay")

	for i := 5; i >= 2; i-- {
		for inc, program := range programs {
			if program.level == i {
				total := program.weight
				if len(program.children) > 1 {
					children := getChildren(program, programs)
					for _, child := range children {
						total += child.totalWeight
					}
				}
				programs[inc].totalWeight = total
			}
		}
	}

	for _, program := range programs {
		if len(program.children) > 1 && program.level == 3 {
			children := getChildren(program, programs)
			lastVal := children[0].totalWeight
			for i := 1; i < (len(children) - 1); i++ {
				if lastVal != children[i].totalWeight {
					fmt.Printf("\n Children aren't equal \n")
					log.Print(program)
					log.Print(children)
					break
				}
			}
		}
	}
}

func getChildren(program Program, programs []Program) []Program {
	var chilren []Program
	for _, name := range program.children {
		childIndex := findProgramIndex(name, programs)
		chilren = append(chilren, programs[childIndex])
	}
	return chilren
}

func markChildrenLevels(program Program, programs []Program, parent string) {
	for _, name := range program.children {
		childIndex := findProgramIndex(name, programs)
		programs[childIndex].level = program.level + 1
		programs[childIndex].parent = parent
		if len(programs[childIndex].children) > 0 {
			markChildrenLevels(programs[childIndex], programs, program.name)
		}
	}
}

func findProgramIndex(name string, programs []Program) int {
	for i, program := range programs {
		if program.name == name {
			return i
		}
	}
	return -1
}

type Program struct {
	name        string
	weight      int
	children    []string
	isChild     bool
	totalWeight int
	level       int
	parent      string
}
