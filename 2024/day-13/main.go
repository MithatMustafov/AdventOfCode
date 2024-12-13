package main

import (
	fileparser "aoc/utils"
	"fmt"
	"strconv"
	"regexp"
)

type Coordinates struct {
	x int
	y int
}

type MachineData struct {
	buttonA Coordinates
	buttonB Coordinates
	prize   Coordinates
}

type MachineList struct {
	Data []MachineData
}

func (ml *MachineList) parseMachinesData(inputMap []string) {
	ml.Data = make([]MachineData, 0)

	REGEX_PATTERN_DATA := regexp.MustCompile(`(Button (A|B)|Prize): X(\+|=)(\d+), Y(\+|=)(\d+)`)

	for i := 0; i < len(inputMap); i += 4 {
		machine := &MachineData{}
		for j := 0; j < 3; j++ {
			matches := REGEX_PATTERN_DATA.FindStringSubmatch(inputMap[i+j])

			x, _ := strconv.Atoi(matches[4])
			y, _ := strconv.Atoi(matches[6])
			switch matches[2] {
			case "A":
				machine.buttonA.x = x
				machine.buttonA.y = y
			case "B":
				machine.buttonB.x = x
				machine.buttonB.y = y
			case "":
				machine.prize.x = x
				machine.prize.y = y
			}
		}

		ml.Data = append(ml.Data, *machine)
	}
}

func (ml *MachineData) solveMachine() int {
	ax, ay := ml.buttonA.x, ml.buttonA.y
	bx, by := ml.buttonB.x, ml.buttonB.y
	px, py := ml.prize.x, ml.prize.y

	a := (px*by - py*bx) / (ax*by - ay*bx)
	b := (ax*py - ay*px) / (ax*by - ay*bx)

	if a*ax+b*bx == px && a*ay+b*by == py { return 3*a + b}

	return 0
}

func PartOne(ml MachineList) int {
	sumTokens  := 0
	for _, machineData := range ml.Data { sumTokens += machineData.solveMachine() }
	return sumTokens
}

func PartTwo(ml MachineList) int {
	conversionError := 10000000000000
	sumTokens := 0
	for _, machineData := range ml.Data {
		machineData.prize.x += conversionError
		machineData.prize.y += conversionError
		sumTokens += machineData.solveMachine()
	}
	return sumTokens
}
func main() {
	fmt.Printf("--- Day 13: Claw Contraption ---\n")
	var machineList MachineList
	inputMap := fileparser.ReadFileLines("input", false)
	machineList.parseMachinesData(inputMap)
	fmt.Printf("PART[1]: %v\n", PartOne(machineList))
	fmt.Printf("PART[2]: %v\n", PartTwo(machineList))
}
