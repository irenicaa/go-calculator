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
			name:       "exponent",
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

		// plus
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
			name: "plus with identifer",
			args: args{code: "one+two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: PlusToken, Value: "+"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		// minus
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
			name: "minus with identifer",
			args: args{code: "one-two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: MinusToken, Value: "-"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		// asterisk
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
			name: "asterisk with identifer",
			args: args{code: "one*two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: AsteriskToken, Value: "*"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		// slash
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
			name: "slash with identifer",
			args: args{code: "one/two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: SlashToken, Value: "/"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		// percent
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
			name: "percent with identifer",
			args: args{code: "one%two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: PercentToken, Value: "%"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},
		// exponentiation
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
			name: "exponentiation with identifer",
			args: args{code: "one^two"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "one"},
				{Kind: ExponentiationToken, Value: "^"},
				{Kind: IdentifierToken, Value: "two"},
			},
			wantErr: "",
		},

		{
			name: "parentheses and comma",
			args: args{code: "test(one,two)"},
			wantTokens: []Token{
				{Kind: IdentifierToken, Value: "test"},
				{Kind: LeftParenthesisToken, Value: "("},
				{Kind: IdentifierToken, Value: "one"},
				{Kind: CommaToken, Value: ","},
				{Kind: IdentifierToken, Value: "two"},
				{Kind: RightParenthesisToken, Value: ")"},
			},
			wantErr: "",
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
