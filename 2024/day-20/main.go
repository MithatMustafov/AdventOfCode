package main

import (
	"aoc/utils/fileparser"
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

type RaceCondition struct {
	RacetrackMap [][]rune
	RaceStart    Coordinates
	RaceEnd      Coordinates
	ShortestPath []Coordinates
}

const RUNE_START = 'S'
const RUNE_END = 'E'
const RUNE_WALL = '#'

func (rc *RaceCondition) parseData(input []string) {
	rc.RacetrackMap = [][]rune{}
	for y, line := range input {
		rc.RacetrackMap = append(rc.RacetrackMap, []rune{})
		for x, cell := range line {
			rc.RacetrackMap[y] = append(rc.RacetrackMap[y], cell)
			if cell == RUNE_START {
				rc.RaceStart.X = x
				rc.RaceStart.Y = y
				continue
			}
			if cell == RUNE_END {
				rc.RaceEnd.X = x
				rc.RaceEnd.Y = y
			}
		}
	}
}

func bfs(grid [][]rune, start Coordinates, end Coordinates, wallRune rune) ([]Coordinates, int) {
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

func abs(x int) int { if x < 0 { return -x }; return x}

func(rc RaceCondition) calcDistance(p1, p2 Coordinates) int { return abs(p2.X-p1.X) + abs(p2.Y-p1.Y) }

func(rc RaceCondition) runCheat(disableCollision int, minSavedPicoseconds int) (int){
	counter := 0
	for idxCurrent := 0; idxCurrent < len(rc.ShortestPath)-1; idxCurrent++ {
		for idxNext := idxCurrent + 1; idxNext < len(rc.ShortestPath); idxNext++ {
			dist := rc.calcDistance(rc.ShortestPath[idxCurrent], rc.ShortestPath[idxNext])
			if dist > 0 && dist <= disableCollision {
				saved := idxNext - idxCurrent - dist
				if saved >= minSavedPicoseconds {
					counter++
				}
			}
		}
	}
	return counter
}

func main() {
	fmt.Printf("--- Day 20: Race Condition ---\n")
	var raceCondition RaceCondition
	puzzleInput := fileparser.ReadFileLines("input", false)
	raceCondition.parseData(puzzleInput)
	raceCondition.ShortestPath, _ = bfs(raceCondition.RacetrackMap, raceCondition.RaceStart, raceCondition.RaceEnd, RUNE_WALL)
	fmt.Printf("PART[1]: %v\n", raceCondition.runCheat(2,100))
	fmt.Printf("PART[2]: %v\n", raceCondition.runCheat(20,100))
}
