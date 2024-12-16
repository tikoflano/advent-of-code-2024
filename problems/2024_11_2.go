package problems

import (
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202411"
)

func init() {
	solutions["2024_11_2"] = year2024Day11Problem2
}

// https://adventofcode.com/2024/day/11#part2
func year2024Day11Problem2(input []string) string {
	var resp int
	values := strings.Split(input[0], " ")
	times := 75

	for _, value := range values {
		node := &sol.Node{
			Value: value,
		}
		node.Blink(times)
		// resp += node.ChildrenAtLevel(times)
	}

	return strconv.Itoa(resp)
}
