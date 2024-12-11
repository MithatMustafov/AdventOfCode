package main

import (
	fileparser "aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type PlutonianPebbles struct {
	BLINKS int
	Stones []int
}

func (pp *PlutonianPebbles) loadStones(inputMap string) {
	arrStrStones := strings.Split(inputMap, " ")
	for _, strStone := range arrStrStones {
		newStone, _ := strconv.Atoi(strStone)
		pp.Stones = append(pp.Stones, newStone)
	}
}

func PartOne(pp PlutonianPebbles) int {
	for idxBlink := 0; idxBlink < pp.BLINKS; idxBlink++ {
		var newStones []int
		for _, stone := range pp.Stones {
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			stoneStr := strconv.Itoa(stone)
			n := len(stoneStr)
			if n%2 == 0 {
				left, _ := strconv.Atoi(stoneStr[:n/2])
				right, _ := strconv.Atoi(stoneStr[n/2:])
				newStones = append(newStones, left)
				newStones = append(newStones, right)
				continue
			}

			newStones = append(newStones, stone*2024)
		}
		pp.Stones = make([]int, len(newStones))
		pp.Stones = newStones
	}
	return len(pp.Stones)
}

func PartTwo(pp PlutonianPebbles) int {
	mapStones := make(map[int]int)
	for _, stone := range pp.Stones {
		mapStones[stone] = 1
	}

	for idxBlink := 0; idxBlink < pp.BLINKS; idxBlink++ {
		newStones := make(map[int]int)
		for key, value := range mapStones {
			if key == 0 {
				newStones[1] += value
				continue
			}

			stoneStr := strconv.Itoa(key)
			n := len(stoneStr)
			if n%2 == 0 {
				left, _ := strconv.Atoi(stoneStr[:n/2])
				right, _ := strconv.Atoi(stoneStr[n/2:])
				newStones[left] += value
				newStones[right] += value
				continue
			}

			newStones[key*2024] += value
		}
		mapStones = newStones
	}

	stoneCounter := 0
	for _, val := range mapStones { stoneCounter += val }
	return stoneCounter
}

func main() {
	fmt.Printf("--- Day 11: Plutonian Pebbles ---\n")
	var plutonianPebbles PlutonianPebbles
	inputMap := fileparser.ReadFileLines("input", false)
	plutonianPebbles.loadStones(inputMap[0])
	plutonianPebbles.BLINKS = 25
	fmt.Printf("PART[1]: %v\n", PartOne(plutonianPebbles))
	plutonianPebbles.BLINKS = 75
	fmt.Printf("PART[2]: %v\n", PartTwo(plutonianPebbles))
}
