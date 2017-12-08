package main

import (
	"log"
	"strconv"
	"strings"
)

func day8_input() ([]Register, []Instruction) {
	var registers []Register
	var instructions []Instruction
	lines := readFile("./day8_input.txt")
	for _, line := range lines {
		inputs := strings.Fields(line)
		if getRegisterIndexByName(inputs[0], registers) == -1 {
			registers = append(registers, Register{name: inputs[0], value: 0})
		}
		adjustment, _ := strconv.Atoi(inputs[2])
		if inputs[1] == "dec" {
			adjustment = adjustment * -1
		}
		checkInt, _ := strconv.Atoi(inputs[6])
		instructions = append(instructions, Instruction{register: inputs[0], adjust: adjustment, checkRegister: inputs[4], checkOperator: inputs[5], checkValue: checkInt})
	}
	return registers, instructions
}

func day8_part1() {
	registers, instructions := day8_input()
	for _, instruction := range instructions {
		checkRegisterIndex := getRegisterIndexByName(instruction.checkRegister, registers)
		checkRegister := registers[checkRegisterIndex]
		adjustRegisterIndex := getRegisterIndexByName(instruction.register, registers)

		if instruction.checkOperator == "==" {
			if checkRegister.value == instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		} else if instruction.checkOperator == ">" {
			if checkRegister.value > instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		} else if instruction.checkOperator == ">=" {
			if checkRegister.value >= instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		} else if instruction.checkOperator == "<=" {
			if checkRegister.value <= instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		} else if instruction.checkOperator == "<" {
			if checkRegister.value < instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		} else if instruction.checkOperator == "!=" {
			if checkRegister.value != instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
			}
		}
	}
	log.Print(registers)
}

func day8_part2() {
	registers, instructions := day8_input()
	highestVal := 0
	for _, instruction := range instructions {
		checkRegisterIndex := getRegisterIndexByName(instruction.checkRegister, registers)
		checkRegister := registers[checkRegisterIndex]
		adjustRegisterIndex := getRegisterIndexByName(instruction.register, registers)

		if instruction.checkOperator == "==" {
			if checkRegister.value == instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		} else if instruction.checkOperator == ">" {
			if checkRegister.value > instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		} else if instruction.checkOperator == ">=" {
			if checkRegister.value >= instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		} else if instruction.checkOperator == "<=" {
			if checkRegister.value <= instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		} else if instruction.checkOperator == "<" {
			if checkRegister.value < instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		} else if instruction.checkOperator == "!=" {
			if checkRegister.value != instruction.checkValue {
				registers[adjustRegisterIndex].value += instruction.adjust
				highestVal = adjustHighest(highestVal, registers[adjustRegisterIndex].value)
			}
		}
	}
	log.Print(highestVal)
}

func adjustHighest(value int, newvalue int) int {
	if newvalue > value {
		return newvalue
	}
	return value
}

func getRegisterIndexByName(name string, registers []Register) int {
	for i, register := range registers {
		if register.name == name {
			return i
		}
	}
	return -1
}

type Register struct {
	name  string
	value int
}

type Instruction struct {
	register      string
	adjust        int
	checkRegister string
	checkOperator string
	checkValue    int
}
