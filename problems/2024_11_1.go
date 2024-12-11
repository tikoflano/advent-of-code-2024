package problems

import (
	"strconv"
	"strings"
	// sol "tikoflano/aoc/problems/lib/solution202411"
)

func init() {
	solutions["2024_11_1"] = year2024Day11Problem1
}

// https://adventofcode.com/2024/day/11
func year2024Day11Problem1(input []string) string {
	values := strings.Split(input[0], " ")
	blinks := 25

	for blink := 0; blink < blinks; blink++ {
		for i := 0; i < len(values); i++ {
			cur := values[i]

			if cur == "0" {
				values[i] = "1"
			} else if len(cur)%2 == 0 {
				left := cur[:(len(cur) / 2)]
				right := cur[(len(cur) / 2):]

				for len(right) > 1 && strings.HasPrefix(right, "0") {
					right = strings.TrimPrefix(right, "0")
				}

				newElements := []string{left, right}
				values = append(values[:i], append(newElements, values[i+1:]...)...)

				i++
			} else {
				value, _ := strconv.Atoi(cur)
				value *= 2024

				newElements := []string{strconv.Itoa(value)}
				values = append(values[:i], append(newElements, values[i+1:]...)...)
			}
		}
	}

	return strconv.Itoa(len(values))
}
