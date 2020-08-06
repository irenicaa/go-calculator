package calculator

import "fmt"

// Evaluate ...
func Evaluate(commands []Command) (float64, error) {
	stack := NumberStack{}
	for _, command := range commands {
		switch command.Kind {
		case PushNumberCommand:
		case PushVariableCommand:
		case CallFunctionCommand:
		}
	}

	number, ok := stack.Pop()
	if !ok {
		return 0, fmt.Errorf("number stack is empty")
	}

	return number, nil
}
