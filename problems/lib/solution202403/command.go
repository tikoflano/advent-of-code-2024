package solution202403

import (
	"strconv"
)

type Command struct {
	Keyword string
	Run     func(executor *executor, operands []string) interface{}
}

var Commands = []*Command{
	{
		Keyword: "mul",
		Run: func(executor *executor, operands []string) interface{} {
			if !executor.Enabled {
				return nil
			}

			if len(operands) != 2 {
				return nil
			}

			num1, err := strconv.Atoi(operands[0])
			if err != nil {
				return nil
			}
			num2, err := strconv.Atoi(operands[1])
			if err != nil {
				return nil
			}

			return num1 * num2
		},
	},
	{
		Keyword: "do",
		Run: func(executor *executor, operands []string) interface{} {
			if len(operands) != 0 {
				return nil
			}

			executor.Enabled = true
			return nil
		},
	},
	{
		Keyword: "don't",
		Run: func(executor *executor, operands []string) interface{} {
			if len(operands) != 0 {
				return nil
			}

			executor.Enabled = false
			return nil
		},
	},
}
