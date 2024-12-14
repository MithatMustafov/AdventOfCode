package main

import (
	fileparser "aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Coordinates struct {
	X int
	Y int
}

type RobotSettings struct {
	Position Coordinates
	Velocity Coordinates
}

type EasterBunnyHeadquarters struct {
	Map    [][]int
	Space  Coordinates
	Timer  int
	Robots []RobotSettings
}

func (ESHQD *EasterBunnyHeadquarters) parseRobotsData(inputMap []string) {
	ESHQD.Robots = []RobotSettings{}
	REGEX_DATA_PATTERN_ROBOT := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	for _, robotInputData := range inputMap {
		newRobot := RobotSettings{}
		robotData := REGEX_DATA_PATTERN_ROBOT.FindStringSubmatch(robotInputData)
		newRobot.Position.X, _ = strconv.Atoi(robotData[1])
		newRobot.Position.Y, _ = strconv.Atoi(robotData[2])
		newRobot.Velocity.X, _ = strconv.Atoi(robotData[3])
		newRobot.Velocity.Y, _ = strconv.Atoi(robotData[4])
		ESHQD.Robots = append(ESHQD.Robots, newRobot)
	}
}

func (robot *RobotSettings) moveRobot(Space Coordinates) {
	x := robot.Position.X + robot.Velocity.X
	y := robot.Position.Y + robot.Velocity.Y
	robot.Position.X = (x + Space.X) % Space.X
	robot.Position.Y = (y + Space.Y) % Space.Y
}

func (EBHQ EasterBunnyHeadquarters) printMap(timer int) {
	var printingSymbol string
	time.Sleep(time.Duration(0.1 * float64(time.Second)))
	for _, horizon := range EBHQ.Map {
		for _, vertice := range horizon {
			printingSymbol = "#"
			if vertice == 0 { printingSymbol = "."}
			fmt.Printf("%v", printingSymbol)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("TIMER: [%v]\n", timer)
}

func (EBHQ EasterBunnyHeadquarters) calcSafetyFactor() int {
	x0, y0 := EBHQ.Space.X/2, EBHQ.Space.Y/2
	quad_I, quad_II, quadIII, quad_IV := 0, 0, 0, 0
	for y, horizon := range EBHQ.Map {
		for x, vertice := range horizon {
			if x > x0 && y < y0 { quad_I += vertice; continue }
			if x < x0 && y < y0 { quad_II += vertice; continue }
			if x < x0 && y > y0 { quadIII += vertice; continue }
			if x > x0 && y > y0 { quad_IV += vertice; continue }
		}
	}
	safetyFactor := quad_I * quad_II * quadIII * quad_IV
	return safetyFactor
}

func PartOne(EBHQ EasterBunnyHeadquarters) int {
	for time := 0; time < EBHQ.Timer; time++ {
		currentEBHQMap := make([][]int, EBHQ.Space.Y)
		for i := range currentEBHQMap {
			currentEBHQMap[i] = make([]int, EBHQ.Space.X)
		}

		for idx := 0; idx < len(EBHQ.Robots); idx++ {
			robot := &EBHQ.Robots[idx]
			robot.moveRobot(EBHQ.Space)
			currentEBHQMap[robot.Position.Y][robot.Position.X]++
		}
		EBHQ.Map = currentEBHQMap
	}
	return EBHQ.calcSafetyFactor()
}

func PartTwo(EBHQ EasterBunnyHeadquarters) int {
	var timer int
	for timer = 0; timer > -1; timer++ {
		currentEBHQMap := make([][]int, EBHQ.Space.Y)
		for i := range currentEBHQMap {
			currentEBHQMap[i] = make([]int, EBHQ.Space.X)
		}

		for idx := 0; idx < len(EBHQ.Robots); idx++ {
			robot := &EBHQ.Robots[idx]
			robot.moveRobot(EBHQ.Space)
			currentEBHQMap[robot.Position.Y][robot.Position.X]++
		}
		EBHQ.Map = currentEBHQMap

		lineOfRobots := 0
		for _, horizon := range EBHQ.Map {
			newLine := 0
			for _, vertice := range horizon {
				if vertice != 0 {
					newLine++
					continue
				}
				if lineOfRobots < newLine {
					lineOfRobots = newLine
				}
				newLine = 0
			}
		}
		if lineOfRobots > 10 {
            //EBHQ.printMap(timer)
            break
		}

	}
	return timer + 1
}

func main() {
	fmt.Printf("--- Day 14: Restroom Redoubt ---\n")
	var EBHQ EasterBunnyHeadquarters
	inputMap := fileparser.ReadFileLines("input", false)
	EBHQ.parseRobotsData(inputMap)
	EBHQ.Space.X, EBHQ.Space.Y, EBHQ.Timer = 101, 103, 100
	fmt.Printf("PART[1]: %v\n", PartOne(EBHQ))
    EBHQ.parseRobotsData(inputMap)
	fmt.Printf("PART[2]: %v\n", PartTwo(EBHQ))
}
