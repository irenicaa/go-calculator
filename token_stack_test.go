package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(test *testing.T) {
	type args struct {
		token Token
	}

	testsCases := []struct {
		name      string
		stack     TokenStack
		args      args
		wantStack TokenStack
	}{
		{
			name: "nonempty",
			stack: []Token{
				{Kind: NumberToken, Value: "1"},
				{Kind: NumberToken, Value: "2"},
			},
			args: args{token: Token{Kind: NumberToken, Value: "3"}},
			wantStack: []Token{
				{Kind: NumberToken, Value: "1"},
				{Kind: NumberToken, Value: "2"},
				{Kind: NumberToken, Value: "3"},
			},
		},
		{
			name:      "empty",
			stack:     []Token{},
			args:      args{token: Token{Kind: NumberToken, Value: "3"}},
			wantStack: []Token{{Kind: NumberToken, Value: "3"}},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			testCase.stack.Push(testCase.args.token)

			assert.Equal(test, testCase.wantStack, testCase.stack)
		})
	}
}
