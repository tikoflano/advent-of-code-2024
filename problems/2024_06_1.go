package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202406"
)

func init() {
	solutions["2024_06_1"] = year2024Day06Problem1
}

// https://adventofcode.com/2024/day/6
func year2024Day06Problem1(input []string) string {
	resp := 0
	floorMap, pos := sol.NewFloorMap(input)
	dir := sol.Directions.N

	for {
		if curCell := floorMap.GetCell(pos); !curCell.IsVisited() {
			resp++
			curCell.AddVisit(dir)
		}

		targetCell, err := floorMap.Look(pos, dir)
		if err != nil {
			break
		}

		if targetCell.Free {
			pos = targetCell.Pos
		} else {
			dir = dir.Rotate90()
		}
	}

	return strconv.Itoa(resp)
}