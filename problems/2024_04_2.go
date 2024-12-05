package problems

import (
	"strconv"
	lib "tikoflano/aoc/problems/lib/solution202404"
)

func init() {
	solutions["2024_04_2"] = year2024Day4Problem2
}

// https://adventofcode.com/2024/day/4#part2
func year2024Day4Problem2(input []string) string {
	resp := 0

	for y, line := range input {
		for x, mapLetter := range line {
			if string(mapLetter) != "A" {
				continue
			}

			curPos := lib.Vector{X: x, Y: y}
			nwPos := curPos.Add(lib.Directions["NW"])
			nePos := curPos.Add(lib.Directions["NE"])
			swPos := curPos.Add(lib.Directions["SW"])
			sePos := curPos.Add(lib.Directions["SE"])

			nwLetter, err := lib.GetMapLetter(input, nwPos)
			if err != nil {
				continue
			}
			neLetter, err := lib.GetMapLetter(input, nePos)
			if err != nil {
				continue
			}
			swLetter, err := lib.GetMapLetter(input, swPos)
			if err != nil {
				continue
			}
			seLetter, err := lib.GetMapLetter(input, sePos)
			if err != nil {
				continue
			}

			word1 := nwLetter + "A" + seLetter
			word2 := neLetter + "A" + swLetter

			if lib.IsMas(word1) && lib.IsMas(word2) {
				resp++
			}
		}
	}

	return strconv.Itoa(resp)
}
