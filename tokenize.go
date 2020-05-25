package calculator

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

// Token ...
type Token struct {
	Kind  TokenKind
	Value string
}

// Tokenize ...
func Tokenize(code string) []Token {
	tokens := []Token{}
	for i := 0; i < len(code); i++ {
		symbol := code[i]
	}

	return tokens
}
