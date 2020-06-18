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
	stack := TokenStack{}
	for tokenIndex, token := range tokens {
		switch {
		case token.Kind == NumberToken:
			command := Command{PushNumberCommand, token.Value}
			commands = append(commands, command)
		case token.Kind == IdentifierToken:
			command := Command{PushVariableCommand, token.Value}
			commands = append(commands, command)
		case token.Kind.IsOperator():
			for {
				tokenOnStack, ok := stack.Pop()
				if !ok {
					break
				}
				if !tokenOnStack.Kind.IsOperator() {
					break
				}
				if tokenOnStack.Kind.Precedence() < token.Kind.Precedence() {
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				commands = append(commands, command)
			}

			stack.Push(token)
		default:
			return nil, fmt.Errorf(
				"unknown token %+v with number #%d",
				token,
				tokenIndex,
			)
		}
	}
	for {
		tokenOnStack, ok := stack.Pop()
		if !ok {
			break
		}

		command := Command{CallFunctionCommand, tokenOnStack.Value}
		commands = append(commands, command)
	}

	return commands, nil
}
