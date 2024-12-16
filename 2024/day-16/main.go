package main

import (
	"aoc/utils/fileparser"
	"container/heap"
	"fmt"
)

type Coordinates struct {
	x, y int
	orientation string 
}

var directions = map[string][]Coordinates{
    "N": {{x: -1, y: 0, orientation: "N"}},
    "S": {{x: 1, y: 0, orientation: "S"}},
    "E": {{x: 0, y: 1, orientation: "E"}},
    "W": {{x: 0, y: -1, orientation: "W"}},
}

var rotations = map[string][]string{
	"N": {"W", "E"},
	"S": {"E", "W"},
	"E": {"N", "S"},
	"W": {"S", "N"},
}

type PriorityQueueItem struct {
	point    Coordinates
	priority int
	path     []Coordinates
}

type PriorityQueue []PriorityQueueItem
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(PriorityQueueItem)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq); item := (*pq)[n-1]; *pq = (*pq)[:n-1] ;return item
}
func isValid(grid [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] != '#'
}
func bfs(grid [][]rune, start, end Coordinates) (int, []Coordinates) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	initialState := PriorityQueueItem{point: start, priority: 0, path: []Coordinates{start}}
	heap.Push(pq, initialState)

	visited := make(map[Coordinates]bool)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(PriorityQueueItem)
		curPoint, curPriority, curPath := item.point, item.priority, item.path

		if curPoint.x == end.x && curPoint.y == end.y { return curPriority, curPath }

		if visited[curPoint] { continue }

		visited[curPoint] = true

		forwardMove := directions[curPoint.orientation]
		nx, ny := curPoint.x+forwardMove[0].x, curPoint.y+forwardMove[0].y
		if isValid(grid, nx, ny) {
			newPoint := Coordinates{nx, ny, curPoint.orientation}
			if !visited[newPoint] {
				heap.Push(pq, PriorityQueueItem{
					point:    newPoint,
					priority: curPriority + 1,
					path:     append([]Coordinates{}, append(curPath, newPoint)...),
				})
			}
		}

		for _, newOrientation := range rotations[curPoint.orientation] {
			newPoint := Coordinates{curPoint.x, curPoint.y, newOrientation}
			if !visited[newPoint] {
				heap.Push(pq, PriorityQueueItem{
					point:    newPoint,
					priority: curPriority + ROTATION_COST,
					path: append([]Coordinates{}, append(curPath, newPoint)...),
				})
			}
		}
	}

	return -1, nil
}

type ReindeerMaze struct {
	MazeMap   [][]rune
}
const ROTATION_COST = 1000

func (rm *ReindeerMaze) parseData(inputMap []string) {
	rm.MazeMap = [][]rune{}
	for y, line := range inputMap {
		rm.MazeMap = append(rm.MazeMap, []rune{})
		for _, cell := range line {
			rm.MazeMap[y] = append(rm.MazeMap[y], cell)
		}
	}
}

func printGrid(grid [][]rune, path []Coordinates) {
	copyGrid := make([][]rune, len(grid))
	for i := range grid { copyGrid[i] = make([]rune, len(grid[i])); copy(copyGrid[i], grid[i]) }
	for _, p := range path { copyGrid[p.x][p.y] = 'O' }
	for _, row := range copyGrid { fmt.Println(string(row))}
}

func simulateMapWithPath(grid [][]rune, path []Coordinates, cost int) (int, [][]rune) {
	versionMapCounter := 0

	savedGrid := make([][]rune, len(grid))
	for i := range grid {
		savedGrid[i] = make([]rune, len(grid[i]))
		copy(savedGrid[i], grid[i])
	}

	for i := 0; i < len(path); i++ {
		newGrid := make([][]rune, len(grid))
		for i := range grid {
			newGrid[i] = make([]rune, len(grid[i]))
			copy(newGrid[i], grid[i])
		}

		newGrid[path[i].x][path[i].y] = '#'
		newCost, newPath := bfs(newGrid, path[0], path[len(path)-1])
		if newCost == cost {
			versionMapCounter++
			// fmt.Printf("[V %v]\n",versionMapCounter)
			// printGrid(newGrid, newPath)
			for _, coords := range newPath {
				savedGrid[coords.x][coords.y] = 'O'
			}
		}
	}

	return versionMapCounter, savedGrid
}

func PartOne(reindeerMaze ReindeerMaze) int {
	grid := reindeerMaze.MazeMap
	start := Coordinates{x: len(reindeerMaze.MazeMap)-2, y: 1, orientation: "E"}
	end := Coordinates{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'E' {
				end = Coordinates{x: i, y: j}
			}
		}
	}
	cost, _ := bfs(grid, start, end)
	return cost
}

func PartTwo(reindeerMaze ReindeerMaze) int {
	grid := reindeerMaze.MazeMap
	start := Coordinates{x: len(reindeerMaze.MazeMap)-2, y: 1, orientation: "E"}
	end := Coordinates{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'E' {
				end = Coordinates{x: i, y: j}
			}
		}
	}
	cost, path := bfs(grid, start, end)

	_, newGrid := simulateMapWithPath(grid, path, cost)
	counter  := 0
	for _, row := range newGrid {
		for _, cell := range row {
			if cell == 'O' { counter++}
		}
	}

	return counter
}

func main() {
	var  reindeerMaze ReindeerMaze
	inputMap := fileparser.ReadFileLines("input", false)
	reindeerMaze.parseData(inputMap)
	fmt.Printf("PART[1]: %v\n", PartOne(reindeerMaze))
	reindeerMaze.parseData(inputMap)
	fmt.Printf("PART[2]: %v\n", PartTwo(reindeerMaze))
}
