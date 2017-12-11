package main

import "strings"
import "log"

func day11_input() []string {
	directions := strings.Split(readFile("./day11_input.txt")[0], ",")
	//directions := strings.Split("se,sw,se,sw,sw", ",")
	return directions
}

func day11_part1() {
	directions := day11_input()
	counts := make(map[string]int)

	for _, direction := range directions {
		//log.Print(direction)
		counts[direction]++
	}
	log.Print(counts)
	log.Print(findDistance(counts))
	log.Print(stepsTotal(counts))
}

func stepsTotal(counts map[string]int) int {
	total := 0
	for _, val := range counts {
		total += val
	}
	return total
}

func findDistance(counts map[string]int) map[string]int {
	if counts["sw"] > counts["ne"] {
		counts["sw"] = counts["sw"] - counts["ne"]
		counts["ne"] = 0
	} else {
		counts["ne"] = counts["ne"] - counts["sw"]
		counts["sw"] = 0
	}
	if counts["s"] > counts["n"] {
		counts["s"] = counts["s"] - counts["n"]
		counts["n"] = 0
	} else {
		counts["n"] = counts["n"] - counts["s"]
		counts["s"] = 0
	}
	if counts["nw"] > counts["se"] {
		counts["nw"] = counts["nw"] - counts["se"]
		counts["se"] = 0
	} else {
		counts["se"] = counts["se"] - counts["nw"]
		counts["nw"] = 0
	}

	if counts["s"] > 0 && counts["nw"] > 0 {
		if counts["s"] > counts["nw"] {
			counts["sw"] += counts["nw"]
			counts["s"] = counts["s"] - counts["nw"]
			counts["nw"] = 0
		} else {
			counts["sw"] += counts["s"]
			counts["s"] = 0
			counts["nw"] = counts["nw"] - counts["s"]
		}
	}

	if counts["s"] > 0 && counts["ne"] > 0 {
		if counts["s"] > counts["ne"] {
			counts["se"] += counts["ne"]
			counts["s"] = counts["s"] - counts["ne"]
			counts["ne"] = 0
		} else {
			counts["se"] += counts["s"]
			counts["s"] = 0
			counts["ne"] = counts["ne"] - counts["s"]
		}
	}

	if counts["n"] > 0 && counts["sw"] > 0 {
		if counts["n"] > counts["sw"] {
			counts["nw"] += counts["sw"]
			counts["n"] = counts["n"] - counts["sw"]
			counts["sw"] = 0
		} else {
			counts["nw"] += counts["n"]
			counts["n"] = 0
			counts["sw"] = counts["sw"] - counts["n"]
		}
	}

	if counts["n"] > 0 && counts["se"] > 0 {
		if counts["n"] > counts["se"] {
			counts["ne"] += counts["se"]
			counts["n"] = counts["n"] - counts["se"]
			counts["se"] = 0
		} else {
			counts["ne"] += counts["n"]
			counts["n"] = 0
			counts["se"] = counts["se"] - counts["n"]
		}
	}

	if counts["sw"] > 0 && counts["se"] > 0 {
		if counts["sw"] > counts["se"] {
			counts["s"] += counts["se"]
			counts["sw"] = counts["sw"] - counts["se"]
			counts["se"] = 0
		} else {
			counts["s"] += counts["sw"]
			counts["sw"] = 0
			counts["se"] = counts["se"] - counts["sw"]
		}
	}

	if counts["nw"] > 0 && counts["ne"] > 0 {
		if counts["nw"] > counts["ne"] {
			counts["n"] += counts["ne"]
			counts["nw"] = counts["nw"] - counts["ne"]
			counts["ne"] = 0
		} else {
			counts["n"] += counts["nw"]
			counts["nw"] = 0
			counts["ne"] = counts["ne"] - counts["nw"]
		}
	}

	return counts
}

func day11_part2() {
	directions := day11_input()
	counts := make(map[string]int)
	countsCopy := make(map[string]int)
	var distances []int

	for _, direction := range directions {
		counts[direction]++
		for k, v := range counts {
			countsCopy[k] = v
		}
		currentCounts := findDistance(countsCopy)
		distances = append(distances, stepsTotal(currentCounts))
	}

	_, max := MinMax(distances)

	log.Print(max)
}
