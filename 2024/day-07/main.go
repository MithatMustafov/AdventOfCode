package main

import (
	fileparser "aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CalibrationEquations struct {
	equations  []Equation
	operations []string
}

type Equation struct {
	value    int
	numbers  []int
	variants []string
}

func (calibrationEquations *CalibrationEquations) parseInput(INPUT *[]string) {
	var newCalibrationEquations CalibrationEquations

	for _, line := range *INPUT {
		equation := strings.Split(line, ":")

		value, _ := strconv.Atoi(strings.TrimSpace(equation[0]))

		numbersStr := strings.Split(strings.TrimSpace(equation[1]), " ")
		numbers := make([]int, len(numbersStr))

		for i := 0; i < len(numbersStr); i++ {
			num, _ := strconv.Atoi(strings.TrimSpace(numbersStr[i]))
			numbers[i] = num
		}
		newEquation := Equation{
			value:   value,
			numbers: numbers,
		}

		newCalibrationEquations.equations = append(newCalibrationEquations.equations, newEquation)
	}

	*calibrationEquations = newCalibrationEquations
}

func (equation *Equation) generateVariants(OPERATORS []string) {
	numbersCount := len(equation.numbers)
	equation.variants = make([]string, 0)

	var generate func(int, string)
	generate = func(idx int, current string) {
		if idx == numbersCount {
			equation.variants = append(equation.variants, current)
			return
		}

		if idx > 0 {
			for _, operator := range OPERATORS {
				generate(idx+1, current+operator+fmt.Sprint(equation.numbers[idx]))
			}
		} else {
			generate(idx+1, fmt.Sprint(equation.numbers[idx]))
		}
	}

	generate(1, fmt.Sprint(equation.numbers[0]))
}

func (equation *Equation) calc(calcEquation string) int {
	re := regexp.MustCompile(`(\d+)|([+*|])`)
	matches := re.FindAllString(calcEquation, -1)
	calcAnswer, _ := strconv.Atoi(matches[0])

	for i := 1; i < len(matches); i += 1 {
		operator := matches[i-1]
		number, _ := strconv.Atoi(matches[i])

		if operator == "+" {
			calcAnswer += number
		} else if operator == "*" {
			calcAnswer *= number
		} else if operator == "|" {
			calcAnswer, _ = strconv.Atoi((strconv.Itoa(calcAnswer) + strconv.Itoa(number)))
		}
	}
    
	return calcAnswer
}

func (calibrationEquations *CalibrationEquations) partAnswer() int {
	sum := 0
	for _, equation := range calibrationEquations.equations {
		equation.generateVariants(calibrationEquations.operations)
		for _, equationVariant := range equation.variants {
			if equation.value == equation.calc(equationVariant) {
				sum += equation.value
				break
			}
		}
	}
	return sum
}

func main() {
	input := fileparser.ReadFileLines("input", false)
	var calibrationEquations CalibrationEquations
	calibrationEquations.parseInput(&input)
	fmt.Printf("--- Day 7: Bridge Repair ---\n")
	calibrationEquations.operations = []string{"+", "*"}
	fmt.Printf("PART[1] %v \n", calibrationEquations.partAnswer())
	calibrationEquations.operations = []string{"+", "*", "|"}
	fmt.Printf("PART[1] %v \n", calibrationEquations.partAnswer())
}
