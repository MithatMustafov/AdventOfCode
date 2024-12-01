package fileparser

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLines(fName string, isPrinting bool) []string {
	var lines []string
	file, _ := os.Open(fName + ".txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	if isPrinting {
		for _, line := range lines {
			fmt.Printf("%v\n", line)
		}
	}
	return lines
}
