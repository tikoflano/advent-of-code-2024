package solution202406

type Vector struct {
	X, Y int
}

type DirectionSet struct {
	N,E,S,W *Vector
}

var Directions = &DirectionSet{
	N:  &Vector{X: 0, Y: -1},
	E:  &Vector{X: 1, Y: 0},
	S:  &Vector{X: 0, Y: 1},
	W:  &Vector{X: -1, Y: 0},
}

func (vector *Vector) Rotate90() *Vector {
	return &Vector{
		X: vector.Y * -1,
		Y: vector.X,
	}
}

func (vector *Vector) Add(v *Vector) *Vector {
	return &Vector{
		X: vector.X + v.X,
		Y: vector.Y + v.Y,
	}
}