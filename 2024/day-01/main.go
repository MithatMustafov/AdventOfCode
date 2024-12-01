package main

import (
	"fmt"
	fileparser "aoc/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func partOne(arrNumL []int, arrNumR []int) {
	var sum int64
	for idx := range arrNumL {
		distance := arrNumL[idx] - arrNumR[idx]
		sum += int64(math.Abs(float64(distance)))
	}

	fmt.Printf("%v\n", sum)

}

func partTwo(arrNumL []int, arrNumR []int) {
	var sum int64
	for _, numL := range arrNumL {
		counter := 0
		for _, numR := range arrNumR {
			if numL == numR{
				counter++
			}
		}
		sum += int64(numL * counter);
	}

	fmt.Printf("%v\n", sum)
}

func main() {
	lines := fileparser.ReadFileLines("input", false)

	arrNumL := make([]int, len(lines))
	arrNumR := make([]int, len(lines))
	for idx, line := range lines {
		nums := strings.Split(line, " ")
		numA, _ := strconv.Atoi(nums[0])
		numB, _ := strconv.Atoi(nums[len(nums)-1])
		arrNumL[idx] = numA
		arrNumR[idx] = numB
	}

	slices.Sort(arrNumL)
	slices.Sort(arrNumR)


	fmt.Printf("--- Day 1: Historian Hysteria ---\n")
	fmt.Printf("PART[1] "); partOne(arrNumL, arrNumR)
	fmt.Printf("PART[2] "); partTwo(arrNumL, arrNumR)
}
