package evaluator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/irenicaa/go-calculator/models"
	"github.com/irenicaa/go-calculator/models/containers"
)

// Evaluator ...
type Evaluator struct {
	stack containers.NumberStack
}

// Evaluate ...
func (evaluator *Evaluator) Evaluate(
	commands []models.Command,
	variables map[string]float64,
	functions models.FunctionGroup,
) error {
	for commandIndex, command := range commands {
		switch command.Kind {
		case models.PushNumberCommand:
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
		case models.PushVariableCommand:
			number, ok := variables[command.Operand]
			if !ok {
				return fmt.Errorf(
					"unknown variable in command %+v with number #%d",
					command,
					commandIndex,
				)
			}

			evaluator.stack.Push(number)
		case models.CallFunctionCommand:
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

			reverseArguments(arguments)

			number := function.Handler(arguments)
			evaluator.stack.Push(number)
		}
	}

	return nil
}

func reverseArguments(arguments []float64) {
	arity := len(arguments)
	for i := 0; i < arity/2; i++ {
		arguments[arity-i-1], arguments[i] =
			arguments[i], arguments[arity-i-1]
	}
}

// Finalize ...
func (evaluator Evaluator) Finalize() (float64, error) {
	number, ok := evaluator.stack.Pop()
	if !ok {
		return 0, errors.New("number stack is empty")
	}

	return number, nil
}
