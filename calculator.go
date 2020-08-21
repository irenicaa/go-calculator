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
	tokenizer  tokenizer.Tokenizer
	translator translator.Translator
	evaluator  evaluator.Evaluator
}

// Calculate ...
func (calculator *Calculator) Calculate(
	code string,
	variables map[string]float64,
	functionsNames map[string]struct{},
	functions map[string]models.Function,
) error {
	tokens, err := calculator.tokenizer.Tokenize(code)
	if err != nil {
		return fmt.Errorf("unable to tokenize the code: %s", err)
	}

	commands, err := calculator.translator.Translate(tokens, functionsNames)
	if err != nil {
		return fmt.Errorf("unable to translate the tokens: %s", err)
	}

	err = calculator.evaluator.Evaluate(commands, variables, functions)
	if err != nil {
		return fmt.Errorf("unable to evaluate the commands: %s", err)
	}

	return nil
}
