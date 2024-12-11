package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202410"
)

func init() {
	solutions["2024_10_2"] = year2024Day10Problem2
}

// https://adventofcode.com/2024/day/10#part2
func year2024Day10Problem2(input []string) string {
	var resp int
	trailheads := []sol.Step{}

	for y, line := range input {
		for x, cell := range line {
			level, _ := strconv.Atoi(string(cell))
			if level == 0 {
				trailhead := sol.Step{
					Pos:   sol.Vector{X: x, Y: y},
					Level: 0,
				}

				trailhead.FindTrails(input)
				trailheads = append(trailheads, trailhead)
			}
		}
	}

	for _, trailhead := range trailheads {
		_, ratings := trailhead.GetMeasurements()
		resp += ratings
	}

	return strconv.Itoa(resp)
}
