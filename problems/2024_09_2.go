package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202409"
)

func init() {
	solutions["2024_09_2"] = year2024Day09Problem2
}

// https://adventofcode.com/2024/day/9#part2
func year2024Day09Problem2(input []string) string {
	diskMap := sol.NewDiskMap(input[0])
	diskMap.Decode()
	diskMap.DefragmentCompleteFiles()

	resp := diskMap.Checksum()

	return strconv.Itoa(resp)
}
