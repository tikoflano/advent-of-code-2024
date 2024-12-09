package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202408"
)

func init() {
	solutions["2024_08_1"] = year2024Day08Problem1
}

// https://adventofcode.com/2024/day/8
func year2024Day08Problem1(input []string) string {
	var antinodes sol.Antinodes
	frequencyMap := make(map[rune]sol.Antinodes)

	for y, line := range input {
		for x, frequency := range line {
			if frequency == '.' {
				continue
			}

			currentCell := &sol.Vector{X: x, Y: y}

			frequencyCells, ok := frequencyMap[frequency]
			if !ok {
				frequencyMap[frequency] = sol.Antinodes{currentCell}
				continue
			}

			for _, frequencyCell := range frequencyCells {
				newAntinodes := sol.GetAntinodesWithNoHarmonics(frequencyCell, currentCell)
				newAntinodes = newAntinodes.FilterAntinodes(len(line), len(input))

				for _, newAntinode := range newAntinodes {
					antinodes = antinodes.AddAntinode(newAntinode)
				}
			}

			frequencyMap[frequency] = append(frequencyMap[frequency], currentCell)
		}
	}

	return strconv.Itoa(len(antinodes))
}
