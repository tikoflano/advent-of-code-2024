package solution202403

type Operation struct {
	command  *Command
	operands []string
}

func (operation *Operation) addOperand(operand string) {
	if operand == "" {
		return
	}

	operation.operands = append(operation.operands, operand)
}
