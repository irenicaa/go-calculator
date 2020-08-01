package calculator

import (
	"errors"
	"fmt"
)

var errStop = errors.New("stop")
var errStopAndRestore = errors.New("stop and restore")

// CommandKind ...
type CommandKind int

// ...
const (
	PushNumberCommand CommandKind = iota
	PushVariableCommand
	CallFunctionCommand
)

// Command ...
type Command struct {
	Kind    CommandKind
	Operand string
}

// Translator ...
type Translator struct {
	commands []Command
	stack    TokenStack
}

// Translate ...
func (translator *Translator) Translate(
	tokens []Token,
	functions map[string]struct{},
) error {
	for tokenIndex, token := range tokens {
		switch {
		case token.Kind == NumberToken:
			translator.addCommand(PushNumberCommand, token)
		case token.Kind == IdentifierToken:
			_, ok := functions[token.Value]
			if ok {
				translator.stack.Push(token)
				continue
			}

			translator.addCommand(PushVariableCommand, token)
		case token.Kind.IsOperator():
			for {
				tokenOnStack, ok := translator.stack.Pop()
				if !ok {
					break
				}
				if !tokenOnStack.Kind.IsOperator() {
					translator.stack.Push(tokenOnStack)
					break
				}
				if tokenOnStack.Kind.Precedence() < token.Kind.Precedence() {
					translator.stack.Push(tokenOnStack)
					break
				}

				translator.addCommand(CallFunctionCommand, tokenOnStack)
			}

			translator.stack.Push(token)
		case token.Kind == LeftParenthesisToken:
			translator.stack.Push(token)
		case token.Kind == RightParenthesisToken:
			err := translator.unwindStack(func(tokenOnStack Token, ok bool) error {
				if !ok {
					return fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					return errStop
				}

				return nil
			})
			if err != nil {
				return err
			}
		case token.Kind == CommaToken:
			err := translator.unwindStack(func(tokenOnStack Token, ok bool) error {
				if !ok {
					return fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					return errStopAndRestore
				}
				if !tokenOnStack.Kind.IsOperator() {
					return errStopAndRestore
				}

				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Finalize ...
func (translator *Translator) Finalize() ([]Command, error) {
	err := translator.unwindStack(func(tokenOnStack Token, ok bool) error {
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

func (translator *Translator) addCommand(kind CommandKind, token Token) {
	command := Command{kind, token.Value}
	translator.commands = append(translator.commands, command)
}

func (translator *Translator) unwindStack(
	checker func(tokenOnStack Token, ok bool) error,
) error {
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

		translator.addCommand(CallFunctionCommand, tokenOnStack)
	}
}
