package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(test *testing.T) {
	type args struct {
		tokens    []Token
		functions map[string]struct{}
	}

	testsCases := []struct {
		name         string
		args         args
		wantCommands []Command
		wantErr      string
	}{}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotCommands, gotErr := Translate(
				testCase.args.tokens,
				testCase.args.functions,
			)

			assert.Equal(test, testCase.wantCommands, gotCommands)
			assert.NoError(test, gotErr)
		})
	}
}
