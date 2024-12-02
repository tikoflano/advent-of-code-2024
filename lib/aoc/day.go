package aoc

import (
	"fmt"
	"io"
	"net/http"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/filemanager"
	"tikoflano/aoc/lib/utils"
)

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

func (day *Day) getInputFilePath() string {
	return fmt.Sprintf("problems/input/%d_%02d.txt", constants.Year, day.Number)
}

func (day *Day) DownloadInput(client *http.Client) {
	inputFilePath := day.getInputFilePath()
	exists, err := filemanager.FileExists(inputFilePath)
	utils.CheckError(err, "Error checking input file")

	if !exists {
		file := filemanager.CreateFile(inputFilePath)

		defer file.Close()

		inputURL := fmt.Sprintf("%s/input", day.Url)
		resp, err := client.Get(inputURL)
		utils.CheckError(err, fmt.Sprintf("Failed to GET input from %s", inputURL))

		defer resp.Body.Close()

		_, err = io.Copy(file, resp.Body)
		utils.CheckError(err, fmt.Sprintf("Failed to copy input contents into the input file %s", inputFilePath))
	}
}

func (day *Day) GetInput() string {
	inputFile := day.getInputFilePath()

	return filemanager.ReadFile(inputFile)
}
