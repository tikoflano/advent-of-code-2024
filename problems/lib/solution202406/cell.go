package solution202406

import (
	"fmt"
	"slices"
)

type Cell struct{
	Pos *Vector
	Visited []*Vector
	Free bool
}

func (cell *Cell) IsVisited() bool {
	return len(cell.Visited) > 0
}

func (cell *Cell) IsVisitedFromDir(dir *Vector) bool {
	for _, visitedDir := range cell.Visited {
		if *visitedDir == *dir {
			return true
		}
	}

	return false
}

func (cell *Cell) AddVisit(dir *Vector) {
	if !slices.Contains(cell.Visited, dir){
		cell.Visited = append(cell.Visited, dir)
	}
}

func (cell *Cell) String() string {
	var notFree, notVisited string
	if !cell.Free {
		notFree = "not "
	}

	if !cell.IsVisited() {
		notVisited = "not "
	}

	return fmt.Sprintf("Cell at (%d,%d) is %sFree and %sVisited", cell.Pos.X, cell.Pos.Y, notFree, notVisited)
}