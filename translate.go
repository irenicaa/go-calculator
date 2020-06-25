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
func Translate(tokens []Token, functions map[string]struct{}) ([]Command, error) {
	commands := []Command{}
	stack := TokenStack{}
	for tokenIndex, token := range tokens {
		switch {
		case token.Kind == NumberToken:
			command := Command{PushNumberCommand, token.Value}
			commands = append(commands, command)
		case token.Kind == IdentifierToken:
			_, ok := functions[token.Value]
			if ok {
				stack.Push(token)
				continue
			}

			command := Command{PushVariableCommand, token.Value}
			commands = append(commands, command)
		case token.Kind.IsOperator():
			for {
				tokenOnStack, ok := stack.Pop()
				if !ok {
					break
				}
				if !tokenOnStack.Kind.IsOperator() {
					stack.Push(tokenOnStack)
					break
				}
				if tokenOnStack.Kind.Precedence() < token.Kind.Precedence() {
					stack.Push(tokenOnStack)
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				commands = append(commands, command)
			}

			stack.Push(token)
		case token.Kind == LeftParenthesisToken:
			stack.Push(token)
		case token.Kind == RightParenthesisToken:
			for {
				tokenOnStack, ok := stack.Pop()
				if !ok {
					return nil, fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					break
				}
				if !tokenOnStack.Kind.IsOperator() && tokenOnStack.Kind != IdentifierToken {
					stack.Push(tokenOnStack)
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				commands = append(commands, command)
			}
		case token.Kind == CommaToken:
			for {
				tokenOnStack, ok := stack.Pop()
				if !ok {
					return nil, fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					stack.Push(tokenOnStack)
					break
				}
				if !tokenOnStack.Kind.IsOperator() {
					stack.Push(tokenOnStack)
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				commands = append(commands, command)
			}
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
		if tokenOnStack.Kind.IsParenthesis() {
			return nil, fmt.Errorf("missed pair for token %+v", tokenOnStack)
		}

		command := Command{CallFunctionCommand, tokenOnStack.Value}
		commands = append(commands, command)
	}

	return commands, nil
}
