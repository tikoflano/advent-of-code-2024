package solution202403

type executor struct {
	Enabled bool
}

func NewExecutor() executor {
	return executor{Enabled: true}
}

func (executor *executor) RunProgram(operations []*Operation) int {
	resp := 0

	for _, operation := range operations {
		result := operation.command.Run(executor, operation.operands)

		if value, ok := result.(int); ok {
			resp += value
		}
	}

	return resp
}
