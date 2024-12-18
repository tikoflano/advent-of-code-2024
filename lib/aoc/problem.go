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

	return fmt.Sprintf("Day %02d, Problem %d: %s", problem.Day.Number, problem.Number, solved)
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
		Day         string
		Problem     int
		SolutionKey string
	}{
		problem.GetURL(),
		problem.Day.Aoc.Year,
		fmt.Sprintf("%02d", problem.Day.Number),
		problem.Number,
		problem.GetSolutionKey(),
	}

	filemanager.CreateFileFromTemplate(problemFilePath, constants.ProblemTemplateFile, tplData)

	return problemFilePath
}

// Using variadic args only to allow empty calls, only 1 alt value is exptected and used
func (problem *Problem) GetSolutionKey(alt ...string) string {
	suffix := ""
	if len(alt) > 1 {
		panic("invalid call to GetSolutionKey, only 0 or 1 arguments are allowed")
	}

	if len(alt) == 1 && alt[0] != "" {
		suffix = "_alternative_" + alt[0]
	}
	return fmt.Sprintf("%d_%02d_%d%s", problem.Day.Aoc.Year, problem.Day.Number, problem.Number, suffix)
}

func (problem *Problem) GetURL() string {
	if problem.Number == 1 {
		return problem.Day.Url
	}

	return fmt.Sprintf("%s#part%d", problem.Day.Url, problem.Number)
}
