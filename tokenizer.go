package calculator

import (
	"fmt"
	"strings"
	"unicode"
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
func (tokenizer *Tokenizer) Tokenize(code string) error {
	for symbolIndex, symbol := range code {
		symbolPosition := position(symbolIndex)
		switch {
		case unicode.IsDigit(symbol):
			if tokenizer.state == defaultTokenizerState {
				tokenizer.state = integerPartTokenizerState
			}
			tokenizer.buffer += string(symbol)
		case unicode.IsLetter(symbol), symbol == '_':
			switch tokenizer.state {
			case integerPartTokenizerState, fractionalPartTokenizerState:
				if tokenizer.areIntegerAndFractionalEmpty() {
					return fmt.Errorf(
						"both integer and fractional parts are empty at position %d",
						symbolIndex,
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
					return fmt.Errorf("empty exponent part at position %d", symbolIndex)
				}

				tokenizer.addTokenFromBuffer(NumberToken)
			}

			tokenizer.state = identifierTokenizerState
			tokenizer.buffer += string(symbol)
		case unicode.IsSpace(symbol):
			err := tokenizer.resetBuffer(symbolPosition)
			if err != nil {
				return err
			}

			tokenizer.state = defaultTokenizerState
		case strings.ContainsRune("+-", symbol):
			if tokenizer.state == exponentTokenizerState && tokenizer.isExponentEmpty() {
				tokenizer.buffer += string(symbol)
				continue
			}

			fallthrough
		case strings.ContainsRune("*/%^(),", symbol):
			err := tokenizer.resetBuffer(symbolPosition)
			if err != nil {
				return err
			}

			tokenizer.addTokenFromSymbol(symbol)
		case symbol == '.':
			switch tokenizer.state {
			case defaultTokenizerState, integerPartTokenizerState:
				tokenizer.state = fractionalPartTokenizerState
				tokenizer.buffer += string(symbol)
				continue
			}

			return fmt.Errorf("unexpected fractional point at position %d", symbolIndex)
		default:
			return fmt.Errorf("unknown symbol %q at position %d", symbol, symbolIndex)
		}
	}

	return nil
}

// Finalize ...
func (tokenizer *Tokenizer) Finalize() ([]Token, error) {
	err := tokenizer.resetBuffer(eoi)
	if err != nil {
		return nil, err
	}

	return tokenizer.tokens, nil
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

func (tokenizer *Tokenizer) addTokenFromSymbol(symbol rune) {
	kind, _ := ParseTokenKind(symbol)
	token := Token{kind, string(symbol)}
	tokenizer.tokens = append(tokenizer.tokens, token)

	tokenizer.state = defaultTokenizerState
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
