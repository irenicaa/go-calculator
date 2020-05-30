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
		wantErr    error
	}{}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotTokens, gotErr := Tokenize(testCase.args.code)
			assert.Equal(test, testCase.wantTokens, gotTokens)
			assert.Equal(test, testCase.wantErr, gotErr)
		})
	}
}
