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

// Tokens ...
func (tokenizer Tokenizer) Tokens() []Token {
	return tokenizer.tokens
}

// Tokenize ...
func (tokenizer *Tokenizer) Tokenize(code string) error {
	for index, symbol := range code {
		position := position(index)
		switch {
		case unicode.IsDigit(symbol):
			if tokenizer.state == defaultTokenizerState {
				tokenizer.state = integerPartTokenizerState
			}
			tokenizer.buffer += string(symbol)
		case unicode.IsLetter(symbol) || symbol == '_':
			switch tokenizer.state {
			case integerPartTokenizerState, fractionalPartTokenizerState:
				if tokenizer.areIntegerAndFractionalEmpty() {
					return fmt.Errorf(
						"both integer and fractional parts are empty at position %d",
						index,
					)
				}
				if symbol == 'e' || symbol == 'E' {
					tokenizer.state = exponentTokenizerState
					tokenizer.buffer += string(symbol)
					continue
				}

				tokenizer.addTokenFromBuffer(NumberToken)
			case exponentTokenizerState:
				if tokenizer.isExponentEmpty() {
					return fmt.Errorf("empty exponent part at position %d", index)
				}

				tokenizer.addTokenFromBuffer(NumberToken)
			}

			tokenizer.state = identifierTokenizerState
			tokenizer.buffer += string(symbol)
		case unicode.IsSpace(symbol):
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			tokenizer.state = defaultTokenizerState
		case symbol == '+':
			if tokenizer.state == integerPartTokenizerState || tokenizer.state == fractionalPartTokenizerState {
				if tokenizer.areIntegerAndFractionalEmpty() {
					return fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == exponentTokenizerState {
				if tokenizer.isExponentEmpty() {
					tokenizer.buffer += string(symbol)
					continue
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
				if tokenizer.areIntegerAndFractionalEmpty() {
					return fmt.Errorf("both integer and fractional parts are empty at position %d", index)
				}
				tokenizer.addTokenFromBuffer(NumberToken)
			}
			if tokenizer.state == exponentTokenizerState {
				if tokenizer.isExponentEmpty() {
					tokenizer.buffer += string(symbol)
					continue
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
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{AsteriskToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '/':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{SlashToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '%':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{PercentToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '^':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{ExponentiationToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '(':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{LeftParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ')':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{RightParenthesisToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == ',':
			err := tokenizer.resetBuffer(position)
			if err != nil {
				return err
			}

			token := Token{CommaToken, string(symbol)}
			tokenizer.tokens = append(tokenizer.tokens, token)
			tokenizer.state = defaultTokenizerState
		case symbol == '.':
			if tokenizer.state == defaultTokenizerState || tokenizer.state == integerPartTokenizerState {
				tokenizer.state = fractionalPartTokenizerState
				tokenizer.buffer += string(symbol)
				continue
			}

			return fmt.Errorf("unexpected fractional point at position %d", index)
		default:
			return fmt.Errorf("unknown symbol %q at position %d", symbol, index)
		}
	}

	return tokenizer.resetBuffer(eoi)
}

func (tokenizer Tokenizer) areIntegerAndFractionalEmpty() bool {
	return tokenizer.buffer == "."
}

func (tokenizer Tokenizer) isExponentEmpty() bool {
	lastSymbol := tokenizer.buffer[len(tokenizer.buffer)-1]
	return lastSymbol == 'e' || lastSymbol == 'E'
}

func (tokenizer *Tokenizer) addTokenFromBuffer(kind TokenKind) {
	token := Token{kind, tokenizer.buffer}
	tokenizer.tokens = append(tokenizer.tokens, token)

	tokenizer.buffer = ""
}

func (tokenizer *Tokenizer) resetBuffer(symbolIndex position) error {
	switch tokenizer.state {
	case integerPartTokenizerState, fractionalPartTokenizerState:
		if tokenizer.areIntegerAndFractionalEmpty() {
			return fmt.Errorf(
				"both integer and fractional parts are empty at %s",
				symbolIndex,
			)
		}

		tokenizer.addTokenFromBuffer(NumberToken)
	case exponentTokenizerState:
		if tokenizer.isExponentEmpty() {
			return fmt.Errorf("empty exponent part at %s", symbolIndex)
		}

		tokenizer.addTokenFromBuffer(NumberToken)
	case identifierTokenizerState:
		tokenizer.addTokenFromBuffer(IdentifierToken)
	}

	return nil
}
