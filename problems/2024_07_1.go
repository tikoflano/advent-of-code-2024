package problems

import (
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202407"
)

func init() {
	solutions["2024_07_1"] = year2024Day07Problem1
}

// https://adventofcode.com/2024/day/7
func year2024Day07Problem1(input []string) string {
	var resp int

	for _, line := range input {
		elements := strings.Split(line, " ")
		targetResult, _ := strconv.Atoi(strings.TrimSuffix(elements[0], ":"))
		tempOperands := elements[1:]
		var operands []int

		for _, tempOperand := range tempOperands {
			operandValue, _ := strconv.Atoi(tempOperand)
			operands = append(operands, operandValue)
		}

		if sol.ValidateEquation(operands, targetResult) {
			resp += targetResult
		}
	}

	return strconv.Itoa(resp)
}
