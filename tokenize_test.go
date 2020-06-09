package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(test *testing.T) {
	type args struct {
		code string
	}

	testsCases := []struct {
		name       string
		args       args
		wantTokens []Token
		wantErr    string
	}{
		// number
		{
			name:       "integer",
			args:       args{code: "23"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23"}},
			wantErr:    "",
		},
		{
			name:       "fractional",
			args:       args{code: ".23"},
			wantTokens: []Token{{Kind: NumberToken, Value: ".23"}},
			wantErr:    "",
		},
		{
			name:       "integer and fractional",
			args:       args{code: "23.42"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23.42"}},
			wantErr:    "",
		},
		{
			name:       "exponent with integers",
			args:       args{code: "23e10"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23e10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with fractionals",
			args:       args{code: "23.42e10"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23.42e10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with plus",
			args:       args{code: "23.42e+10"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23.42e+10"}},
			wantErr:    "",
		},
		{
			name:       "exponent with minus",
			args:       args{code: "23.42e-10"},
			wantTokens: []Token{{Kind: NumberToken, Value: "23.42e-10"}},
			wantErr:    "",
		},

		// identifier
		{
			name:       "identifier",
			args:       args{code: "test"},
			wantTokens: []Token{{Kind: IdentifierToken, Value: "test"}},
			wantErr:    "",
		},
		{
			name:       "identifier with underscore at the start",
			args:       args{code: "_test"},
			wantTokens: []Token{{Kind: IdentifierToken, Value: "_test"}},
			wantErr:    "",
		},
		{
			name:       "identifier with underscore in the middle",
			args:       args{code: "test_23"},
			wantTokens: []Token{{Kind: IdentifierToken, Value: "test_23"}},
			wantErr:    "",
		},
		{
			name: "identifier with integers",
			args: args{code: "23test"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: IdentifierToken, Value: "test"},
			},
			wantErr: "",
		},
		{
			name: "identifier with fractionals",
			args: args{code: "23.5test"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: IdentifierToken, Value: "test"},
			},
			wantErr: "",
		},
		{
			name: "identifier with exponents",
			args: args{code: "23.5e10test"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: IdentifierToken, Value: "test"},
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

		//space
		{
			name: "space with integers",
			args: args{code: "23 42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "space with fractionals",
			args: args{code: "23.5 42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "space with exponents",
			args: args{code: "23.5e10 42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "space with identifers",
			args: args{code: "one two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: PlusToken, Value: "+"}},
			wantErr:    "",
		},
		{
			name: "plus with integers",
			args: args{code: "23+42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: PlusToken, Value: "+"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "plus with fractionals",
			args: args{code: "23.5+42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: PlusToken, Value: "+"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "plus with exponents",
			args: args{code: "23.5e10+42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: PlusToken, Value: "+"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "plus with identifers",
			args: args{code: "one+two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: PlusToken, Value: "+"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: MinusToken, Value: "-"}},
			wantErr:    "",
		},
		{
			name: "minus with integers",
			args: args{code: "23-42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: MinusToken, Value: "-"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "minus with fractionals",
			args: args{code: "23.5-42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: MinusToken, Value: "-"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "minus with exponents",
			args: args{code: "23.5e10-42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: MinusToken, Value: "-"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "minus with identifers",
			args: args{code: "one-two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: MinusToken, Value: "-"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: AsteriskToken, Value: "*"}},
			wantErr:    "",
		},
		{
			name: "asterisk with integers",
			args: args{code: "23*42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: AsteriskToken, Value: "*"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with fractionals",
			args: args{code: "23.5*42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: AsteriskToken, Value: "*"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with exponents",
			args: args{code: "23.5e10*42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: AsteriskToken, Value: "*"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "asterisk with identifers",
			args: args{code: "one*two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: AsteriskToken, Value: "*"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: SlashToken, Value: "/"}},
			wantErr:    "",
		},
		{
			name: "slash with integers",
			args: args{code: "23/42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: SlashToken, Value: "/"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "slash with fractionals",
			args: args{code: "23.5/42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: SlashToken, Value: "/"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "slash with exponents",
			args: args{code: "23.5e10/42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: SlashToken, Value: "/"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "slash with identifers",
			args: args{code: "one/two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: SlashToken, Value: "/"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: PercentToken, Value: "%"}},
			wantErr:    "",
		},
		{
			name: "percent with integers",
			args: args{code: "23%42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: PercentToken, Value: "%"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "percent with fractionals",
			args: args{code: "23.5%42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: PercentToken, Value: "%"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "percent with exponents",
			args: args{code: "23.5e10%42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: PercentToken, Value: "%"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "percent with identifers",
			args: args{code: "one%two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: PercentToken, Value: "%"},
				{Kind: IdentifierToken, Value: "two"},
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
			wantTokens: []Token{{Kind: ExponentiationToken, Value: "^"}},
			wantErr:    "",
		},
		{
			name: "exponentiation with integers",
			args: args{code: "23^42"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: ExponentiationToken, Value: "^"},
				{Kind: NumberToken, Value: "42"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with fractionals",
			args: args{code: "23.5^42.5"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: ExponentiationToken, Value: "^"},
				{Kind: NumberToken, Value: "42.5"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with exponents",
			args: args{code: "23.5e10^42.5e10"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: ExponentiationToken, Value: "^"},
				{Kind: NumberToken, Value: "42.5e10"},
			},
			wantErr: "",
		},
		{
			name: "exponentiation with identifers",
			args: args{code: "one^two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: ExponentiationToken, Value: "^"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		{
			name:       "exponentiation with error (integer and fractional parts are empty)",
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

		//parentheses
		{
			name: "parentheses with integers",
			args: args{code: "23(42)"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23"},
				{Kind: LeftParenthesisToken, Value: "("},
				{Kind: NumberToken, Value: "42"},
				{Kind: RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with fractionals",
			args: args{code: "23.5(42.5)"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5"},
				{Kind: LeftParenthesisToken, Value: "("},
				{Kind: NumberToken, Value: "42.5"},
				{Kind: RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with exponents",
			args: args{code: "23.5e10(42.5e10)"},
			wantTokens: []Token{
				{Kind: NumberToken, Value: "23.5e10"},
				{Kind: LeftParenthesisToken, Value: "("},
				{Kind: NumberToken, Value: "42.5e10"},
				{Kind: RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name: "parentheses with identifers",
			args: args{code: "one(two)"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: LeftParenthesisToken, Value: "("},
				{Kind: IdentifierToken, Value: "two"},
				{Kind: RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
		},
		{
			name:       "left parenthesis with error (integer and fractional parts are empty)",
			args:       args{code: ".(23)"},
			wantTokens: nil,
			wantErr:    "both integer and fractional parts are empty at position 1",
		},
		{
			name:       "left parenthesis with error (exponent part are empty)",
			args:       args{code: "23e(42)"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 3",
		},
		{
			name:       "left parenthesis with error (exponent part are empty)",
			args:       args{code: "23(42e)"},
			wantTokens: nil,
			wantErr:    "empty exponent part at position 6",
		},

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
			gotTokens, gotErr := Tokenize(testCase.args.code)

			assert.Equal(test, testCase.wantTokens, gotTokens)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}
