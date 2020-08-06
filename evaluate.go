package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

// Evaluate ...
func Evaluate(commands []Command) (float64, error) {
	stack := NumberStack{}
	for commandIndex, command := range commands {
		switch command.Kind {
		case PushNumberCommand:
			number, err := strconv.ParseFloat(command.Operand, 64)
			if err != nil {
				return 0, fmt.Errorf(
					"incorrect number for command %+v with number #%d: %s",
					command,
					commandIndex,
					err,
				)
			}

			stack.Push(number)
		case PushVariableCommand:
		case CallFunctionCommand:
		}
	}

	number, ok := stack.Pop()
	if !ok {
		return 0, errors.New("number stack is empty")
	}

	return number, nil
}
