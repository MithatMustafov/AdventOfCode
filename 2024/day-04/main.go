package main

import (
	fileparser "aoc/utils"
	"fmt"
	"strings"
)

var RULE_XMAS =  [][2]int{ {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
var RULE_XedMAS =  [][2]int{ {-1, -1}, {1, -1}, {1, 1}, {-1, 1}}

func transformInput(input []string) []string{
	var emptyLine string
	emptyLine += strings.Repeat(".", len(input)+6)
	newInput := make([]string, len(input)+6);
	for i := 0; i < 3; i++ { newInput[i] = emptyLine}
	for idx,line := range input {newInput[idx+3]="..."+line+"..."}
	for i := len(input)+3; i < len(input)+6; i++ { newInput[i] = emptyLine}
	return newInput;
}

func isXMAS( X int, Y int, RULE [][2]int, input []string) int {
	counter := 0
	const WORD = "MAS"
    for _, rule := range RULE {
		x := X; y := Y
        for _,letter := range WORD {
			y += rule[1]; x += rule[0]
            if string(input[y][x]) != string(letter) {goto NEXTRULE}
        }
        counter++; NEXTRULE:
	}
	return counter;
}

func isXedMAS( X int, Y int, RULE [][2]int, input []string) int {
	generatedWord := []string{"", ""}
    for _, rule := range RULE{
		idx := 0
		if rule[0] != rule[1] {idx = 1}
		letter := string(input[Y + rule[1]][X + rule[0]])
		if letter == "X" {return 0}
		generatedWord[idx] += letter
	}
	validWord := func(word string) bool {return word == "SM" || word == "MS"}
	if validWord(generatedWord[0]) && validWord(generatedWord[1]) {return 1}
	return 0
}

func partOne(input []string) int {
	counter := 0
	const target = "X"
	for y, line := range input{
		x := 0
		for {
			newX := strings.Index(line[x:], target)
			if newX == -1 {break}
			x += newX + 1
			counter += isXMAS(x-1, y, RULE_XMAS, input);
		}
	}
	return counter
}

func partTwo(input []string) int {
	counter := 0
	const target = "A"
	for y, line := range input{
		x := 0
		for {
			newA := strings.Index(line[x:], target)
			if newA == -1 {break}
			x += newA + 1
			counter += isXedMAS(x-1, y, RULE_XedMAS, input);
		}
	}
	return counter
}

func main() {
	input := fileparser.ReadFileLines("input", false)
	fmt.Printf("--- Day 4: Ceres Search ---\n")
	fmt.Printf("PART[1] %v\n", partOne(transformInput(input)))
	fmt.Printf("PART[2] %v\n", partTwo(transformInput(input)))
}
