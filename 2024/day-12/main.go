package main

import (
	fileparser "aoc/utils/fileparser"
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

type GardenGroups struct {
	inputGardenMap  [][]rune
	betterGardenMap [][]rune
	placedPlusSign  map[Coordinates]rune
}

func (gG *GardenGroups) printGardeMap() {

	for _, row := range gG.betterGardenMap {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")
	}
}

func (gG *GardenGroups) printGardeMap1() {

	for _, row := range gG.inputGardenMap {
		for _, r := range row {
			fmt.Printf("[%c]", r)
		}
		fmt.Printf("\n")
	}
}

func (gG *GardenGroups) parseGardenMap(inputMap []string) {
	gG.inputGardenMap = make([][]rune, len(inputMap))
	for y, line := range inputMap {
		gG.inputGardenMap[y] = make([]rune, len(line))
		for x, char := range line {
			gG.inputGardenMap[y][x] = rune(char)
		}
	}
}

func (gG *GardenGroups) initBetterGardenMap() {
	gG.betterGardenMap = make([][]rune, len(gG.inputGardenMap)*2+1)
	for i := range gG.betterGardenMap {
		gG.betterGardenMap[i] = make([]rune, len(gG.inputGardenMap[0])*2+1)
		for j := range gG.betterGardenMap[i] {
			gG.betterGardenMap[i][j] = ' '
		}
	}

	for idxY, line := range gG.inputGardenMap {
		for idxX, char := range line {
			gG.betterGardenMap[idxY*2+1][idxX*2+1] = char
		}
	}
}

func (gG *GardenGroups) isAreaPlantSquare(x int, y int, plantSign rune, typeRune rune) bool {

	if x < 1 || x >= len(gG.betterGardenMap[y])-1 || y < 1 || y >= len(gG.betterGardenMap)-1 {
		return false
	}

	if typeRune == '+' &&
		gG.betterGardenMap[y-1][x-1] == plantSign &&
		gG.betterGardenMap[y-1][x+1] == plantSign &&
		gG.betterGardenMap[y+1][x-1] == plantSign &&
		gG.betterGardenMap[y+1][x+1] == plantSign {
			return true
	}

	if typeRune == '|' &&
		gG.betterGardenMap[y][x-1] == plantSign &&
		gG.betterGardenMap[y][x+1] == plantSign {
			return true
	}

	if typeRune == '-' &&
	gG.betterGardenMap[y+1][x] == plantSign &&
	gG.betterGardenMap[y-1][x] == plantSign {
		return true
	}

	return false
}

func (gG *GardenGroups) setPlotVisuallIndications(x int, y int) {
    pantCell := gG.betterGardenMap
    
    directions := []struct {
        dx, dy int
        symbol   rune
    }{
        {-1, -1, '+'}, {+1, -1, '+'}, {-1, +1, '+'}, {+1, +1, '+'},
        {+1, 0, '|'}, {-1, 0, '|'},
        {0, +1, '-'}, {0, -1, '-'},
    }

    for _, dir := range directions {
        nx, ny := x+dir.dx, y+dir.dy
        if !gG.isAreaPlantSquare(nx, ny, pantCell[y][x], dir.symbol) {
            pantCell[ny][nx] = dir.symbol
            coords := Coordinates{X: nx, Y: ny}
            if _, exists := gG.placedPlusSign[coords]; !exists {
                gG.placedPlusSign[coords] = dir.symbol
            }
        }
    }
}


func (gG *GardenGroups) getPlotData(idxX int, idxY int, plantType rune) (int, int, [][]rune) {
	plantCount := 0
	gG.placedPlusSign = make(map[Coordinates]rune)
	pantCell := gG.inputGardenMap
	
	rows := len(pantCell)
	cols := len(pantCell[0])
	
	directions := []struct{ dx, dy int }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	queue := []struct{ x, y int }{{idxX, idxY}}
	visited := make(map[Coordinates]bool)
	
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]
		cx, cy := current.x, current.y
		
		if cx < 0 || cy < 0 || cx >= cols || cy >= rows || visited[Coordinates{cx, cy}] {
			continue
		}
		
		visited[Coordinates{cx, cy}] = true
		
		if pantCell[cy][cx] < 'A' || pantCell[cy][cx] > 'Z' {
			continue
		}
		
		if pantCell[cy][cx] == plantType {
			pantCell[cy][cx] = '#'
			gG.setPlotVisuallIndications(cx*2+1, cy*2+1)
			plantCount++
			
			for _, d := range directions {
				nx, ny := cx+d.dx, cy+d.dy
				queue = append(queue, struct{ x, y int }{nx, ny})
			}
		}
	}
	
	plusSigns := 0
    fenceMap := make([][]rune, len(gG.betterGardenMap))
    for i := range fenceMap {
        fenceMap[i] = make([]rune, len(gG.betterGardenMap[0]))
    }
	for key, c := range gG.placedPlusSign {
		if c == '+' { plusSigns++; continue}
        fenceMap[key.Y][key.X] = c
	}
    plusSigns = len(gG.placedPlusSign)-plusSigns
	for c := range gG.placedPlusSign {
		delete(gG.placedPlusSign, c)
	}
	
	return plantCount, plusSigns,fenceMap
}

