package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202409"
)

func init() {
	solutions["2024_09_1"] = year2024Day09Problem1
}

// https://adventofcode.com/2024/day/9
func year2024Day09Problem1(input []string) string {
	diskMap := sol.NewDiskMap(input[0])
	diskMap.Decode()
	diskMap.DefragmentPartialFiles()

	resp := diskMap.Checksum()

	return strconv.Itoa(resp)
}
