package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution202406"
)

func init() {
	solutions["2024_06_2"] = year2024Day06Problem2
}

// https://adventofcode.com/2024/day/6#part2
func year2024Day06Problem2(input []string) string {
	var resp int
	floorMap, pos := sol.NewFloorMap(input)
	dir := sol.Directions.N

	for {
		if curCell := floorMap.GetCell(pos); !curCell.IsVisited() {
			curCell.AddVisit(dir)
		}

		targetCell, err := floorMap.Look(pos, dir)
		if err != nil {
			break
		}

		if targetCell.Free {
			// Imagine there is an obstacle, check if that will put us back in an already visited cell facing the same direction
			fakeDir := dir.Rotate90()
			if floorMap.CheckVisitedInPath(pos, fakeDir) {
				resp++
			}
			pos = targetCell.Pos
		} else {
			dir = dir.Rotate90()
		}
	}

	return strconv.Itoa(resp)
}
