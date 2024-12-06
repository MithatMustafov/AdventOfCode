package main

import (
	fileparser "aoc/utils"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var RULE_DIRECTIONS = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var symbolDIRECTIONS = []rune{'|', '-', '|', '-'}

const symbolAllDirections = '+'
const symbolGuard = '^'
const symbolNewGuard = '@'
const symbolGuardStep = 'X'
const symbolEmptyCell = '.'
const symbolObstruction = '#'

type Coordinates struct{ X, Y int }
type Guard struct {
	LocationStart       Coordinates
	LocationCurrent     Coordinates
	LocationLastRoation Coordinates
	MovementDirection   int
	MovementOffset      Coordinates
	LeftTheArea         bool
	StepSymbol          rune
	StepsCount          int
}
type Map [][]rune
type TheWatcher struct {
	GUARD             Guard
	MAP               Map
	SHOW_THE_TIMELINE bool
	SPEED             float64
	PART              bool
	PARADOXES         int
}

func (TheWatcher *TheWatcher)clearTimeline() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (TheWatcher *TheWatcher) stopTimeline() {
	time.Sleep(time.Duration(TheWatcher.SPEED * float64(time.Second)))
	TheWatcher.clearTimeline()
}

func (TheWatcher *TheWatcher) mapPrint() {
	TheWatcher.stopTimeline()
	fmt.Printf("GAME IS RUNNING...\n")
	for _, line := range TheWatcher.MAP {
		fmt.Printf("%v\n", string(line))
	}
}

func (TheWatcher *TheWatcher) mapUpdate(symbolUpdate rune) {
	mapSymbol := &TheWatcher.MAP[TheWatcher.GUARD.LocationCurrent.Y][TheWatcher.GUARD.LocationCurrent.X]

	if TheWatcher.PART {
		*mapSymbol = symbolGuardStep
		return
	}

	if *mapSymbol == symbolEmptyCell || *mapSymbol == symbolNewGuard {
		*mapSymbol = symbolUpdate
		return
	}

	if *mapSymbol == symbolDIRECTIONS[0] || *mapSymbol == symbolDIRECTIONS[1] {
		*mapSymbol = symbolAllDirections
	}
}

func (TheWatcher *TheWatcher) initTheWatcher(INPUT *[]string) {
	var newGuard Guard
	newGuard.LeftTheArea = false
	newGuard.MovementOffset.X = RULE_DIRECTIONS[newGuard.MovementDirection][0]
	newGuard.MovementOffset.Y = RULE_DIRECTIONS[newGuard.MovementDirection][1]
	for y, line := range *INPUT {
		x := strings.IndexRune(line, rune(symbolGuard))
		if x != -1 {
			newGuard.LocationStart.X = x
			newGuard.LocationStart.Y = y
			break
		}
	}

	newGuard.LocationCurrent.X = newGuard.LocationStart.X
	newGuard.LocationCurrent.Y = newGuard.LocationStart.Y
	var newMap = make([][]rune, len(*INPUT))
	for y, line := range *INPUT {
		newMap[y] = []rune(line)
	}

	TheWatcher.GUARD = newGuard
	TheWatcher.MAP = newMap
}

func initWatcherVariant(TheWatcherMain *TheWatcher) *TheWatcher {
	TheWatcherVariant := &TheWatcher{
		GUARD: TheWatcherMain.GUARD,
		MAP:   make([][]rune, len(TheWatcherMain.MAP)),
		SPEED: TheWatcherMain.SPEED,
	}
	for j, row := range TheWatcherMain.MAP {
		TheWatcherVariant.MAP[j] = make([]rune, len(row))
		copy(TheWatcherVariant.MAP[j], row)
	}
	TheWatcherVariant.SHOW_THE_TIMELINE = TheWatcherMain.SHOW_THE_TIMELINE
	TheWatcherVariant.GUARD.LocationLastRoation = TheWatcherMain.GUARD.LocationCurrent
	return TheWatcherVariant
}

func (TheWatcher *TheWatcher) isValidNextMove() bool {
	guardOld := TheWatcher.GUARD
	var guardNew Coordinates
	guardNew.X = guardOld.LocationCurrent.X + guardOld.MovementOffset.X
	guardNew.Y = guardOld.LocationCurrent.Y + guardOld.MovementOffset.Y
	mapX := len(TheWatcher.MAP[0])
	mapY := len(TheWatcher.MAP)

	if guardNew.X < 0 || guardNew.X == mapX || guardNew.Y < 0 || guardNew.Y == mapY {
		TheWatcher.GUARD.LeftTheArea = true
		return false
	}

	if TheWatcher.MAP[guardNew.Y][guardNew.X] == symbolObstruction {
		return false
	}

	TheWatcher.GUARD.StepSymbol = symbolDIRECTIONS[TheWatcher.GUARD.MovementDirection]
	return true
}

func (TheWatcher *TheWatcher) moveGuardForward() {
	guard := &TheWatcher.GUARD
	TheWatcher.mapUpdate(TheWatcher.GUARD.StepSymbol)
	guard.LocationCurrent.X += guard.MovementOffset.X
	guard.LocationCurrent.Y += guard.MovementOffset.Y
	if TheWatcher.MAP[guard.LocationCurrent.Y][guard.LocationCurrent.X] == symbolEmptyCell {
		guard.StepsCount++
	}
	TheWatcher.mapUpdate(symbolNewGuard)
}

func (TheWatcher *TheWatcher) moveGuardRight() {
	guard := &TheWatcher.GUARD
	nextMovementDirection := (guard.MovementDirection + 1) % 4
	guard.MovementDirection = nextMovementDirection
	guard.MovementOffset.X = RULE_DIRECTIONS[nextMovementDirection][0]
	guard.MovementOffset.Y = RULE_DIRECTIONS[nextMovementDirection][1]
	TheWatcher.mapUpdate(symbolAllDirections)
}

func (TheWatcherMain *TheWatcher) updateLastRotation() {
	TheWatcherMain.GUARD.LocationLastRoation.X = TheWatcherMain.GUARD.LocationCurrent.X
	TheWatcherMain.GUARD.LocationLastRoation.Y = TheWatcherMain.GUARD.LocationCurrent.Y
}

func partOne(TheWatcher *TheWatcher) int {
	TheWatcher.stopTimeline()
	for {
		if TheWatcher.SHOW_THE_TIMELINE {
			TheWatcher.mapPrint()
		}
		if !TheWatcher.isValidNextMove() {
			if TheWatcher.GUARD.LeftTheArea {
				break
			}
			TheWatcher.moveGuardRight()
			continue
		}
		TheWatcher.moveGuardForward()
	}

	return TheWatcher.GUARD.StepsCount + 1
}

func partTwo(TheWatcher *TheWatcher) int {
	TheWatcher.stopTimeline()
	counterVariant :=0
	for {
		counterVariant++
		TheWatcherVariant := initWatcherVariant(TheWatcher)
		if TheWatcherVariant.isValidNextMove() {

			newObstructionX := TheWatcher.GUARD.LocationCurrent.X + TheWatcher.GUARD.MovementOffset.X
			newObstructionY := TheWatcher.GUARD.LocationCurrent.Y + TheWatcher.GUARD.MovementOffset.Y
			if TheWatcherVariant.MAP[newObstructionY][newObstructionX] == symbolEmptyCell {
				TheWatcherVariant.MAP[newObstructionY][newObstructionX] = symbolObstruction
			}

			newVariantMap := make(map[Coordinates]int)
			for {
				newVariantMap[TheWatcherVariant.GUARD.LocationCurrent]++
				if newVariantMap[TheWatcherVariant.GUARD.LocationCurrent] > 3 {
					TheWatcher.PARADOXES++
					break
				}

				if TheWatcherVariant.SHOW_THE_TIMELINE { TheWatcherVariant.mapPrint() }

				if !TheWatcherVariant.isValidNextMove() {
					if TheWatcherVariant.GUARD.LeftTheArea { break }
					TheWatcherVariant.moveGuardRight()
					continue
				}

				TheWatcherVariant.moveGuardForward()

				if TheWatcher.SHOW_THE_TIMELINE {
					fmt.Printf("VARIANT: [%v] STEP: [%v]",counterVariant,TheWatcherVariant.GUARD.StepsCount)
				}
			}
		}

		TheWatcher.MAP[TheWatcher.GUARD.LocationStart.Y][TheWatcher.GUARD.LocationStart.X] = symbolNewGuard
		if TheWatcher.SHOW_THE_TIMELINE {
			TheWatcher.mapPrint()
		}
		if !TheWatcher.isValidNextMove() {
			if TheWatcher.GUARD.LeftTheArea {
				break
			}
			TheWatcher.moveGuardRight()
			TheWatcher.updateLastRotation()
			continue
		}
		TheWatcher.moveGuardForward()
	}

	return TheWatcher.PARADOXES
}

func main() {
	input := fileparser.ReadFileLines("input", false)
	var TheWatcher TheWatcher
	TheWatcher.initTheWatcher(&input)
	TheWatcher.PART = true
	TheWatcher.SHOW_THE_TIMELINE = false
	TheWatcher.SPEED = 0.01
	fmt.Printf("--- Day 6: Guard Gallivant ---\n")
	fmt.Printf("PART[1] %v\n", partOne(&TheWatcher))
	var nextPart string
	fmt.Print("Type to start Part 2: ")
	fmt.Scanln(&nextPart)
	TheWatcher.initTheWatcher(&input)
	TheWatcher.PART = false
	TheWatcher.SHOW_THE_TIMELINE = false
	TheWatcher.SPEED = 0.01
	fmt.Printf("PART[2] %v\n", partTwo(&TheWatcher))
}
