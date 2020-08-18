package translator

import (
	"errors"
	"fmt"

	"github.com/irenicaa/go-calculator/models"
	"github.com/irenicaa/go-calculator/models/containers"
)

var errStop = errors.New("stop")
var errStopAndRestore = errors.New("stop and restore")

type stackChecker func(tokenOnStack models.Token, ok bool) error

// Translator ...
type Translator struct {
	commands []models.Command
	stack    containers.TokenStack
}

// Translate ...
func (translator *Translator) Translate(
	tokens []models.Token,
	functions map[string]struct{},
) ([]models.Command, error) {
	for tokenIndex, token := range tokens {
		switch {
		case token.Kind == models.NumberToken:
			translator.addCommand(models.PushNumberCommand, token)
		case token.Kind == models.IdentifierToken:
			_, ok := functions[token.Value]
			if ok {
				translator.stack.Push(token)
				continue
			}

			translator.addCommand(models.PushVariableCommand, token)
		case token.Kind.IsOperator():
			// in this case, all errors will be processed inside the method
			translator.unwindStack(func(tokenOnStack models.Token, ok bool) error {
				if !ok {
					return errStop
				}
				if !tokenOnStack.Kind.IsOperator() {
					return errStopAndRestore
				}
				if tokenOnStack.Kind.Precedence() < token.Kind.Precedence() {
					return errStopAndRestore
				}

				return nil
			})

			translator.stack.Push(token)
		case token.Kind == models.LeftParenthesisToken:
			translator.stack.Push(token)
		case token.Kind == models.RightParenthesisToken:
			err := translator.unwindStack(
				func(tokenOnStack models.Token, ok bool) error {
					if !ok {
						return fmt.Errorf(
							"missed pair for token %+v with number #%d",
							token,
							tokenIndex,
						)
					}
					if tokenOnStack.Kind == models.LeftParenthesisToken {
						return errStop
					}

					return nil
				},
			)
			if err != nil {
				return nil, err
			}
		case token.Kind == models.CommaToken:
			err := translator.unwindStack(
				func(tokenOnStack models.Token, ok bool) error {
					if !ok {
						return fmt.Errorf(
							"missed pair for token %+v with number #%d",
							token,
							tokenIndex,
						)
					}
					if tokenOnStack.Kind == models.LeftParenthesisToken {
						return errStopAndRestore
					}
					if !tokenOnStack.Kind.IsOperator() {
						return errStopAndRestore
					}

					return nil
				},
			)
			if err != nil {
				return nil, err
			}
		}
	}

	commands := translator.commands
	translator.commands = nil

	return commands, nil
}

// Finalize ...
func (translator *Translator) Finalize() ([]models.Command, error) {
	err := translator.unwindStack(func(tokenOnStack models.Token, ok bool) error {
		if !ok {
			return errStop
		}
		if tokenOnStack.Kind.IsParenthesis() {
			return fmt.Errorf("missed pair for token %+v", tokenOnStack)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return translator.commands, nil
}

func (translator *Translator) addCommand(
	kind models.CommandKind,
	token models.Token,
) {
	command := models.Command{Kind: kind, Operand: token.Value}
	translator.commands = append(translator.commands, command)
}

func (translator *Translator) unwindStack(checker stackChecker) error {
	for {
		tokenOnStack, ok := translator.stack.Pop()

		err := checker(tokenOnStack, ok)
		switch err {
		case nil:
		case errStopAndRestore:
			translator.stack.Push(tokenOnStack)

			fallthrough
		case errStop:
			return nil
		default:
			return err
		}

		translator.addCommand(models.CallFunctionCommand, tokenOnStack)
	}
}
