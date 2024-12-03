package problems

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func init() {
	solutions["2024_02_2"] = year2024Day2Problem2
}

// https://adventofcode.com/2024/day/2#part2
func year2024Day2Problem2(input []string) string {
	safe := 0

	for _, line := range input {
		values := strings.Split(line, " ")

		if validateSafeLevels(values) {
			safe++
			continue
		}

		for i := 0; i < len(values); i++ {
			if validateSafeLevels(slices.Concat(values[:i], values[i+1:])) {
				safe++
				break
			}
		}
	}

	return strconv.Itoa(safe)
}

func validateSafeLevels(levels []string) bool {
	var isAscending *bool
	prevLevel, _ := strconv.Atoi(levels[0])

	for i := 1; i < len(levels); i++ {
		level := levels[i]
		numericLevel, _ := strconv.Atoi(level)
		diff := numericLevel - prevLevel
		absDiff := math.Abs(float64(diff))

		// Set direction if it is not already set
		if isAscending == nil {
			isAscending = boolPtr(diff > 0)
		}

		// Too low or high diff
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		// Direction change
		if (*isAscending && diff < 0) || (!*isAscending && diff > 0) {
			return false
		}

		prevLevel = numericLevel
	}

	return true
}
