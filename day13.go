package main

import (
	"log"
	"strconv"
	"strings"
)

func day13_input() [99]Layer {
	lines := readFile("./day13_input.txt")
	var layers [99]Layer
	for _, line := range lines {
		pieces := strings.Fields(line)
		index, _ := strconv.Atoi(pieces[0])
		size, _ := strconv.Atoi(pieces[1])
		layers[index].currentPosition = 0
		layers[index].size = size
		layers[index].direction = 1
	}
	return layers
}

func day13_part1() {
	layers := day13_input()
	position := 0
	severity := 0
	for position <= 98 {
		if layers[position].size > 0 && layers[position].currentPosition == 0 {
			severity += (position * layers[position].size)
		}
		layers = updateAllLayers(layers)
		position++
	}
	log.Print(severity)
}

func day13_part2() {
	layers := day13_input()
	delay := 0
	position := 0
	severity := 0
	search := true
	for search {
		if delay > 0 {
			layers = updateAllLayers(layers)
		}
		position = 0
		severity = 0
		copyLayers := layers
		for position <= 98 {
			if copyLayers[position].size > 0 && copyLayers[position].currentPosition == 0 {
				severity = 1
				break
			}
			copyLayers = updateAllLayers(copyLayers)
			position++
		}
		if severity == 0 {
			search = false
		} else {
			delay++
		}
	}

	log.Print(delay)
}

func resetLayers(layers [99]Layer) [99]Layer {
	for i, layer := range layers {
		if layer.size > 0 {
			layers[i].currentPosition = 0
			layers[i].direction = 1
		}
	}
	return layers
}

func updateAllLayers(layers [99]Layer) [99]Layer {
	for i, layer := range layers {
		if layer.size > 0 {
			direction := layer.direction
			if direction == 1 && layer.currentPosition >= (layer.size-1) {
				direction = -1
			} else if direction == -1 && layer.currentPosition == 0 {
				direction = 1
			}
			nextPostion := layer.currentPosition + direction
			layers[i].currentPosition = nextPostion
			layers[i].direction = direction
		}
	}
	return layers
}

type Layer struct {
	currentPosition int
	size            int
	direction       int
}
