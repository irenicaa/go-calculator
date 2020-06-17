package calculator

import "fmt"

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

// Translate ...
func Translate(tokens []Token) ([]Command, error) {
	commands := []Command{}
	for tokenIndex, token := range tokens {
		switch token.Kind {
		default:
			return nil, fmt.Errorf(
				"unknown token %+v with number #%d",
				token,
				tokenIndex,
			)
		}
	}

	return commands, nil
}
