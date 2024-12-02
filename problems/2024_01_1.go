package problems

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func cmp(a, b int) int { return a - b }

// https://adventofcode.com/2024/day/1
func Year2024Day1Problem1(input []string) string {
	var leftList, rightList []int

	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			break
		}

		values := strings.Split(line, "   ")

		valueLeft, _ := strconv.Atoi(values[0])
		valueRight, _ := strconv.Atoi(values[1])

		leftList = append(leftList, valueLeft)
		rightList = append(rightList, valueRight)
	}

	slices.SortFunc(leftList, cmp)
	slices.SortFunc(rightList, cmp)

	diff := 0
	for index := range leftList {
		diff += int(math.Abs(float64(leftList[index]) - float64(rightList[index])))
	}

	return strconv.Itoa(diff)
}
