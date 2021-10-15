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
		wantCode     string
	}{
		{
			name:         "string with separator",
			args:         args{"test1 = test2 + 1"},
			wantVariable: "test1",
			wantCode:     " test2 + 1",
		},
		{
			name:         "string with separator and without variable",
			args:         args{"= test2 + 1"},
			wantVariable: "",
			wantCode:     " test2 + 1",
		},
		{
			name:         "string without separator",
			args:         args{"test2 + 1"},
			wantVariable: "",
			wantCode:     "test2 + 1",
		},
		{
			name:         "string with several separators",
			args:         args{"test1 = test2 = test3 + 1"},
			wantVariable: "test1",
			wantCode:     " test2 = test3 + 1",
		},
		{
			name:         "empty string",
			args:         args{""},
			wantVariable: "",
			wantCode:     "",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotVariable, gotCode := ExtractVariable(testCase.args.input)

			assert.Equal(test, testCase.wantVariable, gotVariable)
			assert.Equal(test, testCase.wantCode, gotCode)
		})
	}
}
