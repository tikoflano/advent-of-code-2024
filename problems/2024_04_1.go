package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202404"
)

func init() {
	solutions["2024_04_1"] = year2024Day04Problem1
}

// https://adventofcode.com/2024/day/4
func year2024Day04Problem1(input []string) string {
	resp := 0
	wordSearch := sol.NewWordSearch(input)
	const wordToSearch = "XMAS"

	for y, line := range input {
		for x, mapLetter := range line {
			if string(mapLetter) != "X" {
				continue
			}

			curPos := sol.Vector{X: x, Y: y}

			for _, direction := range sol.Directions {
				if wordSearch.CheckWordFromPositionAndDirection(wordToSearch, &curPos, direction) {
					resp++
				}
			}
		}
	}

	return strconv.Itoa(resp)
}
