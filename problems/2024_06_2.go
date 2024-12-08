package problems

import (
	"strconv"
	"sync"
	sol "tikoflano/aoc/problems/lib/solution202406"
)

func init() {
	solutions["2024_06_2"] = year2024Day06Problem2
}

// https://adventofcode.com/2024/day/6#part2
func year2024Day06Problem2(input []string) string {
	var resp int
	c := make(chan int)
	var wg sync.WaitGroup
	floorMap, startingPos := sol.NewFloorMap(input)
	startingDir := sol.Directions.N

	// Tried a bunch of things that didn't work, just brute force it! Takes ~35 secs
	for _, row := range floorMap.Layout {
		for _, cell := range row {
			if cell.Free && (cell.Pos.X != startingPos.X || cell.Pos.Y != startingPos.Y) {
				clonedFloorMap := floorMap.Clone()
				clonedFloorMap.GetCell(cell.Pos).Free = false

				wg.Add(1)
				go clonedFloorMap.CheckIfExits(*startingPos, *startingDir, c, &wg)
			}
		}
	}

	go sol.WaitAndClose(c, &wg)
	for val := range c {
		resp += val
	}

	return strconv.Itoa(resp)
}
