package problems

import (
	"slices"
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202402"
)

func init() {
	solutions["2024_02_2"] = year2024Day02Problem2
}

// https://adventofcode.com/2024/day/2#part2
func year2024Day02Problem2(input []string) string {
	safe := 0

	for _, line := range input {
		values := strings.Split(line, " ")

		if sol.ValidateSafeLevels(values) {
			safe++
			continue
		}

		for i := 0; i < len(values); i++ {
			if sol.ValidateSafeLevels(slices.Concat(values[:i], values[i+1:])) {
				safe++
				break
			}
		}
	}

	return strconv.Itoa(safe)
}
