package problems

import (
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202405"
)

func init() {
	solutions["2024_05_1"] = year2024Day05Problem1
}

// https://adventofcode.com/2024/day/5
func year2024Day05Problem1(input []string) string {
	resp := 0
	rules, lineNumber := sol.CreateRules(input)

	// Validate updates
	for _, line := range input[lineNumber:] {
		update := strings.Split(line, ",")

		if sol.CheckCorrectOrder(update, rules) {
			resp += sol.GetMiddleValue(update)
		}
	}

	return strconv.Itoa(resp)
}
