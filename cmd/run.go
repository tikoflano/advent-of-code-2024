package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "tikoflano/aoc/lib/aoc"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/utils"
	"tikoflano/aoc/problems"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the current problem's code. You can pass an argument with the filename to run.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		aoc := advent.NewAdventOfCode(constants.Year)
		problemToRun := aoc.NextDay().NextProblem()

		if len(args) > 0 && args[0] != "" {
			filename := strings.Split(args[0], ".")[0]
			parts := strings.Split(filename, "_")

			if len(parts) == 3 {
				log.Printf("Running file: %s", args[0])
				dayNumber, err := strconv.Atoi(parts[1])
				utils.CheckError(err, "Failed to parse the day value")

				problemNumber, err := strconv.Atoi(parts[2])
				utils.CheckError(err, "Failed to parse the problem value")

				problemToRun = &aoc.Days[dayNumber-1].Problems[problemNumber-1]
			} else {
				log.Printf("Invalid file to run: '%s', defaulting to next problem", args[0])
			}
		}

		input, exists := problemToRun.Day.GetInput()

		if !exists {
			log.Fatal("Input file not found, run the continue command to get it")
		}

		inputLines := strings.Split(input, "\n")

		// Remove the trailing empty line
		inputLines = inputLines[:len(inputLines)-1]

		output, err := problems.Run(problemToRun.GetSolutionKey(), inputLines)
		utils.CheckError(err, "Failed to run solution")

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
