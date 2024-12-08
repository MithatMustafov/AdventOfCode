package main

import (
	fileparser "aoc/utils"
	"fmt"
)

const EMPTY_TARGET = rune('.')

type Coordinates struct {
	X int
	Y int
}
type ResonantCollinearity struct {
	cityMap   []string
	antennas  map[rune][]Coordinates
	antinodes []Coordinates
}

func (rC *ResonantCollinearity) scanMapForAntennas() {
	rC.antennas = make(map[rune][]Coordinates)

	for y, antennas := range rC.cityMap {
		for x, frequency := range antennas {
			if frequency != EMPTY_TARGET {
				newAntenna := Coordinates{X: x, Y: y}
				rC.antennas[frequency] = append(rC.antennas[frequency], newAntenna)
			}
		}
	}
}

func (rC *ResonantCollinearity) findAntinode(newAntinode Coordinates) bool {
	for _, antinode := range rC.antinodes {
		if newAntinode == antinode {
			return true
		}
	}
	return false
}

func (rC *ResonantCollinearity) printMap(withAntinodes bool) {
	var printableCell rune
	for y, antennas := range rC.cityMap {
		for x, cell := range antennas {
			printableCell = cell
			if withAntinodes && rC.findAntinode(Coordinates{X: x, Y: y}) {
				printableCell = '#'
			}

			fmt.Printf("%v", string(printableCell))
		}
	}
}

func (rC *ResonantCollinearity) isValidAntinode(antinode Coordinates) bool {

	if antinode.X < 0 || antinode.X >= len(rC.cityMap[0]) ||
		antinode.Y < 0 || antinode.Y >= len(rC.cityMap) {
		return false
	}

	return true
}

func (rC *ResonantCollinearity) createAntinode(antennaA Coordinates, antennaB Coordinates) (antinode Coordinates, err bool) {
	antinode.X = antennaB.X + (antennaB.X - antennaA.X)
	antinode.Y = antennaB.Y + (antennaB.Y - antennaA.Y)
	if !rC.isValidAntinode(antinode){
		return antinode, false
	}
	
	if rC.findAntinode((antinode)) {
		return antinode, true
	}

	rC.antinodes = append(rC.antinodes, antinode)
	return antinode, true
}

func PartOne(rC ResonantCollinearity) int {
	for _, frequency := range rC.antennas {
		for idxA := 0; idxA < len(frequency)-1; idxA++ {
			for idxB := idxA + 1; idxB < len(frequency); idxB++ {
				rC.createAntinode(frequency[idxA], frequency[idxB])
				rC.createAntinode(frequency[idxB], frequency[idxA])
			}
		}
	}
	return len(rC.antinodes)
}

func PartTwo(rC ResonantCollinearity) int {
	var previousAntinode, nextAntinode, newAntinode Coordinates
	for _, frequency := range rC.antennas {
		for idxA := 0; idxA < len(frequency)-1; idxA++ {
			for idxB := idxA + 1; idxB < len(frequency); idxB++ {
				previousAntinode, nextAntinode = frequency[idxB] , frequency[idxA]
				for err := true ;err; {
					newAntinode, err = rC.createAntinode(previousAntinode, nextAntinode)
					previousAntinode, nextAntinode = nextAntinode, newAntinode
				}
				previousAntinode, nextAntinode = frequency[idxA], frequency[idxB]
				for err := true ;err; {
					newAntinode, err = rC.createAntinode(previousAntinode, nextAntinode)
					previousAntinode , nextAntinode = nextAntinode, newAntinode
				}
			}
		}
	}

	antinodesCount := 0
    for _, coords := range rC.antennas {
        for _, coord := range coords {
            if(!rC.findAntinode(coord)){antinodesCount++}
        }
    }
	antinodesCount+= len(rC.antinodes)

	return antinodesCount
}

func main() {
	fmt.Printf("--- Day 8: Resonant Collinearity ---\n")
	var resonantCollinearity ResonantCollinearity
	resonantCollinearity.cityMap = fileparser.ReadFileLines("input", false)
	resonantCollinearity.scanMapForAntennas()
	fmt.Printf("PART[1] %v \n", PartOne(resonantCollinearity))
	fmt.Printf("PART[2] %v \n", PartTwo(resonantCollinearity))
}
