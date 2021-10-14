package calculator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/irenicaa/go-calculator/models"
	"github.com/irenicaa/go-calculator/tokenizer"
)

// ErrNoCode ...
var ErrNoCode = errors.New("no code")

// Interpreter ...
type Interpreter struct {
	variables models.VariableGroup
	functions models.FunctionGroup
}

// NewInterpreter ...
func NewInterpreter(
	variables models.VariableGroup,
	functions models.FunctionGroup,
) Interpreter {
	return Interpreter{variables: variables.Copy(), functions: functions}
}

// Interpret ...
func (interpreter Interpreter) Interpret(input string) (float64, error) {
	input = tokenizer.RemoveComment(input)

	variable, input := tokenizer.ExtractVariable(input)
	if variable == "" && strings.TrimSpace(input) == "" {
		return 0, ErrNoCode
	}

	calculator := NewCalculator(interpreter.variables, interpreter.functions)
	if err := calculator.Calculate(input); err != nil {
		return 0, fmt.Errorf("unable to calculate the input: %s", err)
	}

	number, err := calculator.Finalize()
	if err != nil {
		return 0, fmt.Errorf("unable to finalize the calculator: %s", err)
	}

	if variable != "" {
		interpreter.variables[variable] = number
	}

	return number, nil
}
