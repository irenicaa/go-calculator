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
	buffer := ""
	for _, symbol := range code {
		switch {
		case unicode.IsDigit(symbol):
			if state == identifierTokenizerState {
				buffer += string(symbol)
			}
		case unicode.IsLetter(symbol):
			if state == defaultTokenizerState {
				state = identifierTokenizerState
			}
			if state == identifierTokenizerState {
				buffer += string(symbol)
			}
		case symbol == '+':
		case symbol == '-':
		case symbol == '*':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{AsteriskToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '/':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{SlashToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '%':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{PercentToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '^':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{ExponentiationToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '(':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{LeftParenthesisToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == ')':
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{RightParenthesisToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '.':
		case symbol == '_':
			if state == defaultTokenizerState {
				state = identifierTokenizerState
			}
			if state == identifierTokenizerState {
				buffer += string(symbol)
			}
		default:
			return nil, fmt.Errorf("unknown symbol %q", symbol)
		}
	}

	return tokens, nil
}
