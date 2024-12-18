
package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type TheAlgorithm struct {
	FallingBytes []Coordinates
	memorySize int
	memoryRange int
}

const RUNE_MEMORY_SAFE = '.' 
const RUNE_MEMORY_CORRUPTED = '#' 

func (algo *TheAlgorithm) parseData(input []string) {
    algo.FallingBytes = make([]Coordinates, 0)
    for _, line := range input {
        coords := strings.Split(line, ",")
        x, _ := strconv.Atoi(coords[0])
        y, _ := strconv.Atoi(coords[1])
        algo.FallingBytes = append(algo.FallingBytes, Coordinates{X: x, Y: y})
    }
}

func (algo *TheAlgorithm) generateMemorySpace() ([][]rune){
	memorySize, memoryRange := algo.memorySize, algo.memoryRange
	memorySize++
    memorySpace := make([][]rune, memorySize)
    for i := range memorySpace {
        memorySpace[i] = make([]rune, memorySize)
        for j := range memorySpace[i] {
            memorySpace[i][j] = RUNE_MEMORY_SAFE
        }
    }

    for _, fallingByte := range algo.FallingBytes {
		if memoryRange == 0  { break }
		memoryRange--
        memorySpace[fallingByte.Y][fallingByte.X] = RUNE_MEMORY_CORRUPTED		
    }

	return memorySpace
}

func (algo *TheAlgorithm) drawMemorySpace(ms [][]rune) {
	for _, y := range ms { for _, x := range y { fmt.Printf("%c ", x) }; fmt.Println() }
}

func bfsShortestPath(grid [][]rune, start Coordinates, end Coordinates, wallRune rune) ([]Coordinates, int) {
	rows := len(grid)
	if rows == 0 { return nil, -1 }
	cols := len(grid[0])

	directions := []Coordinates{{0, 1},{0, -1},{1, 0},{-1, 0},}

	isValid := func(y, x int) bool { 
		return y >= 0 && y < rows && x >= 0 && x < cols && grid[y][x] != wallRune
	}

	queue := []struct {
		Coord Coordinates
		Path  []Coordinates
	}{{start, []Coordinates{start}}}

	visited := make([][]bool, rows)
	for i := range visited { visited[i] = make([]bool, cols)}
	visited[start.Y][start.X] = true

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if current.Coord == end { return current.Path, 0 }

		for _, dir := range directions {
			newY, newX := current.Coord.Y+dir.Y, current.Coord.X+dir.X

			if isValid(newY, newX) && !visited[newY][newX] {
				visited[newY][newX] = true
				newPath := append([]Coordinates{}, current.Path...)
				newPath = append(newPath, Coordinates{newX, newY})
				queue = append(queue, struct { Coord Coordinates; Path  []Coordinates }{Coordinates{newX, newY}, newPath})
			}
		}
	}

	return nil, -1
}


func PartOne(algo TheAlgorithm) int {
	memorySpace := algo.generateMemorySpace()
	//algo.drawMemorySpace(memorySpace)
	shortestPath, _ := bfsShortestPath(
		memorySpace,
		Coordinates{0,0},
		Coordinates{algo.memorySize,algo.memorySize},
		RUNE_MEMORY_CORRUPTED)
	minSteps := len(shortestPath) - 1
	return minSteps
}


func PartTwo(algo TheAlgorithm) string {
	err := 0
	for ; err != -1 ; algo.memoryRange++ {
		memorySpace := algo.generateMemorySpace()
		_, err := bfsShortestPath(
		memorySpace,
		Coordinates{0,0},
		Coordinates{algo.memorySize,algo.memorySize},
		RUNE_MEMORY_CORRUPTED)
	    if (err == -1) {break;}
	}
	algo.memoryRange--
	theLastByte := algo.FallingBytes[algo.memoryRange]
	answer := strconv.Itoa(theLastByte.X) + "," + strconv.Itoa(theLastByte.Y)
	return answer
}

func main() {
	fmt.Printf("--- Day 18: RAM Run ---\n")
	var theAlgorithm TheAlgorithm
	input := fileparser.ReadFileLines("input", false)
	theAlgorithm.parseData(input)
	theAlgorithm.memorySize = 70
	theAlgorithm.memoryRange = 1024
	fmt.Printf("PART[1]: %v\n", PartOne(theAlgorithm))
	fmt.Printf("PART[2]: %v\n", PartTwo(theAlgorithm))
}
