package problems

import (
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202405"
)

func init() {
	solutions["2024_05_2"] = year2024Day05Problem2
}

// https://adventofcode.com/2024/day/5#part2
func year2024Day05Problem2(input []string) string {
	resp := 0
	rules, lineNumber := sol.CreateRules(input)

	// Fix updates
	for _, line := range input[lineNumber:] {
		update := strings.Split(line, ",")

		if !sol.CheckCorrectOrder(update, rules) {
			sol.FixOrder(update, rules)
			resp += sol.GetMiddleValue(update)
		}
	}

	return strconv.Itoa(resp)
}
