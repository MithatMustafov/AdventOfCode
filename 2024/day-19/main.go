package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"strings"
)

type LinenLayout struct {
	TowelPatterns []string
	Designs       []string
}

func (layout *LinenLayout) parseData(input []string) {
	layout.TowelPatterns = make([]string, 0)
	towelPatterns := strings.Split(input[0], ", ")
	layout.TowelPatterns = append(layout.TowelPatterns, towelPatterns...)
	layout.Designs = make([]string, 0)
	for idx := 2; idx < len(input); idx++ {
		layout.Designs = append(layout.Designs, input[idx])
	}
}

func countCombinations(target string, arr []string) int {
	cache  := make(map[string]int)
	count := 0
	for _, substr := range arr {
		if strings.HasPrefix(target, substr) {
			count += countWays(target[len(substr):], arr, cache)
		}
	}
	return count
}

func countWays(target string, arr []string, cache  map[string]int) int {
	if target == "" { return 1}
	if val, exists := cache [target]; exists { return val }
	count := 0
	for _, substr := range arr {
		if strings.HasPrefix(target, substr) {
			count += countWays(target[len(substr):], arr, cache )
		}
	}
	cache [target] = count
	return count
}

func PartOne(layout LinenLayout) int {
	countPossibleDesigns := 0
	for _, design := range layout.Designs {
		countCombinations := countCombinations(design,layout.TowelPatterns)
		if (countCombinations > 0 ) {countPossibleDesigns++}
	}
	return countPossibleDesigns
}

func PartTwo(layout LinenLayout) int {
	countPossibleDesigns := 0
	for _, design := range layout.Designs {
		countCombinations := countCombinations(design,layout.TowelPatterns)
		if countCombinations > 0  { countPossibleDesigns +=countCombinations }
	}
	return countPossibleDesigns
}

func main() {
	fmt.Printf("--- Day 19: Linen Layout ---\n")
	var linenLayout LinenLayout
	input := fileparser.ReadFileLines("input", false)
	linenLayout.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(linenLayout))
	fmt.Printf("PART[2]: %v\n", PartTwo(linenLayout))
}