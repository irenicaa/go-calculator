package calculator

import (
	"fmt"

	"github.com/irenicaa/go-calculator/evaluator"
	"github.com/irenicaa/go-calculator/models"
	"github.com/irenicaa/go-calculator/tokenizer"
	"github.com/irenicaa/go-calculator/translator"
)

// Calculator ...
type Calculator struct {
	functions      map[string]models.Function
	functionsNames map[string]struct{}

	tokenizer  tokenizer.Tokenizer
	translator translator.Translator
	evaluator  evaluator.Evaluator
}

// NewCalculator ...
func NewCalculator(functions map[string]models.Function) *Calculator {
	calculator := Calculator{
		functions:      functions,
		functionsNames: map[string]struct{}{},
	}

	for name := range functions {
		calculator.functionsNames[name] = struct{}{}
	}

	return &calculator
}

// Calculate ...
func (calculator *Calculator) Calculate(
	code string,
	variables map[string]float64,
) error {
	tokens, err := calculator.tokenizer.Tokenize(code)
	if err != nil {
		return fmt.Errorf("unable to tokenize the code: %s", err)
	}

	commands, err := calculator.translator.Translate(
		tokens,
		calculator.functionsNames,
	)
	if err != nil {
		return fmt.Errorf("unable to translate the tokens: %s", err)
	}

	err = calculator.evaluator.Evaluate(commands, variables, calculator.functions)
	if err != nil {
		return fmt.Errorf("unable to evaluate the commands: %s", err)
	}

	return nil
}
