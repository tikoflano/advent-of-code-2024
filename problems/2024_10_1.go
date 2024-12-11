package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202410"
)

func init() {
	solutions["2024_10_1"] = year2024Day10Problem1
}

// https://adventofcode.com/2024/day/10
func year2024Day10Problem1(input []string) string {
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
		score, _ := trailhead.GetMeasurements()
		resp += score
	}

	return strconv.Itoa(resp)
}
