package solution202402

import (
	"math"
	"strconv"
)

func BoolPtr(b bool) *bool {
	return &b
}

func ValidateSafeLevels(levels []string) bool {
	var isAscending *bool
	prevLevel, _ := strconv.Atoi(levels[0])

	for i := 1; i < len(levels); i++ {
		level := levels[i]
		numericLevel, _ := strconv.Atoi(level)
		diff := numericLevel - prevLevel
		absDiff := math.Abs(float64(diff))

		// Set direction if it is not already set
		if isAscending == nil {
			isAscending = BoolPtr(diff > 0)
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
