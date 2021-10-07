package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractVariable(test *testing.T) {
	type args struct {
		input string
	}

	testsCases := []struct {
		name         string
		args         args
		wantVariable string
		wantInput    string
	}{
		{
			name:         "string with separator",
			args:         args{"test1 = test2 + 1"},
			wantVariable: "test1",
			wantInput:    " test2 + 1",
		},
		{
			name:         "string with separator and without variable",
			args:         args{"= test2 + 1"},
			wantVariable: "",
			wantInput:    " test2 + 1",
		},
		{
			name:         "string without separator",
			args:         args{"test2 + 1"},
			wantVariable: "",
			wantInput:    "test2 + 1",
		},
		{
			name:         "string with several separators",
			args:         args{"test1 = test2 = test3 + 1"},
			wantVariable: "test1",
			wantInput:    " test2 = test3 + 1",
		},
		{
			name:         "empty string",
			args:         args{""},
			wantVariable: "",
			wantInput:    "",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotVariable, gotInput := ExtractVariable(testCase.args.input)

			assert.Equal(test, testCase.wantVariable, gotVariable)
			assert.Equal(test, testCase.wantInput, gotInput)
		})
	}
}
