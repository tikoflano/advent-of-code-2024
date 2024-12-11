package solution202410

type StepsSet []*Step

func (stepsSet StepsSet) addStep(newStep *Step) []*Step {
	for _, step := range stepsSet {
		if step.Pos.X == newStep.Pos.X && step.Pos.Y == newStep.Pos.Y {
			return stepsSet
		}
	}

	return append(stepsSet, newStep)
}
