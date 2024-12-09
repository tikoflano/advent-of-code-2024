package solution202408

type Antinodes []*Vector

func GetAntinodesWithNoHarmonics(posA, posB *Vector) Antinodes {
	diff := posA.Substract(posB)
	invertedDiff := diff.Invert()

	antinodeA := posA.Add(diff)
	antinodeB := posB.Add(invertedDiff)

	return Antinodes{antinodeA, antinodeB}
}

func GetAntinodesWithHarmonics(posA, posB *Vector, mapWidth, mapHeight int) Antinodes {
	var ret Antinodes

	diff := posA.Substract(posB)
	invertedDiff := diff.Invert()

	factor := 1
	for {
		newAntinode := posA.Add(diff.Scale(factor))

		if !newAntinode.IsInbounds(mapWidth, mapHeight) {
			break
		}

		ret = append(ret, newAntinode)
		factor++
	}

	factor = 1
	for {
		newAntinode := posB.Add(invertedDiff.Scale(factor))

		if !newAntinode.IsInbounds(mapWidth, mapHeight) {
			break
		}

		ret = append(ret, newAntinode)
		factor++
	}

	return ret
}

func (antinodes Antinodes) FilterAntinodes(mapWidth, mapHeight int) Antinodes {
	var ret Antinodes

	for _, antinode := range antinodes {
		if antinode.IsInbounds(mapWidth, mapHeight) {
			ret = append(ret, antinode)
		}
	}

	return ret
}

func (antinodes Antinodes) AddAntinode(newAntinode *Vector) Antinodes {
	for _, antinode := range antinodes {
		if antinode.X == newAntinode.X && antinode.Y == newAntinode.Y {
			return antinodes
		}
	}

	return append(antinodes, newAntinode)
}
