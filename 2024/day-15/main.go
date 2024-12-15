package main

import (
	"aoc/utils/fileparser"
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

func (a *Coordinates) Add(b Coordinates) { a.X += b.X; a.Y += b.Y }

var RULE_ROBOT_MOVEMENTS = map[rune]Coordinates{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

type WarehouseWoes struct {
	WarehouseMap  [][]rune
	RobotPosition Coordinates
	RobotMoves    []rune
}

const SYMBOL_ROBOT = '@'
const SYMBOL_WALL = '#'
const SYMBOL_EMPTY = '.'
const SYMBOL_BOX = 'O'
const SYMBOL_BOX_LEFT = '['
const SYMBOL_BOX_RIGHT = ']'

func (ww *WarehouseWoes) parseData(inputMap []string) {
	ww.WarehouseMap, ww.RobotMoves = [][]rune{}, []rune{}
	WarehouseMap, robotMoves, robotPosition := ww.WarehouseMap, ww.RobotMoves, ww.RobotPosition

	lineIndex := 0
	for ; inputMap[lineIndex] != ""; lineIndex++ {
		WarehouseMap = append(WarehouseMap, []rune{})
		for x, cell := range inputMap[lineIndex] {
			WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], cell)
			if cell == SYMBOL_ROBOT {
				robotPosition.X = x
				robotPosition.Y = lineIndex
			}
		}
	}

	for ; lineIndex < len(inputMap); lineIndex++ {
		for _, cell := range inputMap[lineIndex] {
			robotMoves = append(robotMoves, cell)
		}
	}

	ww.WarehouseMap, ww.RobotMoves, ww.RobotPosition = WarehouseMap, robotMoves, robotPosition
}

func (ww WarehouseWoes) printMap() {
	//time.Sleep(time.Duration(0.5 * float64(time.Second)))
	for _, valueY := range ww.WarehouseMap {
		for _, valueX := range valueY {
			fmt.Printf("%v", string(valueX))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (ww *WarehouseWoes) moveRobot(nextStep Coordinates) {
	WarehouseMap := ww.WarehouseMap
	posRobot := Coordinates{X: ww.RobotPosition.X, Y: ww.RobotPosition.Y}
	posNext := Coordinates{X: ww.RobotPosition.X + nextStep.X, Y: ww.RobotPosition.Y + nextStep.Y}
	posCurrent := Coordinates{X: ww.RobotPosition.X + nextStep.X, Y: ww.RobotPosition.Y + nextStep.Y}
	currentCell := WarehouseMap[posNext.Y][posNext.X]

	if currentCell == SYMBOL_WALL {
		return
	}

	for ; ; posCurrent.Add(nextStep) {
		currentCell = WarehouseMap[posCurrent.Y][posCurrent.X]
		if currentCell == SYMBOL_WALL {
			return
		}
		if currentCell == SYMBOL_EMPTY {
			WarehouseMap[posCurrent.Y][posCurrent.X] = SYMBOL_BOX
			break
		}
	}

	WarehouseMap[posRobot.Y][posRobot.X] = SYMBOL_EMPTY
	WarehouseMap[posNext.Y][posNext.X] = SYMBOL_ROBOT

	ww.RobotPosition, ww.WarehouseMap = posNext, WarehouseMap
}

func (ww WarehouseWoes) sumBoxesGPSCoords() int {
	WarehouseMap, sumBoxes := ww.WarehouseMap, 0
	for idxY, valueY := range WarehouseMap {
		for idxX, valueX := range valueY {
			if valueX == SYMBOL_BOX {
				sumBoxes += 100*idxY + idxX
			}
		}
	}
	return sumBoxes
}

func PartOne(ww WarehouseWoes) int {
	robotMoves := ww.RobotMoves
	for _, move := range robotMoves {
		ww.moveRobot(RULE_ROBOT_MOVEMENTS[move])
	}
	return ww.sumBoxesGPSCoords()
}

func (ww *WarehouseWoes) scaledUp(inputMap []string) {
	ww.WarehouseMap, ww.RobotMoves = [][]rune{}, []rune{}
	WarehouseMap, robotMoves, robotPosition := ww.WarehouseMap, ww.RobotMoves, ww.RobotPosition

	lineIndex := 0
	for ; inputMap[lineIndex] != ""; lineIndex++ {
		WarehouseMap = append(WarehouseMap, []rune{})
		for _, cell := range inputMap[lineIndex] {
			if cell == SYMBOL_BOX {
				WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], SYMBOL_BOX_LEFT)
				WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], SYMBOL_BOX_RIGHT)
				continue
			}

			if cell == SYMBOL_ROBOT {
				WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], cell)
				WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], SYMBOL_EMPTY)
				continue
			}

			WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], cell)
			WarehouseMap[lineIndex] = append(WarehouseMap[lineIndex], cell)
		}
	}

	for ; lineIndex < len(inputMap); lineIndex++ {
		for _, cell := range inputMap[lineIndex] {
			robotMoves = append(robotMoves, cell)
		}
	}

	for idxY, valueY := range WarehouseMap {
		for idxX, valueX := range valueY {
			if valueX == SYMBOL_ROBOT {
				robotPosition.X = idxX
				robotPosition.Y = idxY
			}
		}
	}

	ww.WarehouseMap, ww.RobotMoves, ww.RobotPosition = WarehouseMap, robotMoves, robotPosition
}

