package solution202405

import (
	"slices"
	"strconv"
	"strings"
)

type Rules map[string][]string

func CreateRules(input []string) (Rules, int) {
	rules := Rules{}
	lineNumber := 0

	// Parse rules
	for _, line := range input {
		lineNumber++
		if line == "" {
			break
		}

		ruleDef := strings.Split(line, "|")

		if v, ok := rules[ruleDef[1]]; !ok {
			rules[ruleDef[1]] = []string{ruleDef[0]}
		} else {
			rules[ruleDef[1]] = append(v, ruleDef[0])
		}
	}

	return rules, lineNumber
}

func CheckCorrectOrder(update []string, rules map[string][]string) bool {
	for i := 0; i < len(update); i++ {
		updateVal := update[i]
		if ruleSet, ok := rules[updateVal]; ok {
			for j := i + 1; j < len(update); j++ {
				checkVal := update[j]
				if slices.Contains(ruleSet, checkVal) {
					return false
				}
			}
		}
	}

	return true
}

func FixOrder(update []string, rules map[string][]string) {
	for i := 0; i < len(update); i++ {
		updateVal := update[i]
		if ruleSet, ok := rules[updateVal]; ok {
			for j := i + 1; j < len(update); j++ {
				checkVal := update[j]
				if slices.Contains(ruleSet, checkVal) {
					// Swap i, j
					update[i], update[j] = update[j], update[i]
					i--
					break
				}
			}
		}
	}
}

func GetMiddleValue(s []string) int {
	middleVal := s[len(s)/2]
	val, _ := strconv.Atoi(middleVal)

	return val
}
