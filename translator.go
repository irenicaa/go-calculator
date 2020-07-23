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
			command := Command{PushNumberCommand, token.Value}
			translator.commands = append(translator.commands, command)
		case token.Kind == IdentifierToken:
			_, ok := functions[token.Value]
			if ok {
				translator.stack.Push(token)
				continue
			}

			command := Command{PushVariableCommand, token.Value}
			translator.commands = append(translator.commands, command)
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

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				translator.commands = append(translator.commands, command)
			}

			translator.stack.Push(token)
		case token.Kind == LeftParenthesisToken:
			translator.stack.Push(token)
		case token.Kind == RightParenthesisToken:
			for {
				tokenOnStack, ok := translator.stack.Pop()
				if !ok {
					return fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				translator.commands = append(translator.commands, command)
			}
		case token.Kind == CommaToken:
			for {
				tokenOnStack, ok := translator.stack.Pop()
				if !ok {
					return fmt.Errorf(
						"missed pair for token %+v with number #%d",
						token,
						tokenIndex,
					)
				}
				if tokenOnStack.Kind == LeftParenthesisToken {
					translator.stack.Push(tokenOnStack)
					break
				}
				if !tokenOnStack.Kind.IsOperator() {
					translator.stack.Push(tokenOnStack)
					break
				}

				command := Command{CallFunctionCommand, tokenOnStack.Value}
				translator.commands = append(translator.commands, command)
			}
		}
	}

	return nil
}

// Finalize ...
func (translator *Translator) Finalize() ([]Command, error) {
	for {
		tokenOnStack, ok := translator.stack.Pop()
		if !ok {
			break
		}
		if tokenOnStack.Kind.IsParenthesis() {
			return nil, fmt.Errorf("missed pair for token %+v", tokenOnStack)
		}

		command := Command{CallFunctionCommand, tokenOnStack.Value}
		translator.commands = append(translator.commands, command)
	}

	return translator.commands, nil
}
