package containers

import (
	"testing"

	"github.com/irenicaa/go-calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestTokenStack_Push(test *testing.T) {
	type args struct {
		token models.Token
	}

	testsCases := []struct {
		name      string
		stack     TokenStack
		args      args
		wantStack TokenStack
	}{
		{
			name: "nonempty",
			stack: []models.Token{
				{Kind: models.NumberToken, Value: "1"},
				{Kind: models.NumberToken, Value: "2"},
			},
			args: args{token: models.Token{Kind: models.NumberToken, Value: "3"}},
			wantStack: []models.Token{
				{Kind: models.NumberToken, Value: "1"},
				{Kind: models.NumberToken, Value: "2"},
				{Kind: models.NumberToken, Value: "3"},
			},
		},
		{
			name:      "empty",
			stack:     []models.Token{},
			args:      args{token: models.Token{Kind: models.NumberToken, Value: "3"}},
			wantStack: []models.Token{{Kind: models.NumberToken, Value: "3"}},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			testCase.stack.Push(testCase.args.token)

			assert.Equal(test, testCase.wantStack, testCase.stack)
		})
	}
}

func TestTokenStack_Pop(test *testing.T) {
	testsCases := []struct {
		name      string
		stack     TokenStack
		wantToken models.Token
		wantOk    bool
	}{
		{
			name: "nonempty",
			stack: []models.Token{
				{Kind: models.NumberToken, Value: "1"},
				{Kind: models.NumberToken, Value: "2"},
			},
			wantToken: models.Token{Kind: models.NumberToken, Value: "2"},
			wantOk:    true,
		},
		{
			name:      "empty",
			stack:     []models.Token{},
			wantToken: models.Token{},
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
