package solution202406

import (
	"errors"
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

func (floorMap *FloorMap) CheckVisitedInPath(pos *Vector, dir *Vector) bool {
	for {
		nextCell, err := floorMap.Look(pos, dir)
		if err != nil {
			return false
		}

		if nextCell.IsVisitedFromDir(dir) {
			return true
		}

		pos = nextCell.Pos
	}
}
