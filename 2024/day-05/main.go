package main

import (
	fileparser "aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Instructions struct {
	Rules   map[int][]int
	Updates [][]int
}

func transformInput(INPUT *[]string) Instructions {
	var instructions Instructions
	instructions.Rules = make(map[int][]int)
	instructions.Updates = make([][]int, 0)

	idx := 0
	for ; idx < len(*INPUT); idx++ {
		if (*INPUT)[idx] == "" {break}
		ruleParts := strings.Split((*INPUT)[idx], "|")
		x, _ := strconv.Atoi(ruleParts[0])
		y, _ := strconv.Atoi(ruleParts[1])
		instructions.Rules[x] = append(instructions.Rules[x], y)
	}
	idx++
	for ; idx < len(*INPUT); idx++ {
		updatePages := strings.Split((*INPUT)[idx], ",")
		newUpdate := make([]int, len(updatePages))
		for idx, page := range updatePages { newUpdate[idx], _ = strconv.Atoi(page) }
		instructions.Updates = append(instructions.Updates, newUpdate)
	}

	return instructions
}

func isValidRule(PAGE *int, RULE *[]int) bool{

	for _, pageRule := range *RULE {
		if *PAGE == pageRule {return true}
	}

	return false
}

func isValidUpdate(UPDATE *[]int, RULES *map[int][]int) (bool, int, int) {
    for idxPage, x := range *UPDATE {
        for idxNextPage := idxPage + 1; idxNextPage < len(*UPDATE); idxNextPage++ {
            pagePtr := &(*UPDATE)[idxNextPage]
            ruleSlice := (*RULES)[x]
            rulePtr := &ruleSlice
            if isValidRule(pagePtr, rulePtr) { continue }
            return false, idxPage, idxNextPage
        }
    }

    return true, -1, -1
}


func orderTheUpdate(UPDATE *[]int, RULES *map[int][]int) *[]int{
	orderedUpdate := UPDATE
	for{
		isOrdered, idxPageOne, idxPageTwo := isValidUpdate(orderedUpdate, RULES);
		if isOrdered { break; }
		(*orderedUpdate)[idxPageOne], (*orderedUpdate)[idxPageTwo] =  (*orderedUpdate)[idxPageTwo], (*orderedUpdate)[idxPageOne]
	}
	return orderedUpdate;
}

func partOne(INSTRUCTIONS *Instructions) int {
	sumMiddlePageNumbers := 0

	for _,update := range INSTRUCTIONS.Updates {
		isValid, _, _ := isValidUpdate(&(update), &(INSTRUCTIONS.Rules))
		if isValid { sumMiddlePageNumbers += update[len(update)/2] }
	}

	return sumMiddlePageNumbers
}

func partTwo(INSTRUCTIONS *Instructions) int {
	sumMiddlePageNumbers := 0

	for _,update := range INSTRUCTIONS.Updates {
		isValid, _, _ := isValidUpdate(&(update), &(INSTRUCTIONS.Rules));
		if !isValid { 
			orderedUpdate := orderTheUpdate(&(update), &(INSTRUCTIONS.Rules));
			sumMiddlePageNumbers += (*orderedUpdate)[len((*orderedUpdate))/2] 
		}
	}

	return sumMiddlePageNumbers
}

func main() {
	fmt.Printf("--- Day 5: Print Queue ---\n")
	input := fileparser.ReadFileLines("input", false)
	instructions := transformInput(&input)
	fmt.Printf("PART[1] %v\n", partOne(&instructions))
	fmt.Printf("PART[2] %v\n", partTwo(&instructions))
}
