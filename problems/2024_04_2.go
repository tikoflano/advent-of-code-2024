package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202404"
)

func init() {
	solutions["2024_04_2"] = year2024Day04Problem2
}

// https://adventofcode.com/2024/day/4#part2
func year2024Day04Problem2(input []string) string {
	resp := 0
	wordSearch := sol.NewWordSearch(input)

	for y, line := range input {
		for x, mapLetter := range line {
			if string(mapLetter) != "A" {
				continue
			}

			curPos := sol.Vector{X: x, Y: y}
			nwPos := curPos.Add(sol.Directions["NW"])
			nePos := curPos.Add(sol.Directions["NE"])
			swPos := curPos.Add(sol.Directions["SW"])
			sePos := curPos.Add(sol.Directions["SE"])

			nwLetter, err := wordSearch.GetLetterAt(nwPos)
			if err != nil {
				continue
			}
			neLetter, err := wordSearch.GetLetterAt(nePos)
			if err != nil {
				continue
			}
			swLetter, err := wordSearch.GetLetterAt(swPos)
			if err != nil {
				continue
			}
			seLetter, err := wordSearch.GetLetterAt(sePos)
			if err != nil {
				continue
			}

			word1 := nwLetter + "A" + seLetter
			word2 := neLetter + "A" + swLetter

			if sol.IsMas(word1) && sol.IsMas(word2) {
				resp++
			}
		}
	}

	return strconv.Itoa(resp)
}