type boxCapsule struct {
	boxLeft   Coordinates
	leftRune  rune
	boxRight  Coordinates
	rightRune rune
}

func getBox(givenBoxPartSymbol rune, y int, givenBoxPartCoord Coordinates) boxCapsule {
	var boxData boxCapsule
	boxData.boxLeft.Y = y
	boxData.boxLeft.X = givenBoxPartCoord.X
	if givenBoxPartSymbol == SYMBOL_BOX_RIGHT {
		boxData.boxLeft.X -= 1
	}
	boxData.boxRight.Y = y
	boxData.boxRight.X = boxData.boxLeft.X + 1
	return boxData
}

func (ww *WarehouseWoes) moveRobotScaledUp(nextStep Coordinates) {
	//fmt.Printf("%v\n", nextStep)
	originalMap := ww.WarehouseMap
	WarehouseMap := make([][]rune, len(originalMap))

	for i := range originalMap {
		WarehouseMap[i] = make([]rune, len(originalMap[i]))
		copy(WarehouseMap[i], originalMap[i])
	}

	var posRobot = Coordinates{X: ww.RobotPosition.X, Y: ww.RobotPosition.Y}
	var posNext = Coordinates{X: ww.RobotPosition.X + nextStep.X, Y: ww.RobotPosition.Y + nextStep.Y}
	var cellNext = WarehouseMap[posNext.Y][posNext.X]
	posCurrent := Coordinates{X: ww.RobotPosition.X + nextStep.X, Y: ww.RobotPosition.Y + nextStep.Y}

	if cellNext == SYMBOL_WALL {
		return
	}

	if nextStep.Y != 0 {
		if cellNext == SYMBOL_EMPTY {
			WarehouseMap[posRobot.Y][posRobot.X] = SYMBOL_EMPTY
			WarehouseMap[posNext.Y][posNext.X] = SYMBOL_ROBOT
		} else {
			eventAffectedBoxes := make([]boxCapsule, 0)
			//fmt.Printf("%v %v\n", string(cellNext), posRobot)

			var tempBoxCapsule = getBox(cellNext, posNext.Y, posRobot)
			tempBoxCapsule.leftRune = WarehouseMap[tempBoxCapsule.boxLeft.Y][tempBoxCapsule.boxLeft.X]
			tempBoxCapsule.rightRune = WarehouseMap[tempBoxCapsule.boxRight.Y][tempBoxCapsule.boxRight.X]

			eventAffectedBoxes = append(eventAffectedBoxes, tempBoxCapsule)

			processedBoxes := make(map[string]bool) // Use a unique key for map
			for idx := 0; idx < len(eventAffectedBoxes); idx++ {
				event := eventAffectedBoxes[idx]
				eventKey := fmt.Sprintf("%d:%d-%d:%d", event.boxLeft.X, event.boxLeft.Y, event.boxRight.X, event.boxRight.Y)
				if processedBoxes[eventKey] {
					continue
				}
				processedBoxes[eventKey] = true

				//fmt.Printf("processedBoxes: %v\n", processedBoxes)

				left := event.boxLeft
				cellNextLeft := WarehouseMap[left.Y+nextStep.Y][left.X]

				right := event.boxRight
				cellNextRight := WarehouseMap[right.Y+nextStep.Y][right.X]

				//fmt.Printf("Events: \n %v %v %v \n %v %v %v \n", string(event.leftRune), left, string(cellNextLeft), string(event.rightRune), right, string(cellNextRight))

				if event.leftRune == SYMBOL_WALL || event.rightRune == SYMBOL_WALL {
					break
				}
				if cellNextLeft == SYMBOL_WALL || cellNextRight == SYMBOL_WALL {
					//fmt.Printf("Wall encountered, stopping processing.\n")
					return
				}

				if cellNextLeft != SYMBOL_EMPTY {
					tempBoxCapsule = getBox(cellNextLeft, left.Y+nextStep.Y, left)
					tempBoxCapsule.leftRune = WarehouseMap[tempBoxCapsule.boxLeft.Y][tempBoxCapsule.boxLeft.X]
					tempBoxCapsule.rightRune = WarehouseMap[tempBoxCapsule.boxRight.Y][tempBoxCapsule.boxRight.X]
					eventAffectedBoxes = append(eventAffectedBoxes, tempBoxCapsule)
					//fmt.Printf("L")
				} else {
					WarehouseMap[left.Y+nextStep.Y][left.X] = event.leftRune
					
				}
				WarehouseMap[left.Y][left.X] = SYMBOL_EMPTY
				if cellNextRight != SYMBOL_EMPTY {
					tempBoxCapsule = getBox(cellNextRight, right.Y+nextStep.Y, right)
					tempBoxCapsule.leftRune = WarehouseMap[tempBoxCapsule.boxLeft.Y][tempBoxCapsule.boxLeft.X]
					tempBoxCapsule.rightRune = WarehouseMap[tempBoxCapsule.boxRight.Y][tempBoxCapsule.boxRight.X]
					eventAffectedBoxes = append(eventAffectedBoxes, tempBoxCapsule)
					//fmt.Printf("R")
				} else {
					WarehouseMap[right.Y+nextStep.Y][right.X] = event.rightRune
				}
				WarehouseMap[right.Y][right.X] = SYMBOL_EMPTY

				//fmt.Printf("\n")
				//ww.printMap()
			}

			//fmt.Printf("---FIX----\n")
			for idx := range eventAffectedBoxes {
				box := eventAffectedBoxes[idx]
				WarehouseMap[box.boxLeft.Y+nextStep.Y][box.boxLeft.X] = box.leftRune
				WarehouseMap[box.boxRight.Y+nextStep.Y][box.boxRight.X] = box.rightRune
				//ww.printMap()
			}
			//fmt.Printf("%v\n",posRobot)
			ww.WarehouseMap = WarehouseMap
			//ww.printMap()
		}
	} else {
		var prevCell rune
		prevCell = SYMBOL_EMPTY
		for ; ; posCurrent.Add(nextStep) {
			currentCell := WarehouseMap[posCurrent.Y][posCurrent.X]
			if currentCell == SYMBOL_WALL {
				return
			}
			if currentCell == SYMBOL_EMPTY {
				WarehouseMap[posCurrent.Y][posCurrent.X] = prevCell
				prevCell = currentCell
				break
			}
			WarehouseMap[posCurrent.Y][posCurrent.X], prevCell = prevCell, currentCell
		}
	}

	WarehouseMap[ww.RobotPosition.Y][ww.RobotPosition.X] = SYMBOL_EMPTY
	WarehouseMap[ww.RobotPosition.Y+nextStep.Y][ww.RobotPosition.X+nextStep.X] = SYMBOL_ROBOT

	ww.WarehouseMap = WarehouseMap
	ww.RobotPosition = Coordinates{X: ww.RobotPosition.X + nextStep.X, Y: ww.RobotPosition.Y + nextStep.Y}
}

