package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

// Function ...
type Function struct {
	Arity   int
	Handler func(arguments []float64) float64
}

// Evaluator ...
type Evaluator struct {
	stack NumberStack
}

// Evaluate ...
func (evaluator *Evaluator) Evaluate(
	commands []Command,
	variables map[string]float64,
	functions map[string]Function,
) error {
	for commandIndex, command := range commands {
		switch command.Kind {
		case PushNumberCommand:
			number, err := strconv.ParseFloat(command.Operand, 64)
			if err != nil {
				return fmt.Errorf(
					"incorrect number for command %+v with number #%d: %s",
					command,
					commandIndex,
					err,
				)
			}

			evaluator.stack.Push(number)
		case PushVariableCommand:
			number, ok := variables[command.Operand]
			if !ok {
				return fmt.Errorf(
					"unknown variable in command %+v with number #%d",
					command,
					commandIndex,
				)
			}

			evaluator.stack.Push(number)
		case CallFunctionCommand:
			function, ok := functions[command.Operand]
			if !ok {
				return fmt.Errorf(
					"unknown function in command %+v with number #%d",
					command,
					commandIndex,
				)
			}

			arguments := []float64{}
			for argumentIndex := 0; argumentIndex < function.Arity; argumentIndex++ {
				number, ok := evaluator.stack.Pop()
				if !ok {
					return fmt.Errorf(
						"number stack is empty for argument #%d in command %+v with number #%d",
						argumentIndex,
						command,
						commandIndex,
					)
				}

				arguments = append(arguments, number)
			}

			for i := 0; i < function.Arity/2; i++ {
				arguments[function.Arity-i-1], arguments[i] =
					arguments[i], arguments[function.Arity-i-1]
			}

			number := function.Handler(arguments)
			evaluator.stack.Push(number)
		}
	}

	return nil
}

// Finalize ...
func (evaluator Evaluator) Finalize() (float64, error) {
	number, ok := evaluator.stack.Pop()
	if !ok {
		return 0, errors.New("number stack is empty")
	}

	return number, nil
}
