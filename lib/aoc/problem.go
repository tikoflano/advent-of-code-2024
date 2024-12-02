package aoc

import (
	"fmt"
	"io"
	"net/http"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/filemanager"
	"tikoflano/aoc/lib/utils"
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

func (problem *Problem) getInputFilePath() string {
	return fmt.Sprintf("problems/input/%d_%02d_%d.txt", constants.Year, problem.Day.Number, problem.Number)
}

func (problem *Problem) MakeProblemFile() string {
	problemFilePath := problem.getProblemFilePath()
	exists, err := filemanager.FileExists(problemFilePath)
	utils.CheckError(err, "Error checking problem file")

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

func (problem *Problem) DownloadInput(client *http.Client) {
	inputFilePath := problem.getInputFilePath()
	exists, err := filemanager.FileExists(inputFilePath)
	utils.CheckError(err, "Error checking input file")

	if !exists {
		file := filemanager.CreateFile(inputFilePath)

		defer file.Close()

		inputURL := fmt.Sprintf("%s/input", problem.Day.Url)
		resp, err := client.Get(inputURL)
		utils.CheckError(err, fmt.Sprintf("Failed to GET input from %s", inputURL))

		defer resp.Body.Close()

		_, err = io.Copy(file, resp.Body)
		utils.CheckError(err, fmt.Sprintf("Failed to copy input contents into the input file %s", inputFilePath))
	}
}

func (problem *Problem) GetInput() string {
	inputFile := fmt.Sprintf("problems/input/%d_%02d_%d.txt", problem.Day.Aoc.Year, problem.Day.Number, problem.Number)

	return filemanager.ReadFile(inputFile)
}