func countHorizontalWalls(fenceMap [][]rune) int {
    wallCount := 0
    for i := 0; i < len(fenceMap); i++ {
        if i%2 == 0 {
            prevChar := '\u0000'
            for j := 0; j < len(fenceMap[i]); j++ {
                if j%2 == 1 {
                    currentChar := fenceMap[i][j]
                    if currentChar != prevChar {
                        prevChar = currentChar
                        if prevChar != '\u0000' {
                            wallCount++
                        }
                    }
                    if i > 0 && i+1 < len(fenceMap) && j+2 < len(fenceMap[0]) &&
                        fenceMap[i][j+2] == '-' && fenceMap[i+1][j+1] == '|' && fenceMap[i-1][j+1] == '|' {
                        wallCount++
                    }
                }
            }
        }
    }
    return wallCount
}


func countVerticalWalls(fenceMap [][]rune) int {
    wallCount := 0
    for i := 0; i < len(fenceMap[0]); i++ {
        if i%2 == 0 {
            prevChar := '\u0000'
            for j := 0; j < len(fenceMap); j++ {
                if j%2 == 1 {
                    currentChar := fenceMap[j][i]
                    if currentChar != prevChar {
                        prevChar = currentChar
                        if prevChar != '\u0000' {
                            wallCount++
                        }
                    }
                    if j+2 < len(fenceMap) && i+1 < len(fenceMap[0]) && i-1 >= 0 &&
                        fenceMap[j+2][i] == '|' && fenceMap[j+1][i+1] == '-' && fenceMap[j+1][i-1] == '-' {
                        wallCount++
                    }
                }
            }
        }
    }
    return wallCount
}

func PartOne(gG GardenGroups) int {
	sumPrice := 0
	for idxY, y := range gG.inputGardenMap {
		for idxX, x := range y {
			newPlant := rune(x)
			plantCount, plusSigns, _ := gG.getPlotData(idxX, idxY, newPlant)
			sumPrice += plantCount * plusSigns
		}
	}
	return sumPrice
}

func PartTwo(gG GardenGroups) int {
    totalPrice := 0

    for idxY, row := range gG.inputGardenMap {
        for idxX, cell := range row {
            plant := rune(cell)
            plantCount, _, fenceMap := gG.getPlotData(idxX, idxY, plant)
            
            if plantCount == 0 {
                continue
            }

            horizontalWallCount := countHorizontalWalls(fenceMap)
            verticalWallCount := countVerticalWalls(fenceMap)

            totalPrice += plantCount * (horizontalWallCount + verticalWallCount)
        }
    }
    return totalPrice
}

func main() {
	fmt.Printf("--- Day 12: Garden Groups ---\n")
	var gardenGroups GardenGroups
	inputMap := fileparser.ReadFileLines("input", false)
	gardenGroups.parseGardenMap(inputMap)
	gardenGroups.initBetterGardenMap()
	fmt.Printf("PART[1] %v \n", PartOne(gardenGroups))
    gardenGroups.parseGardenMap(inputMap)
	gardenGroups.initBetterGardenMap()
	fmt.Printf("PART[2] %v \n", PartTwo(gardenGroups))
}
