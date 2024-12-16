package problems

import (
	"strconv"
	"strings"
	sol "tikoflano/aoc/problems/lib/solution202411"
)

func init() {
	solutions["2024_11_1"] = year2024Day11Problem1
}

// https://adventofcode.com/2024/day/11
func year2024Day11Problem1(input []string) string {
	values := strings.Split(input[0], " ")
	times := 25

	head := &sol.Node1{
		Value: values[0],
	}

	currentNode := head
	for i := 1; i < len(values); i++ {
		next := &sol.Node1{
			Value: values[i],
		}

		currentNode.Next = next
		currentNode = next
	}

	var nextNode *sol.Node1
	currentNode = head
	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Blink(times)
		currentNode = nextNode
	}

	return strconv.Itoa(head.LengthToTail())
}
