package solution202410

type Vector struct {
	X, Y int
}

var Directions = map[string]*Vector{
	"N": {X: 0, Y: -1},
	"E": {X: 1, Y: 0},
	"S": {X: 0, Y: 1},
	"W": {X: -1, Y: 0},
}

func (vector *Vector) Add(v *Vector) *Vector {
	return &Vector{
		X: vector.X + v.X,
		Y: vector.Y + v.Y,
	}
}

func (vector *Vector) IsInbounds(input []string) bool {
	return vector.X >= 0 && vector.Y >= 0 && vector.X < len(input[0]) && vector.Y < len(input)
}
