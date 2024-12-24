package main

import (
	"aoc/utils/fileparser"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type GateType string

const (
	AND GateType = "AND"
	OR  GateType = "OR"
	XOR GateType = "XOR"
)

type GateConnection struct {
	WireA string
	WireB string
	GateType GateType
	Result string
}

type MonitoringDevice struct {
	Wires map[string]bool
	GateConnections  []GateConnection
}

func (md *MonitoringDevice) parseData(input []string) {
	md.Wires = make(map[string]bool)
	md.GateConnections = make([]GateConnection, 0)
	
	var idx int
	for idx = 0; input[idx] !=  ""; idx++ {
		line := input[idx]
		initialWireValues := strings.Split(line,": ")
		wireName := initialWireValues[0]
		wireValue,_ := strconv.Atoi(initialWireValues[1])
		md.Wires[wireName] = wireValue != 0
	}
	idx++
	var newGateConnection GateConnection
	for ; idx < len(input); idx++ {
		line := input[idx]
		gateConnection := strings.Split(line," ")
		newGateConnection.WireA = gateConnection[0]
		newGateConnection.GateType  = GateType(gateConnection[1])
		newGateConnection.WireB = gateConnection[2]
		newGateConnection.Result = gateConnection[4]
		md.GateConnections = append(md.GateConnections, newGateConnection)
	}
}

func(md *MonitoringDevice) isWireDecoded(wire string) (bool) { _, exists := md.Wires[wire]; return exists }
func(md *MonitoringDevice) getWireValue(wire string) (bool) { value := md.Wires[wire]; return value }

func (md *MonitoringDevice) decodeWires() {
	DECODE_AGAIN:
	isWiresDecoded := true
	for _, gateConnection := range md.GateConnections {
		if md.isWireDecoded(gateConnection.Result) { continue }

		if(!md.isWireDecoded(gateConnection.WireA) || !md.isWireDecoded(gateConnection.WireB)) { 
			isWiresDecoded = false;
			continue 
		}

		md.Wires[gateConnection.Result] = logicGate(
			gateConnection.GateType,
			md.getWireValue(gateConnection.WireA),
			md.getWireValue(gateConnection.WireB))
	}
	if(!isWiresDecoded) {goto DECODE_AGAIN}
}

func logicGate(gateType GateType, a bool, b bool) bool {
	switch gateType {
	case AND:
		return a && b
	case OR:
		return a || b
	case XOR:
		return a != b
	default:
		return false
	}
}

func boolToString(b bool) string {
	if b { return "1" }
	return "0"
}

func PartOne(md MonitoringDevice) int64 {
	md.decodeWires()
	wires := make([]string, 0, len(md.Wires))
	for wire := range md.Wires { wires = append(wires, wire) }
	sort.Strings(wires)
	var binaryStr string
	for _, wire := range wires {
		if(wire[0] == 'z') { 
			binaryStr = boolToString(md.Wires[wire]) + binaryStr
		}
	}
	decimalCode, _ := strconv.ParseInt(binaryStr, 2, 0)
	return decimalCode
}

func PartTwo(md MonitoringDevice) int { return 0 }

func main() {
	fmt.Printf("--- Day 24: Crossed Wires ---\n")
	var monitoringDevice MonitoringDevice
	input := fileparser.ReadFileLines("input", false)
	monitoringDevice.parseData(input)
	fmt.Printf("PART[1]: %v\n", PartOne(monitoringDevice))
	fmt.Printf("PART[2]: %v\n", PartTwo(monitoringDevice))
}