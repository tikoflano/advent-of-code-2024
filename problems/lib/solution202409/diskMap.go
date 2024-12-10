package solution202409

import (
	"strconv"
)

type DiskMap struct {
	Input  string
	Memory []Block
}

func NewDiskMap(input string) *DiskMap {
	return &DiskMap{
		Input:  input,
		Memory: []Block{},
	}
}

func (diskMap *DiskMap) Decode() {
	fileParser := makeFileParser()
	freeMemoryParser := makeFreeMemoryParser()

	blockParsers := []BlockParser{fileParser, freeMemoryParser}

	var parserIndex int
	for _, char := range diskMap.Input {
		blockParser := blockParsers[parserIndex]

		lenght, _ := strconv.Atoi(string(char))
		diskMap.Memory = append(diskMap.Memory, blockParser(lenght)...)

		parserIndex = (parserIndex + 1) % len(blockParsers)
	}
}

func (diskMap *DiskMap) DefragmentPartialFiles() {
	p1 := 0
	p2 := len(diskMap.Memory) - 1

	for {
		var dest, origin Block

		for ; p1 < len(diskMap.Memory); p1++ {
			origin = diskMap.Memory[p1]
			if origin.IsFree() {
				break
			}
		}

		for ; p2 > 0; p2-- {
			dest = diskMap.Memory[p2]
			if !dest.IsFree() {
				break
			}
		}

		if p1 > p2 {
			return
		}

		diskMap.Memory[p1], diskMap.Memory[p2] = diskMap.Memory[p2], diskMap.Memory[p1]
	}
}

func (diskMap *DiskMap) DefragmentCompleteFiles() {
	p1 := 0
	p2 := len(diskMap.Memory) - 1

	for p2 > 0 {
		var dest, origin Block

		dest = diskMap.Memory[p2]
		if dest.IsFree() {
			p2--
			continue
		}

		fileLength, fileStart := diskMap.getLengthFromEnd(p2)

		for p1 < len(diskMap.Memory) {
			origin = diskMap.Memory[p1]
			if origin.IsFree() {
				freeMemoryLength, _ := diskMap.getLengthFromStart(p1)
				if freeMemoryLength >= fileLength {
					break
				}
				p1 += freeMemoryLength
			}
			p1++
		}

		if p1 < p2 {
			// Swap free blocks with full file
			for i := 0; i < fileLength; i++ {
				diskMap.Memory[p1+i], diskMap.Memory[fileStart+i] = diskMap.Memory[fileStart+i], diskMap.Memory[p1+i]
			}
		}

		p1 = 0
		p2 -= fileLength
	}
}

func (diskMap *DiskMap) getLengthFromEnd(pos int) (length, start int) {
	blockId := diskMap.Memory[pos].GetId()

	for ; pos > 0 && diskMap.Memory[pos].GetId() == blockId; pos-- {
		length++
		start = pos
	}

	return
}

func (diskMap *DiskMap) getLengthFromStart(pos int) (length, end int) {
	blockId := diskMap.Memory[pos].GetId()

	for ; pos < len(diskMap.Memory) && diskMap.Memory[pos].GetId() == blockId; pos++ {
		length++
		end = pos
	}

	return
}

func (diskMap *DiskMap) Checksum() (ret int) {
	for index, block := range diskMap.Memory {
		ret += index * block.GetId()
	}

	return
}
