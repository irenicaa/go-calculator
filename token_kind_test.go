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
