package calculator

import (
	"fmt"
	"unicode"
)

// TokenKind ...
type TokenKind int

// ...
const (
	PlusToken TokenKind = iota
	MinusToken
	AsteriskToken
	SlashToken
	PercentToken
	ExponentiationToken
	NumberToken
	IdentifierToken
	LeftParenthesisToken
	RightParenthesisToken
)

type tokenizerState int

const (
	defaultTokenizerState tokenizerState = iota
	numberTokenizerState
	identifierTokenizerState
)

// Token ...
type Token struct {
	Kind  TokenKind
	Value string
}

// Tokenize ...
func Tokenize(code string) ([]Token, error) {
	tokens := []Token{}
	state := defaultTokenizerState
	for _, symbol := range code {
		switch {
		case unicode.IsDigit(symbol):
		case unicode.IsLetter(symbol):
		case symbol == '+':
		case symbol == '-':
		case symbol == '*':
		case symbol == '/':
		case symbol == '%':
		case symbol == '^':
		case symbol == '(':
		case symbol == ')':
		case symbol == '.':
		case symbol == '_':
		default:
			return nil, fmt.Errorf("unknown symbol %q", symbol)
		}
	}

	return tokens, nil
}
