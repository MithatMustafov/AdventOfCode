package main

import (
	"aoc/utils/fileparser"
	"fmt"
)

type Schematic struct {
	Pins []int
}

type Schematics  struct {
	Locks []Schematic
	Keys  []Schematic
}

const SCHEMATIC_PIN_WIDTH = 5
const SCHEMATIC_PIN_HIGHT = 7
const SCHEMATIC_LOCK_IDENTIFIER = "#####"
const SCHEMATIC_PIN_PART = '#'

func (schematics *Schematics) parseData(input []string) {
	schematics.Locks = make([]Schematic, 0)
	schematics.Keys = make([]Schematic, 0)

	for idx := 0; idx < len(input); idx+=SCHEMATIC_PIN_HIGHT+1 {
		line := input[idx]
		if line == "" { continue }

		newSchematic := Schematic{ Pins: make([]int, SCHEMATIC_PIN_WIDTH),}
		for i := range newSchematic.Pins { newSchematic.Pins[i] = -1 }

		for pinIdx := 0; pinIdx < SCHEMATIC_PIN_HIGHT; pinIdx++ {
			currentLine := input[pinIdx+idx]
			for idxPin, valuePin := range currentLine {
				if valuePin == SCHEMATIC_PIN_PART {
					newSchematic.Pins[idxPin]++
				}
			}
		}

		if  line == SCHEMATIC_LOCK_IDENTIFIER { schematics.Locks = append(schematics.Locks, newSchematic)
		} else { schematics.Keys = append(schematics.Keys, newSchematic) }
	}
}

func(schematics Schematics) isPinAligned(pinLock int, pinKey int) (bool) {
	return SCHEMATIC_PIN_WIDTH - pinLock - pinKey >= 0
}

func PartOne(schematics Schematics) int {
	var countFits int = 0
	for _, lock := range schematics.Locks{
		for _, key := range schematics.Keys{
			isFit := true
			for idxPin:= 0; idxPin < SCHEMATIC_PIN_WIDTH; idxPin++ {
				isPinAligned := schematics.isPinAligned(lock.Pins[idxPin], key.Pins[idxPin])
				if !isPinAligned { isFit= false; break }
			}
			if isFit { countFits++ }
		}
	}
	return countFits
}

func PartTwo(schematics Schematics) int { return 0 }

func main() {
	fmt.Printf("--- Day 25: Code Chronicle ---\n")
	var schematics Schematics
	input := fileparser.ReadFileLines("input", false)
	schematics.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(schematics))
	fmt.Printf("PART[2]: %v\n", PartTwo(schematics))
}