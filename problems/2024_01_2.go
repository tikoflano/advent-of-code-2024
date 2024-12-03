package problems

import (
	"strconv"
	"strings"
)

func init() {
	solutions["2024_01_2"] = year2024Day1Problem2
}

// https://adventofcode.com/2024/day/1#part2
func year2024Day1Problem2(input []string) string {
	var leftList, rightList []int

	for _, line := range input {
		values := strings.Split(line, "   ")

		valueLeft, _ := strconv.Atoi(values[0])
		valueRight, _ := strconv.Atoi(values[1])

		leftList = append(leftList, valueLeft)
		rightList = append(rightList, valueRight)
	}

	countMap := make(map[int]int)

	for _, valueRight := range rightList {
		if _, exists := countMap[valueRight]; exists {
			countMap[valueRight] = countMap[valueRight] + 1
		} else {
			countMap[valueRight] = 1
		}
	}

	result := 0

	for _, valueLeft := range leftList {
		if count, exists := countMap[valueLeft]; exists {
			result += valueLeft * count
		}
	}

	return strconv.Itoa(result)
}
