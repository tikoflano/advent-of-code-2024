package aoc

import (
	"fmt"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/filemanager"
)

type Problem struct {
	Number int
	Solved bool
	Day    *Day
}

func (problem Problem) String() string {
	return fmt.Sprint(problem.Day)
}

func (problem *Problem) getProblemFilePath() string {
	return fmt.Sprintf("problems/%d_%02d_%d.go", constants.Year, problem.Day.Number, problem.Number)
}

func (problem *Problem) MakeProblemFile() string {
	problemFilePath := problem.getProblemFilePath()
	exists := filemanager.FileExists(problemFilePath)

	if exists {
		return problemFilePath
	}

	tplData := struct {
		URL     string
		Year    int
		Day     int
		Problem int
	}{
		problem.Day.Url,
		problem.Day.Aoc.Year,
		problem.Day.Number,
		problem.Number,
	}

	filemanager.CreateFileFromTemplate(problemFilePath, constants.ProblemTemplateFile, tplData)

	return problemFilePath
}
