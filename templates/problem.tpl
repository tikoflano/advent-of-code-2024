package problems

import (
	"strconv"
	sol "tikoflano/aoc/problems/lib/solution{{.Year}}{{.Day}}"
)

func init() {
	solutions["{{.SolutionKey}}"] = year{{.Year}}Day{{.Day}}Problem{{.Problem}}
}

// {{.URL}}
func year{{.Year}}Day{{.Day}}Problem{{.Problem}}(input []string) string {
	var resp

	for _, line := range input {

	}

	return strconv.Itoa(resp)
}