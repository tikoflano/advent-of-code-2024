package problems

import (
	"regexp"
	"strconv"
)

func init() {
	solutions["2024_03_1"] = year2024Day03Problem1
}

// https://adventofcode.com/2024/day/3
func year2024Day03Problem1(input []string) string {
	resp := 0
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			resp += num1 * num2
		}
	}

	return strconv.Itoa(resp)
}
