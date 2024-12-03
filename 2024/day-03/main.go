package main

import (
	fileparser "aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var REGEX_MUL = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func partOne(lines []string) {
	var sum = 0
	for _, corruptedMemory := range lines {
		matches := REGEX_MUL.FindAllStringSubmatch(corruptedMemory, -1)
		for _, match := range matches {
			number1, _ := strconv.Atoi(match[1])
			number2, _ := strconv.Atoi(match[2])
			sum += number1 * number2
		}
	}

	fmt.Printf("%v \n", sum)
}

func partTwo(lines []string) {
	var totalSum = 0
	theLine := strings.Join(lines, "")
	corruptedMemory := strings.Split(theLine, "do()")
	var transformedMemory []string

	for _, memory  := range corruptedMemory {
		memoryDO := "DO" + memory
		memoryDONOT := strings.Split(memoryDO, "don't()")
		for _, newMemory := range memoryDONOT {
			transformedMemory = append(transformedMemory, "NOT"+newMemory)
		}
	}

	for _, memory := range transformedMemory {
		if len(memory) > 3 && memory[3] == 'D' {
			matches := REGEX_MUL.FindAllStringSubmatch(memory, -1)
			for _, match := range matches {
				number1, _ := strconv.Atoi(match[1])
				number2, _ := strconv.Atoi(match[2])
				totalSum += number1 * number2
			}
		}
	}

	fmt.Printf("%v \n", totalSum)
}

func main() {
	lines := fileparser.ReadFileLines("inputFull",false)
	fmt.Printf("--- Day 3: Mull It Over ---\n")
	fmt.Printf("PART[1] ")
	partOne(lines)
	fmt.Printf("PART[2] ")
	partTwo(lines)
}
