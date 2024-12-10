package main

import (
	fileparser "aoc/utils/fileparser"
	"fmt"
)

type Hoofit struct {
	TopographicMap [][]rune
}

func (hooit *Hoofit) parseData(input []string) {
	hooit.TopographicMap = [][]rune{}
	for y, line := range input {
		hooit.TopographicMap = append(hooit.TopographicMap, []rune{})
		for _, cell := range line {
			hooit.TopographicMap[y] = append(hooit.TopographicMap[y], cell)
		}
	}
}

func dfs_findAllNines(grid [][]rune, r, c rune, target rune, visited map[[2]int]bool) int {
    rows := len(grid)
    cols := len(grid[0])

    count := 0
    if grid[r][c] == '9' { count++ }

    visited[[2]int{int(r), int(c)}] = true

    directions := [][2]int{ {-1, 0}, {1, 0}, {0, -1}, {0, 1}, }

    for _, dir := range directions {
        nr, nc := int(r)+dir[0], int(c)+dir[1]

        if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[[2]int{nr, nc}] {
            if grid[nr][nc] == target {
                count += dfs_findAllNines(grid, rune(nr), rune(nc), target+1, visited)
            }
        }
    }

    return count
}

func dfs_findAllPathsToNines(grid [][]rune, r, c rune, target rune, visited map[[2]int]bool) int {
    rows := len(grid)
    cols := len(grid[0])

    if grid[r][c] == '9' { return 1 }
    visited[[2]int{int(r), int(c)}] = true
    defer func() { visited[[2]int{int(r), int(c)}] = false }() 

    directions := [][2]int{ {-1, 0}, {1, 0}, {0, -1}, {0, 1},}

    pathCount := 0

    for _, dir := range directions {
        nr, nc := int(r)+dir[0], int(c)+dir[1]
        if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[[2]int{nr, nc}] {
            if grid[nr][nc] == target {
                pathCount += dfs_findAllPathsToNines(grid, rune(nr), rune(nc), target+1, visited)
            }
        }
    }

    return pathCount
}

func PartOne(hoofit Hoofit) int {
	grid := hoofit.TopographicMap
    rows := len(grid)
    cols := len(grid[0])

    var startRow, startCol int
    count := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '0' {
                startRow, startCol = r, c
                visited := make(map[[2]int]bool)
				count += dfs_findAllNines(grid, rune(startRow), rune(startCol), '1', visited)
            }
        }
    }

    return count
}

func PartTwo(hoofit Hoofit) int {
	grid := hoofit.TopographicMap
    rows := len(grid)
    cols := len(grid[0])

    var startRow, startCol int
    count := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '0' {
                startRow, startCol = r, c
                visited := make(map[[2]int]bool)
				count += dfs_findAllPathsToNines(grid, rune(startRow), rune(startCol), '1', visited)
            }
        }
    }

    return count
}

func main() {
	fmt.Printf("--- Day 10: Hoof It ---\n")
	var hoofit  Hoofit
	input := fileparser.ReadFileLines("input", false)
	hoofit.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(hoofit))
	fmt.Printf("PART[2]: %v\n", PartTwo(hoofit))
}
