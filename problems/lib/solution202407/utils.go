package solution202407

func ValidateEquation(operands []int, target int) bool {
	if len(operands) == 1 {
		return operands[0] == target
	}

	for _, operator := range Operators {
		newOperand := operator.Exec(operands[0], operands[1])
		newOperands := append([]int{newOperand}, operands[2:]...)

		if ValidateEquation(newOperands, target) {
			return true
		}
	}

	return false
}
