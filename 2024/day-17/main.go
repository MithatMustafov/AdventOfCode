package main

import (
	fileparser "aoc/utils/fileparser"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ChronospatialComputer struct {
	RegisterA int
	RegisterB int
	RegisterC int

	Program []int
}

func (cc *ChronospatialComputer) parseMachinesData(input []string) {

	REGEX_PATTERN_DATA := regexp.MustCompile(`(Register (A|B|C): (\d+))`)
	idxLine := 0
	for ; idxLine < 3; idxLine++ {
		matches := REGEX_PATTERN_DATA.FindStringSubmatch(input[idxLine])
		if len(matches) == 4 {
			value, _ := strconv.Atoi(matches[3])
			switch matches[2] {
			case "A":
				cc.RegisterA = value
			case "B":
				cc.RegisterB = value
			case "C":
				cc.RegisterC = value
			}
		}
	}

	idxLine++
	cc.Program = make([]int, 0)
	REGEX_PATTERN_DATA = regexp.MustCompile(`Program: (\d+(?:,\d+)*)`)
	matches := REGEX_PATTERN_DATA.FindStringSubmatch(input[idxLine])
	strLiteralOperands := strings.Split(matches[1], ",")
	for _, literalOperand := range strLiteralOperands {
		num, _ := strconv.Atoi(literalOperand)
		cc.Program = append(cc.Program, num)
	}
}

func (cc *ChronospatialComputer) combo(operand int) int {
	if operand < 4 { return operand }
	if operand == 4 { return cc.RegisterA }
	if operand == 5 { return cc.RegisterB }
	return cc.RegisterC
}

func (cc *ChronospatialComputer) adv(operand int) { cc.RegisterA >>= cc.combo(operand) }
func (cc *ChronospatialComputer) bxl(operand int) { cc.RegisterB ^= operand }
func (cc *ChronospatialComputer) bst(operand int) { cc.RegisterB = cc.combo(operand)  & 0x7 }
func (cc *ChronospatialComputer) jnz(operand int, instructionPointer int) (int) { if cc.RegisterA != 0 { return operand - 2 }; return instructionPointer}
func (cc *ChronospatialComputer) bxc(operand int) { cc.RegisterB ^= cc.RegisterC }
func (cc *ChronospatialComputer) out(operand int) (int) { return cc.combo(operand)  & 0x7 }
func (cc *ChronospatialComputer) bdv(operand int) { cc.RegisterB = cc.RegisterA >> cc.combo(operand) }
func (cc *ChronospatialComputer) cdv(operand int) { cc.RegisterC = cc.RegisterA >> cc.combo(operand) }

func (cc *ChronospatialComputer) executeProgram() []int {
	output := make([]int, 0)
	var instructionPointer int
	for instructionPointer = 0; instructionPointer < int(len(cc.Program)); instructionPointer += 2 {
		opcode := cc.Program[instructionPointer]
		operand := int(cc.Program[instructionPointer+1])

		switch opcode {
		case 0: cc.adv(operand)
		case 1: cc.bxl(operand)
		case 2: cc.bst(operand)
		case 3: instructionPointer = cc.jnz(operand,instructionPointer)
		case 4: cc.bxc(operand)
		case 5: output = append(output, cc.out(operand))
		case 6: cc.bdv(operand)
		case 7: cc.cdv(operand)
		}
	}

	return output
}

func (cc ChronospatialComputer) printData() {
	fmt.Printf("---Data---\n")
	fmt.Printf("Register A: %v\n", cc.RegisterA)
	fmt.Printf("Register B: %v\n", cc.RegisterB)
	fmt.Printf("Register C: %v\n", cc.RegisterC)
	fmt.Printf("Program: %v\n", cc.Program)
}

func (cc ChronospatialComputer) getOutput( output []int) (string) {
	var strOutput string
	for i, out := range output {
		if i == len(output) - 1 { strOutput += strconv.Itoa(out)
		} else { strOutput += strconv.Itoa(out) + ","}
	}

	return strOutput
}

func PartOne(cc ChronospatialComputer) string {
	output := cc.executeProgram()
	return cc.getOutput(output)
}

func max(a, b int) int { if a > b { return a}; return b }

func isOutputCopyofProgram(a, b []int) bool {
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}

func isEndingWithMatch(output, program []int, length int) bool {
	if len(output) < length || len(program) < length {
		return false
	}
	return isOutputCopyofProgram(output[len(output)-length:], program[len(program)-length:])
}

func PartTwo(cc ChronospatialComputer) int {
	cc.RegisterA = int(math.Pow(8, float64(len(cc.Program)-1)))
	power := len(cc.Program)-2
	matchedLength := 1

	for {
		cc.RegisterA += int(math.Pow(8, float64(power)))

		newCC := ChronospatialComputer{
			RegisterA: cc.RegisterA,
			RegisterB: cc.RegisterB,
			RegisterC: cc.RegisterC,
			Program:   append([]int{}, cc.Program...),
		}

		output := newCC.executeProgram()

		if isEndingWithMatch(output, cc.Program, matchedLength) {
			power = max(0, power-1)
			matchedLength++
		}

		if isEndingWithMatch(output, cc.Program, len(cc.Program)) { break }
	}

	return cc.RegisterA
}

func main() {
	fmt.Printf("--- Day 17: Chronospatial Computer ---\n")
	var chronospatialComputer ChronospatialComputer
	inputMap := fileparser.ReadFileLines("input", false)
	chronospatialComputer.parseMachinesData(inputMap)
	fmt.Printf("PART[1]: %v\n", PartOne(chronospatialComputer))
	chronospatialComputer.parseMachinesData(inputMap)
	fmt.Printf("PART[2]: %v\n", PartTwo(chronospatialComputer))
}
