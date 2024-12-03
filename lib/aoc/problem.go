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
	solved := "⛔"

	if problem.Solved {
		solved = "⭐"
	}

	return fmt.Sprintf("Day %d, Problem %d: %s", problem.Day.Number, problem.Number, solved)
}

func (problem *Problem) getProblemFilePath() string {
	return fmt.Sprintf("problems/%s.go", problem.GetSolutionKey())
}

func (problem *Problem) MakeProblemFile() string {
	problemFilePath := problem.getProblemFilePath()
	exists := filemanager.FileExists(problemFilePath)

	if exists {
		return problemFilePath
	}

	tplData := struct {
		URL         string
		Year        int
		Day         int
		Problem     int
		SolutionKey string
	}{
		problem.GetURL(),
		problem.Day.Aoc.Year,
		problem.Day.Number,
		problem.Number,
		problem.GetSolutionKey(),
	}

	filemanager.CreateFileFromTemplate(problemFilePath, constants.ProblemTemplateFile, tplData)

	return problemFilePath
}

func (problem *Problem) GetSolutionKey() string {
	return fmt.Sprintf("%d_%02d_%d", problem.Day.Aoc.Year, problem.Day.Number, problem.Number)
}

func (problem *Problem) GetURL() string {
	if problem.Number == 1 {
		return problem.Day.Url
	}

	return fmt.Sprintf("%s#part%d", problem.Day.Url, problem.Number)
}
