package tokenizer

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/irenicaa/go-calculator/models"
)

type tokenizerState int

const (
	defaultTokenizerState tokenizerState = iota
	integerPartTokenizerState
	fractionalPartTokenizerState
	exponentTokenizerState
	identifierTokenizerState
)

// Tokenizer ...
type Tokenizer struct {
	tokens []models.Token
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
				if unicode.ToLower(symbol) == 'e' {
					tokenizer.state = exponentTokenizerState
					tokenizer.buffer += string(symbol)
					continue
				}
			}
			if tokenizer.state != identifierTokenizerState {
				err := tokenizer.resetBuffer(symbolPosition)
				if err != nil {
					return err
				}
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

			return fmt.Errorf("unexpected fractional point at %s", symbolPosition)
		default:
			return fmt.Errorf("unknown symbol %q at %s", symbol, symbolPosition)
		}
	}

	return nil
}

// Finalize ...
func (tokenizer *Tokenizer) Finalize() ([]models.Token, error) {
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
	return unicode.ToLower(rune(lastSymbol)) == 'e'
}

func (tokenizer *Tokenizer) addTokenFromBuffer(kind models.TokenKind) {
	token := models.Token{Kind: kind, Value: tokenizer.buffer}
	tokenizer.tokens = append(tokenizer.tokens, token)

	tokenizer.buffer = ""
}

func (tokenizer *Tokenizer) addTokenFromSymbol(symbol rune) {
	// lack of the error is guaranteed by the calling function
	kind, _ := models.ParseTokenKind(symbol)
	token := models.Token{Kind: kind, Value: string(symbol)}
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

		tokenizer.addTokenFromBuffer(models.NumberToken)
	case exponentTokenizerState:
		if tokenizer.isExponentEmpty() {
			return fmt.Errorf("empty exponent part at %s", symbolIndex)
		}

		tokenizer.addTokenFromBuffer(models.NumberToken)
	case identifierTokenizerState:
		tokenizer.addTokenFromBuffer(models.IdentifierToken)
	}

	return nil
}
