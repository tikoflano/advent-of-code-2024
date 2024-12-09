package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202408"
)

func init() {
	solutions["2024_08_2"] = year2024Day08Problem2
}

// https://adventofcode.com/2024/day/8#part2
func year2024Day08Problem2(input []string) string {
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
				newAntinodes := sol.GetAntinodesWithHarmonics(frequencyCell, currentCell, len(line), len(input))

				for _, newAntinode := range newAntinodes {
					antinodes = antinodes.AddAntinode(newAntinode)
				}
			}

			if len(frequencyMap[frequency]) == 1 {
				antinodes = antinodes.AddAntinode(frequencyMap[frequency][0])
			}
			antinodes = antinodes.AddAntinode(currentCell)

			frequencyMap[frequency] = append(frequencyMap[frequency], currentCell)

		}
	}

	return strconv.Itoa(len(antinodes))
}
