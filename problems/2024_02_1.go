package problems

import (
	"math"
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202402"
)

func init() {
	solutions["2024_02_1"] = year2024Day02Problem1
}

// https://adventofcode.com/2024/day/2
func year2024Day02Problem1(input []string) string {
	safe := 0

	for _, line := range input {
		values := strings.Split(line, " ")

		var isAscending *bool
		prevValue, _ := strconv.Atoi(values[0])
		isSafe := true

		for _, value := range values[1:] {
			numericValue, _ := strconv.Atoi(value)
			diff := numericValue - prevValue
			absDiff := math.Abs(float64(diff))

			// Set direction if it is not already set
			if isAscending == nil {
				isAscending = sol.BoolPtr(diff > 0)
			}

			// Too low or high diff
			if absDiff < 1 || absDiff > 3 {
				isSafe = false
				break
			}

			// Direction change
			if (*isAscending && diff < 0) || (!*isAscending && diff > 0) {
				isSafe = false
				break
			}

			prevValue = numericValue
		}

		if isSafe {
			safe++
		}
	}

	return strconv.Itoa(safe)
}
