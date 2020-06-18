package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTokenKind(test *testing.T) {
	type args struct {
		symbol rune
	}

	testsCases := []struct {
		name          string
		args          args
		wantTokenKind TokenKind
		wantErr       string
	}{
		{
			name:          "plus",
			args:          args{symbol: '+'},
			wantTokenKind: PlusToken,
			wantErr:       "",
		},
		{
			name:          "minus",
			args:          args{symbol: '-'},
			wantTokenKind: MinusToken,
			wantErr:       "",
		},
		{
			name:          "asterisk",
			args:          args{symbol: '*'},
			wantTokenKind: AsteriskToken,
			wantErr:       "",
		},
		{
			name:          "slash",
			args:          args{symbol: '/'},
			wantTokenKind: SlashToken,
			wantErr:       "",
		},
		{
			name:          "percent",
			args:          args{symbol: '%'},
			wantTokenKind: PercentToken,
			wantErr:       "",
		},
		{
			name:          "exponentiation",
			args:          args{symbol: '^'},
			wantTokenKind: ExponentiationToken,
			wantErr:       "",
		},
		{
			name:          "left parenthesis",
			args:          args{symbol: '('},
			wantTokenKind: LeftParenthesisToken,
			wantErr:       "",
		},
		{
			name:          "right parenthesis",
			args:          args{symbol: ')'},
			wantTokenKind: RightParenthesisToken,
			wantErr:       "",
		},
		{
			name:          "comma",
			args:          args{symbol: ','},
			wantTokenKind: CommaToken,
			wantErr:       "",
		},
		{
			name:          "error",
			args:          args{symbol: '!'},
			wantTokenKind: 0,
			wantErr:       "unknown symbol '!'",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotTokenKind, gotErr := ParseTokenKind(testCase.args.symbol)

			assert.Equal(test, testCase.wantTokenKind, gotTokenKind)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}

func TestKind_IsOperator(test *testing.T) {
	testsCases := []struct {
		name   string
		kind   TokenKind
		wantOk bool
	}{
		{
			name:   "plus",
			kind:   PlusToken,
			wantOk: true,
		},
		{
			name:   "minus",
			kind:   MinusToken,
			wantOk: true,
		},
		{
			name:   "asteriks",
			kind:   AsteriskToken,
			wantOk: true,
		},
		{
			name:   "slash",
			kind:   SlashToken,
			wantOk: true,
		},
		{
			name:   "percent",
			kind:   PercentToken,
			wantOk: true,
		},
		{
			name:   "exponent",
			kind:   ExponentiationToken,
			wantOk: true,
		},
		{
			name:   "not operator",
			kind:   LeftParenthesisToken,
			wantOk: false,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			ok := testCase.kind.IsOperator()

			assert.Equal(test, testCase.wantOk, ok)
		})
	}
}

func TestKind_Precedence(test *testing.T) {
	testsCases := []struct {
		name           string
		kind           TokenKind
		wantPrecedence int
	}{
		{
			name:           "plus",
			kind:           PlusToken,
			wantPrecedence: 1,
		},
		{
			name:           "minus",
			kind:           MinusToken,
			wantPrecedence: 1,
		},
		{
			name:           "asteriks",
			kind:           AsteriskToken,
			wantPrecedence: 2,
		},
		{
			name:           "slash",
			kind:           SlashToken,
			wantPrecedence: 2,
		},
		{
			name:           "percent",
			kind:           PercentToken,
			wantPrecedence: 2,
		},
		{
			name:           "exponent",
			kind:           ExponentiationToken,
			wantPrecedence: 3,
		},
		{
			name:           "not operator",
			kind:           LeftParenthesisToken,
			wantPrecedence: 0,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotPrecedence := testCase.kind.Precedence()

			assert.Equal(test, testCase.wantPrecedence, gotPrecedence)
		})
	}
}
