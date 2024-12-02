package cmd

import (
	"fmt"
	"log"
	"strings"
	advent "tikoflano/aoc/lib/aoc"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/problems"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the current problem's code",
	Long:  `Execute the current problem's code`,
	Run: func(cmd *cobra.Command, args []string) {
		aoc := advent.NewAdventOfCode(constants.Year)
		nextProblem := aoc.NextDay().NextProblem()

		input, exists := nextProblem.Day.GetInput()

		if !exists {
			log.Fatal("Input file not found, run the continue command to get it")
		}

		inputLines := strings.Split(input, "\n")
		output := problems.Run(inputLines)

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
