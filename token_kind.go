package calculator

import "fmt"

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
	CommaToken
)

// lack of the error is guaranteed by the calling function
func ParseTokenKind(symbol rune) (TokenKind, error) {
	switch symbol {
	case '+':
		return PlusToken, nil
	case '-':
		return MinusToken, nil
	case '*':
		return AsteriskToken, nil
	case '/':
		return SlashToken, nil
	case '%':
		return PercentToken, nil
	case '^':
		return ExponentiationToken, nil
	case '(':
		return LeftParenthesisToken, nil
	case ')':
		return RightParenthesisToken, nil
	case ',':
		return CommaToken, nil
	default:
		return 0, fmt.Errorf("unknown symbol %q", symbol)
	}
}

// IsParenthesis ...
func (kind TokenKind) IsParenthesis() bool {
	return kind == LeftParenthesisToken || kind == RightParenthesisToken
}

// IsOperator ...
func (kind TokenKind) IsOperator() bool {
	switch kind {
	case PlusToken, MinusToken,
		AsteriskToken, SlashToken, PercentToken,
		ExponentiationToken:
		return true
	default:
		return false
	}
}

// Precedence ...
func (kind TokenKind) Precedence() int {
	switch kind {
	case PlusToken, MinusToken:
		return 1
	case AsteriskToken, SlashToken, PercentToken:
		return 2
	case ExponentiationToken:
		return 3
	default:
		return 0
	}
}
