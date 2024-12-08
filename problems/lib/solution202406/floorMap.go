package solution202406

import (
	"errors"
	"sync"
)

type FloorMap struct {
	Layout        [][]*Cell
	Width, Height int
}

func NewFloorMap(input []string) (*FloorMap, *Vector) {
	var startPos *Vector
	floorMap := FloorMap{
		Layout: [][]*Cell{},
	}

	for y, line := range input {
		floorMap.Layout = append(floorMap.Layout, []*Cell{})

		for x, cellData := range line {
			pos := &Vector{X: x, Y: y}
			cell := &Cell{
				Pos:     pos,
				Visited: []*Vector{},
				Free:    cellData != '#',
			}

			floorMap.Layout[y] = append(floorMap.Layout[y], cell)

			if cellData == '^' {
				startPos = pos
			}
		}
	}

	floorMap.Height = len(floorMap.Layout)
	floorMap.Width = len(floorMap.Layout[0])

	return &floorMap, startPos
}

func (floorMap *FloorMap) Clone() FloorMap {
	clone := FloorMap{
		Width:  floorMap.Width,
		Height: floorMap.Height,
		Layout: [][]*Cell{},
	}

	for y, row := range floorMap.Layout {
		clone.Layout = append(clone.Layout, []*Cell{})
		for _, cell := range row {
			clonedCell := cell.Clone()
			clone.Layout[y] = append(clone.Layout[y], &clonedCell)
		}
	}

	return clone
}

func (floorMap *FloorMap) Look(pos *Vector, dir *Vector) (*Cell, error) {
	targetPos := pos.Add(dir)

	if targetPos.X < 0 {
		return nil, errors.New("exit W")
	} else if targetPos.X >= floorMap.Width {
		return nil, errors.New("exit E")
	} else if targetPos.Y < 0 {
		return nil, errors.New("exit N")
	} else if targetPos.Y >= floorMap.Height {
		return nil, errors.New("exit S")
	}

	targetCell := floorMap.GetCell(targetPos)

	return targetCell, nil

}

func (floorMap *FloorMap) GetCell(pos *Vector) *Cell {
	return floorMap.Layout[pos.Y][pos.X]
}

func (floorMap *FloorMap) CheckIfExits(startingPos Vector, startingDir Vector, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	pos := &startingPos
	dir := &startingDir

	for {
		if curCell := floorMap.GetCell(pos); curCell.IsVisitedFromDir(dir) {
			c <- 1
			return
		} else {
			curCell.AddVisit(dir)
		}

		targetCell, err := floorMap.Look(pos, dir)
		if err != nil {
			c <- 0
			return
		}

		if targetCell.Free {
			pos = targetCell.Pos
		} else {
			dir = dir.Rotate90()
		}
	}
}
