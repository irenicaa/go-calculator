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
	variables      models.VariableGroup
	functions      models.FunctionGroup
	functionsNames models.FunctionNameGroup

	tokenizer  tokenizer.Tokenizer
	translator translator.Translator
	evaluator  evaluator.Evaluator
}

// NewCalculator ...
func NewCalculator(
	variables models.VariableGroup,
	functions models.FunctionGroup,
) *Calculator {
	return &Calculator{
		variables:      variables,
		functions:      functions,
		functionsNames: functions.Names(),
	}
}

// Calculate ...
func (calculator *Calculator) Calculate(code string) error {
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

	err = calculator.evaluator.Evaluate(
		commands,
		calculator.variables,
		calculator.functions,
	)
	if err != nil {
		return fmt.Errorf("unable to evaluate the commands: %s", err)
	}

	return nil
}

// Finalize ...
func (calculator *Calculator) Finalize() (float64, error) {
	tokens, err := calculator.tokenizer.Finalize()
	if err != nil {
		return 0, fmt.Errorf("unable to finalize the tokenizer: %s", err)
	}

	// data that came from that Finalize() call
	// can't lead to an error inside the Translate() call
	commands, _ := calculator.translator.Translate(
		tokens,
		calculator.functionsNames,
	)
	additionalCommands, err := calculator.translator.Finalize()
	if err != nil {
		return 0, fmt.Errorf("unable to finalize the translator: %s", err)
	}
	commands = append(commands, additionalCommands...)

	err = calculator.evaluator.Evaluate(
		commands,
		calculator.variables,
		calculator.functions,
	)
	if err != nil {
		return 0, fmt.Errorf("unable to evaluate the commands: %s", err)
	}

	number, err := calculator.evaluator.Finalize()
	if err != nil {
		return 0, fmt.Errorf("unable to finalize the evaluator: %s", err)
	}

	return number, nil
}
