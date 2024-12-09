package main

import (
	fileparser "aoc/utils"
	"fmt"
	"strconv"
)

type Data struct {
	free  bool
	size  int
	value int
}

type DiskFragmenter struct {
	diskMap         string
	diskDenseFormat []Data
}

func (d *Data) DeepCopy() Data { return Data{ free:  d.free, size:  d.size, value: d.value,}}

func (df *DiskFragmenter) DeepCopy() DiskFragmenter {
	newDiskDenseFormat := make([]Data, len(df.diskDenseFormat))
	for i, data := range df.diskDenseFormat { newDiskDenseFormat[i] = data.DeepCopy() }
	return DiskFragmenter{ diskMap: df.diskMap,diskDenseFormat: newDiskDenseFormat,}
}

func (df *DiskFragmenter) parseDenseFormat() {
	dataType := false
	dataValue := 0
	df.diskDenseFormat = make([]Data, len(df.diskMap))

	for idx := 0; idx < len(df.diskMap); idx++ {
		valueStr := string(df.diskMap[idx])
		dataSize, _ := strconv.Atoi(valueStr)

		df.diskDenseFormat[idx] = Data{
			free:  dataType,
			size:  dataSize,
			value: dataValue,
		}

		if !dataType {
			dataValue++
		}
		dataType = !dataType
	}
}

func (df *DiskFragmenter) defragmentByBlock() []int {
	var blocks []int

	idxEnd := len(df.diskDenseFormat) - 1
	for idxStart := 0; !(idxStart > idxEnd); idxStart++ {
		currentData := &df.diskDenseFormat[idxStart]

		if !currentData.free {
			for i := 0; i < currentData.size; i++ {
				blocks = append(blocks, currentData.value)
			}
			continue
		}

		freeSpaceBlocks := currentData.size

		for freeSpaceBlocks > 0 {
			endData := &df.diskDenseFormat[idxEnd]

			if endData.free {
				idxEnd--
				continue
			}

			for freeSpaceBlocks > 0 {
				blocks = append(blocks, endData.value)
				freeSpaceBlocks--
				endData.size--
				if endData.size == 0 {
					idxEnd--
					break
				}
			}
		}

	}

	return blocks
}

func (df *DiskFragmenter) defragmentByFile() []int {
	var blocks []int

	for idxFile := len(df.diskDenseFormat) - 1; idxFile > -1; idxFile-- {
		currentFile := df.diskDenseFormat[idxFile]
		if currentFile.size == 0 || currentFile.free {
			continue
		}
		for idx := 0; idx < len(df.diskDenseFormat) && idx < idxFile; idx++ {
			emptySpce := &df.diskDenseFormat[idx]
			if !emptySpce.free || emptySpce.size < currentFile.size {
				continue
			}

			temp := *emptySpce
			temp.size -= currentFile.size
			df.diskDenseFormat[idx] = temp
			df.diskDenseFormat[idxFile].value = -1
			df.diskDenseFormat = append(df.diskDenseFormat[:idx], append([]Data{currentFile}, df.diskDenseFormat[idx:]...)...)
			break
		}
	}

	for _, disk := range df.diskDenseFormat {
		for counter := 0; counter < disk.size; counter++ {
			if !disk.free && disk.value!= -1{
				blocks = append(blocks,disk.value)
			} else {
				blocks = append(blocks, 0)
			}
		}
	}

	return blocks
}

func PartOne(df DiskFragmenter) int {
	filesystemChecksum := 0
	for idx, block := range df.defragmentByBlock() { filesystemChecksum += (idx * block)}
	return filesystemChecksum
}

func PartTwo(df DiskFragmenter) int {
	filesystemChecksum := 0
	for idx, block := range df.defragmentByFile() { filesystemChecksum += (idx * block) }
	return filesystemChecksum
}

func main() {
	fmt.Printf("--- Day 9: Disk Fragmenter ---\n")
	var diskFragmenter DiskFragmenter
	diskFragmenter.diskMap = fileparser.ReadFileLines("input", false)[0]
	diskFragmenter.parseDenseFormat()
	fmt.Printf("PART[1] %v \n", PartOne(diskFragmenter.DeepCopy()))
	fmt.Printf("PART[2] %v \n", PartTwo(diskFragmenter.DeepCopy()))
}
