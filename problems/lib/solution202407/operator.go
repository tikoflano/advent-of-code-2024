package solution202407

import (
	"strconv"
)

type Operator struct {
	Symbol string
	Exec   func(operand1, operand2 int) int
}

var Operators = []*Operator{
	{
		Symbol: "+",
		Exec:   func(operand1, operand2 int) int { return operand1 + operand2 },
	},
	{
		Symbol: "*",
		Exec:   func(operand1, operand2 int) int { return operand1 * operand2 },
	},
	{
		Symbol: "||",
		Exec: func(operand1, operand2 int) int {
			strOperand1 := strconv.Itoa(operand1)
			strOperand2 := strconv.Itoa(operand2)

			resp, _ := strconv.Atoi(strOperand1 + strOperand2)

			return resp
		},
	},
}
