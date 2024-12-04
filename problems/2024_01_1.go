package problems

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"tikoflano/aoc/problems/lib/solution202401"
)

func init() {
	solutions["2024_01_1"] = year2024Day1Problem1
}

// https://adventofcode.com/2024/day/1
func year2024Day1Problem1(input []string) string {
	var leftList, rightList []int

	for _, line := range input {
		values := strings.Split(line, "   ")

		valueLeft, _ := strconv.Atoi(values[0])
		valueRight, _ := strconv.Atoi(values[1])

		leftList = append(leftList, valueLeft)
		rightList = append(rightList, valueRight)
	}

	slices.SortFunc(leftList, solution202401.Cmp)
	slices.SortFunc(rightList, solution202401.Cmp)

	diff := 0
	for index := range leftList {
		diff += int(math.Abs(float64(leftList[index]) - float64(rightList[index])))
	}

	return strconv.Itoa(diff)
}
