package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"math"
	"strconv"
)

type Coordinates struct {
	X int
	Y int
}

var KEYPAD_CTV = map[Coordinates]rune{
	{X: 0, Y: 0}: '7', {X: 1, Y: 0}: '8', {X: 2, Y: 0}: '9',
	{X: 0, Y: 1}: '4', {X: 1, Y: 1}: '5', {X: 2, Y: 1}: '6',
	{X: 0, Y: 2}: '3', {X: 1, Y: 2}: '2', {X: 2, Y: 2}: '3',
	{X: 0, Y: 3}: ' ', {X: 1, Y: 3}: '0', {X: 2, Y: 3}: 'A',
}

var KEYPAD_VTC = map[rune]Coordinates{
	'7': {X: 0, Y: 0}, '8': {X: 1, Y: 0}, '9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1}, '5': {X: 1, Y: 1}, '6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2}, '2': {X: 1, Y: 2}, '3': {X: 2, Y: 2},
	' ': {X: 0, Y: 3}, '0': {X: 1, Y: 3}, 'A': {X: 2, Y: 3},
}

var KEYPAD_GRID = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'-', '0', 'A'},
}

var REMOTE_GRID = [][]rune{
	{'-', '^', 'A'},
	{'<', 'v', '>'},
}

var REMOTE_CTV = map[Coordinates]rune{
	{X: 0, Y: 0}: ' ', {X: 1, Y: 0}: '^', {X: 2, Y: 0}: 'A',
	{X: 0, Y: 1}: '<', {X: 1, Y: 1}: 'v', {X: 2, Y: 1}: '>',
}

var REMOTE_VTC = map[rune]Coordinates{
	' ': {X: 0, Y: 0}, '^': {X: 1, Y: 0}, 'A': {X: 2, Y: 0},
	'<': {X: 0, Y: 1}, 'v': {X: 1, Y: 1}, '>': {X: 2, Y: 1},
}

var RULE_DIRECTIONS = map[Coordinates]rune{
	{X: 0, Y: -1}: '^',
	{X: 1, Y: 0}:  '>',
	{X: -1, Y: 0}: '<',
	{X: 0, Y: 1}:  'v',
}

type KeypadConundrum struct {
	Codes []string
}

func bfsContains(path []Coordinates, coord Coordinates) bool {
	for _, step := range path { if step == coord { return true } }
	return false
}

func bfsAllShortestPaths(grid [][]rune, start Coordinates, end Coordinates, wallRune rune) ([][]Coordinates, int) {
	rows := len(grid)
	cols := len(grid[0])
	directions := []Coordinates{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	isValid := func(y, x int) bool {
		return y >= 0 && y < rows && x >= 0 && x < cols && grid[y][x] != wallRune
	}
	queue := []struct {
		Coord Coordinates
		Path  []Coordinates
	}{{start, []Coordinates{start}}}

	var shortestPaths [][]Coordinates
	minPathLength := -1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Coord == end {
			if minPathLength == -1 || len(current.Path) == minPathLength {
				shortestPaths = append(shortestPaths, current.Path)
				minPathLength = len(current.Path)
			}
			continue
		}

	
		if minPathLength != -1 && len(current.Path) >= minPathLength { continue }

		for _, dir := range directions {
			newY, newX := current.Coord.Y+dir.Y, current.Coord.X+dir.X

			if isValid(newY, newX) {
				newCoord := Coordinates{newX, newY}

				if bfsContains(current.Path, newCoord) {continue}

				newPath := append([]Coordinates{}, current.Path...)
				newPath = append(newPath, newCoord)
				queue = append(queue, struct { Coord Coordinates; Path  []Coordinates }{newCoord, newPath})
			}
		}
	}

	if len(shortestPaths) == 0 { return nil, -1 }
	return shortestPaths, len(shortestPaths[0]) - 1
}

func (kk *KeypadConundrum) parseData(input []string) {
	kk.Codes = make([]string, 0);
	kk.Codes = append(kk.Codes, input...)
}

func getDirections(coords []Coordinates) string {
	var result string
	for i := 0; i < len(coords)-1; i++ {
		current, next := coords[i], coords[i+1]
		deltaX := next.X - current.X
		deltaY := next.Y - current.Y
		delta := Coordinates{X: deltaX, Y: deltaY}
		result += string(RULE_DIRECTIONS[delta])
	}
	return result
}

func generatePaths(keypadVTC map[rune]Coordinates, keypadGrid [][]rune, currentPos Coordinates, code string, pathOne []string) []string {
    for _, button := range code {
        targetPos := keypadVTC[button]
        pathCoords, _ := bfsAllShortestPaths(keypadGrid, currentPos, targetPos, '-')

        newPaths := []string{}
        for _, newPath := range pathCoords {
            newPathStr := getDirections(newPath) + "A"
            for _, existingPath := range pathOne {
                combinedPath := existingPath + newPathStr
                newPaths = append(newPaths, combinedPath)
            }
        }
        pathOne = newPaths
        currentPos = targetPos
    }

    return pathOne
}

func generateRemotePaths(cPathOne string, currentPos Coordinates, remoteVTC map[rune]Coordinates, remoteGrid [][]rune) []string {
    pathTwo := []string{""}
    for _, r := range cPathOne {
        targetPos := remoteVTC[r]
        pathCoords, _ := bfsAllShortestPaths(remoteGrid, currentPos, targetPos, '-')

        newPaths := []string{}
        for _, newPath := range pathCoords {
            newPathStr := getDirections(newPath) + "A"
            for _, existingPath := range pathTwo {
                combinedPath := existingPath + newPathStr
                newPaths = append(newPaths, combinedPath)
            }
        }
        pathTwo = newPaths
        currentPos = targetPos
    }
    return pathTwo
}

func calculateFinalCost(paths []string) int {
    min := math.MaxInt
    newId := -1

    for idx, path := range paths {
        costCalc := 0
        currentPos := Coordinates{X: 2, Y: 0}
        for _, step := range path {
            targetPos := REMOTE_VTC[step]
            costCalc += 
			int(math.Abs(float64(targetPos.X-currentPos.X))) + 
			int(math.Abs(float64(targetPos.Y-currentPos.Y)))
            currentPos = targetPos
        }

        if costCalc < min { min, newId = costCalc, idx}
    }
    return len(paths[newId]) + min
}

func PartOne(kk KeypadConundrum) int {
    sum := 0
    for _, code := range kk.Codes {
        pathOne := []string{""}
        currentPos := Coordinates{X: 2, Y: 3}

        pathOne = generatePaths(KEYPAD_VTC, KEYPAD_GRID, currentPos, code, pathOne)

        newPathTwo := []string{}
        currentPos = Coordinates{X: 2, Y: 0}
        for _, cPathOne := range pathOne {
            pathTwo := generateRemotePaths(cPathOne, currentPos, REMOTE_VTC, REMOTE_GRID)
            newPathTwo = append(newPathTwo, pathTwo...)
        }

        finalCost := calculateFinalCost(newPathTwo)
        newStr := code[0:3]
        numstr, _ := strconv.ParseInt(newStr, 10, 64)
        sum += int(numstr) * finalCost
    }

    return sum
}

func PartTwo(kk KeypadConundrum) int {

	return 0
}

func main() {
	fmt.Printf("--- Day 21: Keypad Conundrum ---\n")
	var keypadConundrum KeypadConundrum
	input := fileparser.ReadFileLines("input", false)
	keypadConundrum.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(keypadConundrum))
	fmt.Printf("PART[2]: %v\n", PartTwo(keypadConundrum))
}
