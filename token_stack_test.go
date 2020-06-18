package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(test *testing.T) {
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

func TestStack_Pop(test *testing.T) {
	testsCases := []struct {
		name      string
		stack     TokenStack
		wantToken Token
		wantOk    bool
	}{
		{
			name: "nonempty",
			stack: []Token{
				{Kind: NumberToken, Value: "1"},
				{Kind: NumberToken, Value: "2"},
			},
			wantToken: Token{Kind: NumberToken, Value: "2"},
			wantOk:    true,
		},
		{
			name:      "empty",
			stack:     []Token{},
			wantToken: Token{},
			wantOk:    false,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotToken, gotOk := testCase.stack.Pop()

			assert.Equal(test, testCase.wantToken, gotToken)
			assert.Equal(test, testCase.wantOk, gotOk)
		})
	}
}
