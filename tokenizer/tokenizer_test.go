package tokenizer

import (
	"testing"

	"github.com/irenicaa/go-calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer(test *testing.T) {
	type args struct {
		code string
	}

	testsCases := []struct {
		name       string
		args       args
		wantTokens []models.Token
		wantErr    string
	}{
		// number
		{
			name:       "integer",
			args:       args{code: "23"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23"}},
			wantErr:    "",
		},
		{
			name:       "fractional",
			args:       args{code: ".23"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: ".23"}},
			wantErr:    "",
		},
		{
			name:       "integer and fractional",
			args:       args{code: "23.42"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23.42"}},
			wantErr:    "",
		},
		{
			name:       "exponent with integers",
			args:       args{code: "23e10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23e10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with integers (in upper case)",
			args:       args{code: "23E10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23E10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with fractionals",
			args:       args{code: "23.42e10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23.42e10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with plus",
			args:       args{code: "23.42e+10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23.42e+10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with plus (in upper case)",
			args:       args{code: "23.42E+10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23.42E+10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with minus",
			args:       args{code: "23.42e-10"},
			wantTokens: []models.Token{{Kind: models.NumberToken, Value: "23.42e-10"}},
			wantErr:    "",
		},

		// identifier
		{
			name:       "identifier",
			args:       args{code: "test"},
			wantTokens: []models.Token{{Kind: models.IdentifierToken, Value: "test"}},
			wantErr:    "",
		},
		{
			name:       "identifier with underscore at the start",
			args:       args{code: "_test"},
			wantTokens: []models.Token{{Kind: models.IdentifierToken, Value: "_test"}},
			wantErr:    "",
		},
		{
			name:       "identifier with underscore in the middle",
			args:       args{code: "test_23"},
			wantTokens: []models.Token{{Kind: models.IdentifierToken, Value: "test_23"}},
			wantErr:    "",
		},
		{
			name: "identifier with integers",
			args: args{code: "23test"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.IdentifierToken, Value: "test"},
			},
			wantErr: "",
		},
		{
			name: "identifier with fractionals",
			args: args{code: "23.5test"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.IdentifierToken, Value: "test"},
			},
			wantErr: "",
		},
		{
			name: "identifier with exponents",
			args: args{code: "23.5e10test"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.IdentifierToken, Value: "test"},
			},
			wantErr: "",
		},
		{
			name:       "identifier with error (integer and fractional parts are empty)",
			args:       args{code: ".test"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "identifier with error (exponent part are empty)",
			args:       args{code: "23etest"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// space
		{
			name: "space with integers",
			args: args{code: "23 42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "space with fractionals",
			args: args{code: "23.5 42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "space with exponents",
			args: args{code: "23.5e10 42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "space with identifers",
			args: args{code: "one two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "space with error (integer and fractional parts are empty)",
			args:       args{code: ". 23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "space with error (exponent part are empty)",
			args:       args{code: "23e 42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// plus
		{
			name:       "plus",
			args:       args{code: "+"},
			wantTokens: []models.Token{{Kind: models.PlusToken, Value: "+"}},
			wantErr:    "",
		},
		{
			name: "plus with integers",
			args: args{code: "23+42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.PlusToken, Value: "+"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "plus with fractionals",
			args: args{code: "23.5+42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.PlusToken, Value: "+"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "plus with exponents",
			args: args{code: "23.5e10+42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.PlusToken, Value: "+"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "plus with identifers",
			args: args{code: "one+two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.PlusToken, Value: "+"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "plus with error (integer and fractional parts are empty)",
			args:       args{code: ".+23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},

		// minus
		{
			name:       "minus",
			args:       args{code: "-"},
			wantTokens: []models.Token{{Kind: models.MinusToken, Value: "-"}},
			wantErr:    "",
		},
		{
			name: "minus with integers",
			args: args{code: "23-42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.MinusToken, Value: "-"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "minus with fractionals",
			args: args{code: "23.5-42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.MinusToken, Value: "-"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "minus with exponents",
			args: args{code: "23.5e10-42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.MinusToken, Value: "-"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "minus with identifers",
			args: args{code: "one-two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.MinusToken, Value: "-"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "minus with error (integer and fractional parts are empty)",
			args:       args{code: ".-23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},

		// asterisk
		{
			name:       "asterisk",
			args:       args{code: "*"},
			wantTokens: []models.Token{{Kind: models.AsteriskToken, Value: "*"}},
			wantErr:    "",
		},
		{
			name: "asterisk with integers",
			args: args{code: "23*42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.AsteriskToken, Value: "*"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with fractionals",
			args: args{code: "23.5*42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.AsteriskToken, Value: "*"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with exponents",
			args: args{code: "23.5e10*42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.AsteriskToken, Value: "*"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with identifers",
			args: args{code: "one*two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.AsteriskToken, Value: "*"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "asterisk with error (integer and fractional parts are empty)",
			args:       args{code: ".*23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "asterisk with error (exponent part are empty)",
			args:       args{code: "23e*42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// slash
		{
			name:       "slash",
			args:       args{code: "/"},
			wantTokens: []models.Token{{Kind: models.SlashToken, Value: "/"}},
			wantErr:    "",
		},
		{
			name: "slash with integers",
			args: args{code: "23/42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.SlashToken, Value: "/"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "slash with fractionals",
			args: args{code: "23.5/42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.SlashToken, Value: "/"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "slash with exponents",
			args: args{code: "23.5e10/42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.SlashToken, Value: "/"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "slash with identifers",
			args: args{code: "one/two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.SlashToken, Value: "/"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "slash with error (integer and fractional parts are empty)",
			args:       args{code: "./23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "slash with error (exponent part are empty)",
			args:       args{code: "23e/42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// percent
		{
			name:       "percent",
			args:       args{code: "%"},
			wantTokens: []models.Token{{Kind: models.PercentToken, Value: "%"}},
			wantErr:    "",
		},
		{
			name: "percent with integers",
			args: args{code: "23%42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.PercentToken, Value: "%"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "percent with fractionals",
			args: args{code: "23.5%42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.PercentToken, Value: "%"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "percent with exponents",
			args: args{code: "23.5e10%42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.PercentToken, Value: "%"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "percent with identifers",
			args: args{code: "one%two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.PercentToken, Value: "%"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "percent with error (integer and fractional parts are empty)",
			args:       args{code: ".%23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "percent with error (exponent part are empty)",
			args:       args{code: "23e%42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// exponentiation
		{
			name:       "exponentiation",
			args:       args{code: "^"},
			wantTokens: []models.Token{{Kind: models.ExponentiationToken, Value: "^"}},
			wantErr:    "",
		},
		{
			name: "exponentiation with integers",
			args: args{code: "23^42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.ExponentiationToken, Value: "^"},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with fractionals",
			args: args{code: "23.5^42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.ExponentiationToken, Value: "^"},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with exponents",
			args: args{code: "23.5e10^42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.ExponentiationToken, Value: "^"},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with identifers",
			args: args{code: "one^two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.ExponentiationToken, Value: "^"},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with error" +
				"(integer and fractional parts are empty)",
			args:       args{code: ".^23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "exponentiation with error (exponent part are empty)",
			args:       args{code: "23e^42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// parentheses
		{
			name: "parentheses with integers",
			args: args{code: "23(42)"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.LeftParenthesisToken, Value: "("},
				{Kind: models.NumberToken, Value: "42"},
				{Kind: models.RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with fractionals",
			args: args{code: "23.5(42.5)"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.LeftParenthesisToken, Value: "("},
				{Kind: models.NumberToken, Value: "42.5"},
				{Kind: models.RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with exponents",
			args: args{code: "23.5e10(42.5e10)"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.LeftParenthesisToken, Value: "("},
				{Kind: models.NumberToken, Value: "42.5e10"},
				{Kind: models.RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with identifers",
			args: args{code: "one(two)"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.LeftParenthesisToken, Value: "("},
				{Kind: models.IdentifierToken, Value: "two"},
				{Kind: models.RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "left parenthesis with error" +
				" (integer and fractional parts are empty)",
			args:       args{code: ".(23)"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name: "right parenthesis with error" +
				"(integer and fractional parts are empty)",
			args:       args{code: "23(.)"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 4",
		},
		{
			name:       "left parenthesis with error (exponent part are empty)",
			args:       args{code: "23e(42)"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},
		{
			name:       "right parenthesis with error (exponent part are empty)",
			args:       args{code: "23(42e)"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 6",
		},

		// comma
		{
			name: "comma with integers",
			args: args{code: "23,42"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.CommaToken, Value: ","},
				{Kind: models.NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "comma with fractionals",
			args: args{code: "23.5,42.5"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5"},
				{Kind: models.CommaToken, Value: ","},
				{Kind: models.NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "comma with exponents",
			args: args{code: "23.5e10,42.5e10"},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23.5e10"},
				{Kind: models.CommaToken, Value: ","},
				{Kind: models.NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "comma with identifers",
			args: args{code: "one,two"},
			wantTokens: []models.Token{
				{Kind: models.IdentifierToken, Value: "one"},
				{Kind: models.CommaToken, Value: ","},
				{Kind: models.IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "comma with error (integer and fractional parts are empty)",
			args:       args{code: ".,23"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "comma with error (exponent part are empty)",
			args:       args{code: "23e,42"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},

		// misc. errors
		{
			name:       "error with a fractional point after fractional part",
			args:       args{code: "23.42."},
			wantTokens: nil,
			wantErr:    "unexpected fractional point at position 5",
		},
		{
			name:       "error with a fractional point after exponent part",
			args:       args{code: "23.42e10."},
			wantTokens: nil,
			wantErr:    "unexpected fractional point at position 8",
		},
		{
			name:       "error with a fractional point after identifier part",
			args:       args{code: "test."},
			wantTokens: nil,
			wantErr:    "unexpected fractional point at position 4",
		},
		{
			name:       "error with an unknown symbol",
			args:       args{code: "23!"},
			wantTokens: nil,
			wantErr:    "unknown symbol '!' at position 2",
		},
		{
			name:       "error with empty integer and fractional parts at EOI",
			args:       args{code: "."},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at EOI",
		},
		{
			name:       "error with an empty exponent part at EOI",
			args:       args{code: "23.42e"},
			wantTokens: nil,
			wantErr:    "empty exponent part at EOI",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotTokens := []models.Token(nil)

			tokenizer := Tokenizer{}
			tokens, gotErr := tokenizer.Tokenize(testCase.args.code)
			gotTokens = append(gotTokens, tokens...)

			if gotErr == nil {
				tokens, gotErr = tokenizer.Finalize()
				gotTokens = append(gotTokens, tokens...)
			}

			assert.Equal(test, testCase.wantTokens, gotTokens)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}

func TestTokenizer_withSequentialCalls(test *testing.T) {
	type args struct {
		codeParts []string
	}

	testsCases := []struct {
		name       string
		args       args
		wantTokens []models.Token
	}{
		{
			name: "different tokens in separate parts",
			args: args{codeParts: []string{"23", "test"}},
			wantTokens: []models.Token{
				{Kind: models.NumberToken, Value: "23"},
				{Kind: models.IdentifierToken, Value: "test"},
			},
		},
		{
			name:       "single token in separate parts",
			args:       args{codeParts: []string{"test", "23"}},
			wantTokens: []models.Token{{Kind: models.IdentifierToken, Value: "test23"}},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotTokens, gotErr := []models.Token(nil), error(nil)

			tokenizer := Tokenizer{}
			for _, codePart := range testCase.args.codeParts {
				tokens, err := tokenizer.Tokenize(codePart)

				gotTokens = append(gotTokens, tokens...)
				gotErr = err
				if gotErr != nil {
					break
				}
			}
			if gotErr == nil {
				tokens, err := tokenizer.Finalize()

				gotTokens = append(gotTokens, tokens...)
				gotErr = err
			}

			assert.Equal(test, testCase.wantTokens, gotTokens)
			assert.NoError(test, gotErr)
		})
	}
}
