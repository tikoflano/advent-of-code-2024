package aoc

import "fmt"

type Day struct {
	Number   int
	Url      string
	Problems [2]Problem
	Aoc      *AdventOfCode
}

func (day Day) String() string {
	var solved string

	if day.Problems[1].Solved {
		solved = "⭐⭐"
	} else if day.Problems[0].Solved {
		solved = "⭐"
	} else {
		solved = "⛔"
	}

	return fmt.Sprintf("Day %d: %s", day.Number, solved)
}

func (day *Day) IsSolved() bool {
	return day.Problems[0].Solved && day.Problems[1].Solved
}

func (day *Day) NextProblem() *Problem {
	for _, problem := range day.Problems {
		if !problem.Solved {
			return &problem
		}
	}

	return nil
}