func (ww WarehouseWoes) sumBoxesGPSCoordsv2() int {
	WarehouseMap, sumBoxes := ww.WarehouseMap, 0
	for idxY, valueY := range WarehouseMap {
		for idxX, valueX := range valueY {
			if valueX == SYMBOL_BOX_LEFT {
				sumBoxes += 100*idxY + idxX
			}
		}
	}
	return sumBoxes
}

func PartTwo(ww WarehouseWoes) int {
	robotMoves := ww.RobotMoves
	for _, move := range robotMoves {
		//fmt.Printf("MOVE: %v\n", string(move))
		ww.moveRobotScaledUp(RULE_ROBOT_MOVEMENTS[move])
		//ww.printMap()
	}

	//ww.printMap()
	return ww.sumBoxesGPSCoordsv2()
}

func main() {
	fmt.Printf("--- Day 15: Warehouse Woes ---\n")
	var warehouseWoes WarehouseWoes
	inputMap := fileparser.ReadFileLines("input", false)
	warehouseWoes.parseData(inputMap)
	fmt.Printf("PART[1]: %v\n", PartOne(warehouseWoes))
	warehouseWoes.scaledUp(inputMap)
	//warehouseWoes.printMap()
	fmt.Printf("PART[2]: %v\n", PartTwo(warehouseWoes))
}
