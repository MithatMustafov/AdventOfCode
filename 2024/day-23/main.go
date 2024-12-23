package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"sort"
	"strings"
)

type NetworkMap struct {
	Connections map[string][]string
}

func (nm *NetworkMap) parseData(input []string) {
	nm.Connections = make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, "-")
		source, destination := parts[0], parts[1]
		nm.Connections[source] = append(nm.Connections[source], destination)
		nm.Connections[destination] = append(nm.Connections[destination], source)
	}
}

func Factorial(n int) int {
	if n == 0 || n == 1 { return 1 }
	return n * Factorial(n-1)
}

func calcCombinations(n int, r int) int { return Factorial(n) / (Factorial(r) * Factorial(n-r)) }

func isConnected(nm NetworkMap, computerA, computerB string) bool {
	for _, connected := range nm.Connections[computerA] {
		if connected == computerB {
			return true
		}
	}
	return false
}

func PartOne(nm NetworkMap) int {
	seenCombinations := make(map[string]bool)

	for computerA, value := range nm.Connections {
		for i := 0; i < len(value)-1; i++ {
			for j := i + 1; j < len(value); j++ {
				computerB := value[i]
				computerC := value[j]

				if !isConnected(nm, computerB, computerC) { continue }

				if !(
				computerA[0] == 't' ||
				computerB[0] == 't'||
				computerC[0] == 't') {
					continue
				}

				elements := []string{computerA, computerB, computerC}
				sort.Strings(elements)
				combination := elements[0] + "," + elements[1] + "," + elements[2]

				if seenCombinations[combination] { continue }

				seenCombinations[combination] = true
			}
		}
	}

	return len(seenCombinations)
}

func PartTwo(nm NetworkMap) string {
	finalPassword := make(map[string]int)
	for computerA, value := range nm.Connections {
		combinations := 0
		items := make(map[string]int)
		for i := 0; i < len(value)-1; i++ {
			for j := i + 1; j < len(value); j++ {
				computerB := value[i]
				computerC := value[j]
				if !isConnected(nm, computerB, computerC) { continue }
				items[computerB]++
				items[computerC]++
				combinations++
			}
		}

		if combinations == calcCombinations(len(items), 2) && len(items) > len(finalPassword)-1 {
			finalPassword = items
			finalPassword[computerA]++
		}
	}

	computers := make([]string, 0, len(finalPassword))
	for key := range finalPassword { computers = append(computers, key) }
	sort.Strings(computers)
	var password string
	for _, computer := range computers { password += computer + "," }
	password = password[:len(password)-1]
	
	return password
}

func main() {
	fmt.Printf("--- Day 23: LAN Party ---\n")
	var networkMap NetworkMap
	input := fileparser.ReadFileLines("input", false)
	networkMap.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(networkMap))
	fmt.Printf("PART[2]: %v\n", PartTwo(networkMap))
}
