package main

import (
	fileparser "aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func popByIdx(stack []int, idx int) []int {
	if idx < 0 || idx >= len(stack) { return stack }
	stack = append(stack[:idx], stack[idx+1:]...)
	return stack
}

func isReportValid(report []int) (bool, int) {
	const DIFF_MIN, DIFF_MAX = 1, 3
	checkChange := report[1] - report[0]
	action := checkChange > 0
	isValid := true

	idxLevel := 0;
	for ; idxLevel < len(report)-1; idxLevel++ {
		difference := report[idxLevel] - report[idxLevel+1]
		diffAbs := int(math.Abs(float64(difference)))
		if diffAbs < DIFF_MIN || diffAbs > DIFF_MAX {isValid = false; break;}

		checkChange := report[idxLevel+1] - report[idxLevel]
		tempAction := checkChange > 0
		if tempAction != action {isValid = false; break}
	}

	return isValid, idxLevel
}

func reportLineToArr(line string) []int{
	strReport := strings.Split(line, " ")
	report := make([]int, len(strReport))
	for idxLevel, level := range strReport {
		report[idxLevel], _ = strconv.Atoi(level)
	}

	return report
}

func partOne(lines []string) {
	var sumSAFE int
	for _, line := range lines {
		report := reportLineToArr(line)
		isValid, _ := isReportValid(report)
		if isValid {sumSAFE++}
	}
	fmt.Printf("%v\n", sumSAFE)
}

type UnsafeReports struct {
    ReportError []int
    ReportErrorLeft []int
	ReportErrorRight []int
}

func partTwo(lines []string) {
	var sumSAFE int
	var newUnsafeReports []UnsafeReports
	for _, line := range lines {
		report := reportLineToArr(line)
		isValid, idx := isReportValid(report)
		if isValid {
			sumSAFE++
		} else {
			newUnsafeReports = append(newUnsafeReports, UnsafeReports{
				ReportError: popByIdx(append([]int(nil), report...), idx),
				ReportErrorLeft: popByIdx(append([]int(nil), report...), idx-1),
				ReportErrorRight: popByIdx(append([]int(nil), report...), idx+1),
			})
		}
	}

	sumUNSAFE := 0
	for _, report := range newUnsafeReports {
		isValidReportError,_ := isReportValid(report.ReportError)
		isValidReportErrorLeft,_ := isReportValid(report.ReportErrorLeft)
		isValidReportErrorRight,_ := isReportValid(report.ReportErrorRight)
		if isValidReportError || isValidReportErrorLeft || isValidReportErrorRight {sumUNSAFE++}
	}

	fmt.Printf("%v\n", (sumSAFE+ sumUNSAFE));
}

func main() {
	lines := fileparser.ReadFileLines("input", false)
	fmt.Printf("--- Day 2: Red-Nosed Reports ---\n")
	fmt.Printf("PART[1] "); partOne(lines)
	fmt.Printf("PART[2] "); partTwo(lines)
}