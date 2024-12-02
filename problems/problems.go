package problems

import "errors"

type solution func(input []string) string

var solutions map[string]solution = make(map[string](solution))

func Run(problem string, input []string) (string, error) {
	fn, ok := solutions[problem]

	if !ok {
		return "", errors.New("solution not found")
	}

	return fn(input), nil
}
