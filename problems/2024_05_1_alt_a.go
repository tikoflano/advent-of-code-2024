package problems

import (
	"strconv"
	"strings"
	"sync"
	sol "tikoflano/aoc/problems/lib/solution202405"
)

func init() {
	solutions["2024_05_1_alternative_a"] = year2024Day05Problem1AlternativeA
}

// https://adventofcode.com/2024/day/5
func year2024Day05Problem1AlternativeA(input []string) string {
	resp := 0
	rules, lineNumber := sol.CreateRules(input)
	c := make(chan int)

	// Validate updates
	var wg sync.WaitGroup
	for _, line := range input[lineNumber:] {
		update := strings.Split(line, ",")

		wg.Add(1)
		go sol.CheckUpdate(update, rules, c, &wg)
	}

	/**
		Wait for all goroutines to finish and close the channel

		IMPORTANT: if this is not done in a goroutine, the `c <-` calls will be blocked
					because no one is reading from the channel
	**/
	go sol.WaitAndClose(c, &wg)
	for val := range c {
		resp += val
	}

	return strconv.Itoa(resp)
}
