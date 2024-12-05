package problems

import (
	"math"
	"slices"
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202401"
)

func init() {
	solutions["2024_01_1"] = year2024Day01Problem1
}

// https://adventofcode.com/2024/day/1
func year2024Day01Problem1(input []string) string {
	var leftList, rightList []int

	for _, line := range input {
		values := strings.Split(line, "   ")

		valueLeft, _ := strconv.Atoi(values[0])
		valueRight, _ := strconv.Atoi(values[1])

		leftList = append(leftList, valueLeft)
		rightList = append(rightList, valueRight)
	}

	slices.SortFunc(leftList, sol.Cmp)
	slices.SortFunc(rightList, sol.Cmp)

	diff := 0
	for index := range leftList {
		diff += int(math.Abs(float64(leftList[index]) - float64(rightList[index])))
	}

	return strconv.Itoa(diff)
}
