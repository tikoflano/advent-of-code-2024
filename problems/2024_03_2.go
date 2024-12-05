package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202403"
)

func init() {
	solutions["2024_03_2"] = year2024Day3Problem2
}

// https://adventofcode.com/2024/day/3#part2
func year2024Day3Problem2(input []string) string {
	resp := 0
	parser := sol.NewParser()
	executor := sol.NewExecutor()

	for _, line := range input {
		for _, char := range line {
			parser.NextChar(string(char))
		}
	}

	resp += executor.RunProgram(parser.Operations)

	return strconv.Itoa(resp)
}
