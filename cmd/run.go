package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	advent "tikoflano/aoc/lib/aoc"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/utils"
	"tikoflano/aoc/problems"

	"github.com/PuerkitoBio/goquery"
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
		alt := ""

		if len(args) > 0 && args[0] != "" {
			filename := strings.Split(args[0], ".")[0]
			parts := strings.Split(filename, "_")

			if len(parts) == 3 || len(parts) == 5 {
				log.Printf("Running file: %s", args[0])
				dayNumber, err := strconv.Atoi(parts[1])
				utils.CheckError(err, "Failed to parse the day value")

				problemNumber, err := strconv.Atoi(parts[2])
				utils.CheckError(err, "Failed to parse the problem value")

				problemToRun = &aoc.Days[dayNumber-1].Problems[problemNumber-1]

				// Alternative solution file called
				if len(parts) == 5 {
					alt = parts[4]
				}
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
		if inputLines[len(inputLines)-1] == "" {
			inputLines = inputLines[:len(inputLines)-1]
		}

		output, err := problems.Run(problemToRun.GetSolutionKey(alt), inputLines)
		utils.CheckError(err, "Failed to run solution")

		fmt.Printf("Result: %s\n", output)

		// Submit answer
		if submitAnswer, _ := cmd.Flags().GetBool("submit"); submitAnswer {
			data := url.Values{}
			data.Set("level", strconv.Itoa(problemToRun.Number))
			data.Set("answer", output)

			res, err := aoc.Client.Post(
				fmt.Sprintf("%s/%d/day/%d/answer", constants.BaseURL, aoc.Year, problemToRun.Day.Number),
				"application/x-www-form-urlencoded",
				strings.NewReader(data.Encode()),
			)
			utils.CheckError(err, "Error while submitting the answer")

			defer res.Body.Close()

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			utils.CheckError(err, "Failed to create HTML document")

			submissionResponse := doc.Find("article").First().Text()
			submissionResponse = strings.ReplaceAll(submissionResponse, "  ", "\n")

			fmt.Println(submissionResponse)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("submit", "s", false, "submit answer")
}
