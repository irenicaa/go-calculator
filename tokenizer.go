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

// Tokenizer ...
type Tokenizer struct {
	tokens []Token
	state  tokenizerState
	buffer string
}

// Tokenize ...
func (tokenizer *Tokenizer) Tokenize(code string) ([]Token, error) {
	for index, symbol := range code {
		switch {
		case unicode.IsDigit(symbol):
			if tokenizer.state == defaultTokenizerState {
				tokenizer.state = integerPartTokenizerState
			}
			tokenizer.buffer += string(symbol)
		case unicode.IsLetter(symbol) || symbol == '_':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if tokenizer.buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				if symbol == 'e' || symbol == 'E' {
					tokenizer.state = exponentTokenizerState
					tokenizer.buffer += string(symbol)
					break
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == exponentTokenizerState {
				if tokenizer.isExponentEmpty() {
					return nil, fmt.Errorf("empty exponent part at position %d", index)
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			tokenizer.state = identifierTokenizerState
			tokenizer.buffer += string(symbol)
		case unicode.IsSpace(symbol):
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}
			tokenizer.state = defaultTokenizerState
		case symbol == '+':
			if tokenizer.buffer == "." {
				return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
			}
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == exponentTokenizerState {
				if tokenizer.isExponentEmpty() {
					tokenizer.buffer += string(symbol)
					break
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == identifierTokenizerState {
				tokenizer.addTokenFromBuffer(IdentifierToken)
			}
			token := Token{PlusToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '-':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if tokenizer.buffer == "." {
					return nil, fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == exponentTokenizerState {
				if tokenizer.isExponentEmpty() {
					tokenizer.buffer += string(symbol)
					break
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == identifierTokenizerState {
				tokenizer.addTokenFromBuffer(IdentifierToken)
			}
			token := Token{MinusToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '*':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{AsteriskToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '/':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{SlashToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '%':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{PercentToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '^':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{ExponentiationToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '(':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{LeftParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ')':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{RightParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ',':
			err := tokenizer.resetBuffer(index)
			if err != nil {
				return nil, err
			}

			token := Token{CommaToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '.':
			if tokenizer.state == defaultTokenizerState || tokenizer.state == integerPartTokenizerState {
				tokenizer.state = fractionalPartTokenizerState
				tokenizer.buffer += string(symbol)
				break
			}
			return nil, fmt.Errorf("unexpected fractional point at position %d", index)
		default:
			return nil, fmt.Errorf("unknown symbol %q at position %d", symbol, index)
		}
	}
	if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
		if tokenizer.buffer == "." {
			return nil, errors.New("both integer and fractional parts are empty at EOI")
		}
		tokenizer.addTokenFromBuffer(NumberToken)
	}
	if tokenizer.state == exponentTokenizerState {
		if tokenizer.isExponentEmpty() {
			return nil, errors.New("empty exponent part at EOI")
		}

		tokenizer.addTokenFromBuffer(NumberToken)
	}
	if tokenizer.state == identifierTokenizerState {
		tokenizer.addTokenFromBuffer(IdentifierToken)
	}

	return tokenizer.tokens, nil
}

func (tokenizer *Tokenizer) addTokenFromBuffer(kind TokenKind) {
	token := Token{kind, tokenizer.buffer}
	tokenizer.tokens = append(tokenizer.tokens, token)

	tokenizer.buffer = ""
}

func (tokenizer Tokenizer) isExponentEmpty() bool {
	lastSymbol := tokenizer.buffer[len(tokenizer.buffer)-1]
	return lastSymbol == 'e' || lastSymbol == 'E'
}

func (tokenizer *Tokenizer) resetBuffer(symbolIndex int) error {
	switch tokenizer.state {
	case integerPartTokenizerState, fractionalPartTokenizerState:
		if tokenizer.buffer == "." {
			return fmt.Errorf(
				"both integer and fractional parts are empty at position %d",
				symbolIndex,
			)
		}

		tokenizer.addTokenFromBuffer(NumberToken)
	case exponentTokenizerState:
		if tokenizer.isExponentEmpty() {
			return fmt.Errorf("empty exponent part at position %d", symbolIndex)
		}

		tokenizer.addTokenFromBuffer(NumberToken)
	case identifierTokenizerState:
		tokenizer.addTokenFromBuffer(IdentifierToken)
	}

	return nil
}
