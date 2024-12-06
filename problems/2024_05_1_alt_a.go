package problems

import (
	"strconv"
	"strings"
	"sync"
	lib "tikoflano/aoc/problems/lib/solution202405"
)

func init() {
	solutions["2024_05_1_alternative_a"] = year2024Day05Problem1AlternativeA
}

func checkUpdate(update []string, rules lib.Rules, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if lib.CheckCorrectOrder(update, rules) {
		c <- lib.GetMiddleValue(update)
	}
}

func waitAndClose(c chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(c)
}

// https://adventofcode.com/2024/day/5
func year2024Day05Problem1AlternativeA(input []string) string {
	resp := 0
	rules, lineNumber := lib.CreateRules(input)
	c := make(chan int)

	// Validate updates
	var wg sync.WaitGroup
	for _, line := range input[lineNumber:] {
		update := strings.Split(line, ",")

		wg.Add(1)
		go checkUpdate(update, rules, c, &wg)
	}

	/**
		Wait for all goroutines to finish and close the channel

		IMPORTANT: if this is not done in a goroutine, the `c <-` calls will be blocked
					because no one is reading from the channel
	**/
	go waitAndClose(c, &wg)
	for val := range c {
		resp += val
	}

	return strconv.Itoa(resp)
}
