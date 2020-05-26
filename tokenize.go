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
	integerPartTokenizerState
	fractionalPartTokenizerState
	exponentTokenizerState
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
	for index, symbol := range code {
		switch {
		case unicode.IsDigit(symbol):
			if state == defaultTokenizerState {
				state = integerPartTokenizerState
			}
			buffer += string(symbol)
		case unicode.IsLetter(symbol) || symbol == '_':
			if state == defaultTokenizerState {
				state = identifierTokenizerState
			}
			if state == integerPartTokenizerState {
				state = identifierTokenizerState
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == fractionalPartTokenizerState {
				if symbol == 'e' || symbol == 'E' {
					state = exponentTokenizerState
					buffer += string(symbol)
					break
				}
				state = identifierTokenizerState
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				state = identifierTokenizerState
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			buffer += string(symbol)
		case symbol == '+':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					buffer += string(symbol)
					break
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{PlusToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '-':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					buffer += string(symbol)
					break
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{MinusToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{AsteriskToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '*':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{AsteriskToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '/':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{SlashToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '%':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{PercentToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '^':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{ExponentiationToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '(':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{LeftParenthesisToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == ')':
			if state == integerPartTokenizerState || state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			if state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokens = append(tokens, token)
				buffer = ""
			}
			token := Token{RightParenthesisToken, string(symbol)}
			tokens = append(tokens, token)
			state = defaultTokenizerState
		case symbol == '.':
		default:
			return nil, fmt.Errorf("unknown symbol %q at position %d", symbol, index)
		}
	}

	return tokens, nil
}
