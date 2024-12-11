package solution202410

import (
	"strconv"
)

type Step struct {
	Pos   Vector
	Next  []*Step
	Level int
}

func (step *Step) FindTrails(input []string) {
	for _, dir := range Directions {
		newPos := step.Pos.Add(dir)

		if !newPos.IsInbounds(input) {
			continue
		}

		curr, _ := strconv.Atoi(string(input[newPos.Y][newPos.X]))

		if curr == step.Level+1 {
			newStep := &Step{
				Pos:   Vector{X: newPos.X, Y: newPos.Y},
				Level: step.Level + 1,
			}
			step.Next = append(step.Next, newStep)

			newStep.FindTrails(input)
		}
	}
}

func (step *Step) GetMeasurements() (int, int) {
	if step.Level != 0 {
		return 0, 0
	}

	queue := []*Step{step}
	trailends := StepsSet{}
	ret := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Level == 9 {
			ret++
			trailends = trailends.addStep(curr)
		}

		if curr.Next != nil {
			queue = append(queue, curr.Next...)
		}
	}

	return len(trailends), ret
}
