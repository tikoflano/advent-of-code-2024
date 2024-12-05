package solution202404

type Vector struct {
	X, Y int
}

var Directions = map[string]*Vector{
	"N":  {X: 0, Y: -1},
	"NE": {X: 1, Y: -1},
	"E":  {X: 1, Y: 0},
	"SE": {X: 1, Y: 1},
	"S":  {X: 0, Y: 1},
	"SW": {X: -1, Y: 1},
	"W":  {X: -1, Y: 0},
	"NW": {X: -1, Y: -1},
}

func (vector *Vector) Scale(n int) *Vector {
	return &Vector{
		X: vector.X * n,
		Y: vector.Y * n,
	}
}

func (vector *Vector) Add(v *Vector) *Vector {
	return &Vector{
		X: vector.X + v.X,
		Y: vector.Y + v.Y,
	}
}
