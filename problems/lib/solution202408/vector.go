package solution202408

type Vector struct {
	X, Y int
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return &Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector) Substract(v2 *Vector) *Vector {
	return &Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v *Vector) Invert() *Vector {
	return &Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v *Vector) Scale(factor int) *Vector {
	return &Vector{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

func (v *Vector) IsInbounds(mapWidth, mapHeight int) bool {
	return v.X >= 0 && v.Y >= 0 && v.X < mapWidth && v.Y < mapHeight
}
