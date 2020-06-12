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

// ParseOperation ...
func ParseOperation(symbol rune) (TokenKind, error) {
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
	default:
		return 0, fmt.Errorf("unknown symbol %q", symbol)
	}
}
