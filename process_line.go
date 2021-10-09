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

// ProcessLine ...
func ProcessLine(
	line string,
	variables map[string]float64,
	functions models.FunctionGroup,
) (float64, error) {
	line = tokenizer.RemoveComment(line)

	variable, line := tokenizer.ExtractVariable(line)
	if variable == "" && strings.TrimSpace(line) == "" {
		return 0, ErrNoCode
	}

	calculator := NewCalculator(variables, functions)
	if err := calculator.Calculate(line); err != nil {
		return 0, fmt.Errorf("unable to calculate the line: %s", err)
	}

	number, err := calculator.Finalize()
	if err != nil {
		return 0, fmt.Errorf("unable to finalize the calculator: %s", err)
	}

	if variable != "" {
		variables[variable] = number
	}

	return number, nil
}
