package calculator

import (
	"errors"
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
	CommaToken
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

// Tokinizer ...
type Tokenizer struct {
	tokens []Token
	state  tokenizerState
}

// Tokenize ...
func (tokenizer Tokenizer) Tokenize(code string) ([]Token, error) {
	buffer := ""
	for index, symbol := range code {
		switch {
		case unicode.IsDigit(symbol):
			if tokenizer.state == defaultTokenizerState {
				tokenizer.state = integerPartTokenizerState
			}
			buffer += string(symbol)
		case unicode.IsLetter(symbol) || symbol == '_':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				if symbol == 'e' || symbol == 'E' {
					tokenizer.state = exponentTokenizerState
					buffer += string(symbol)
					break
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			tokenizer.state = identifierTokenizerState
			buffer += string(symbol)
		case unicode.IsSpace(symbol):
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			tokenizer.state = defaultTokenizerState
		case symbol == '+':
			if buffer == "." {
				return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
			}
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					buffer += string(symbol)
					break
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{PlusToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '-':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					buffer += string(symbol)
					break
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{MinusToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '*':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{AsteriskToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '/':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{SlashToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '%':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{PercentToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '^':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{ExponentiationToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '(':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{LeftParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ')':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{RightParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ',':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == exponentTokenizerState {
				lastSymbol := buffer[len(buffer)-1]
				if lastSymbol == 'e' || lastSymbol == 'E' {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				token := Token{NumberToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			if tokenizer.state == identifierTokenizerState {
				token := Token{IdentifierToken, buffer}
				tokenizer.tokens = append(tokenizer.tokens, token)
				buffer = ""
			}
			token := Token{CommaToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '.':
			if tokenizer.state == defaultTokenizerState || tokenizer.state == integerPartTokenizerState {
				tokenizer.state = fractionalPartTokenizerState
				buffer += string(symbol)
				break
			}
			return nil, fmt.Errorf("unexpected fractional point at position %d", index)
		default:
			return nil, fmt.Errorf("unknown symbol %q at position %d", symbol, index)
		}
	}
	if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
		if buffer == "." {
			return nil, errors.New("both integer and fractional parts are empty at EOI")
		}
		token := Token{NumberToken, buffer}
		tokenizer.tokens = append(tokenizer.tokens, token)
	}
	if tokenizer.state == exponentTokenizerState {
		lastSymbol := buffer[len(buffer)-1]
		if lastSymbol == 'e' || lastSymbol == 'E' {
			return nil, errors.New("empty exponent part at EOI")
		}
		token := Token{NumberToken, buffer}
		tokenizer.tokens = append(tokenizer.tokens, token)
	}
	if tokenizer.state == identifierTokenizerState {
		token := Token{IdentifierToken, buffer}
		tokenizer.tokens = append(tokenizer.tokens, token)
	}

	return tokenizer.tokens, nil
}
