package cmd

import (
	"os/exec"
	advent "tikoflano/aoc/lib/aoc"
	"tikoflano/aoc/lib/constants"

	"github.com/spf13/cobra"
)

// continueCmd represents the continue command
var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Continue with the next problem",
	Long:  `This will create the necessary files, get the input and set up the run command.`,
	Run: func(cmd *cobra.Command, args []string) {
		aoc := advent.NewAdventOfCode(constants.Year)
		nextProblem := aoc.NextDay().NextProblem()

		problemFilePath := nextProblem.MakeProblemFile()
		nextProblem.Day.DownloadInput(aoc.Client)

		cmnd := exec.Command(constants.VSCodeExecutable, problemFilePath)
		cmnd.Start()
	},
}

func init() {
	rootCmd.AddCommand(continueCmd)
}
