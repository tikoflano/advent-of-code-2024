package problems

import (
	"strconv"
	lib "tikoflano/aoc/problems/lib/solution202404"
)

func init() {
	solutions["2024_04_1"] = year2024Day4Problem1
}

// https://adventofcode.com/2024/day/4
func year2024Day4Problem1(input []string) string {
	resp := 0

	for y, line := range input {
		for x, mapLetter := range line {
			if string(mapLetter) != "X" {
				continue
			}

			curPos := lib.Vector{X: x, Y: y}

			for _, direction := range lib.Directions {
				if lib.CheckFromPositionAndDirection("XMAS", input, &curPos, direction) {
					resp++
				}
			}
		}
	}

	return strconv.Itoa(resp)
}
