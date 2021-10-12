package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionGroup_Names(test *testing.T) {
	testsCases := []struct {
		name      string
		functions FunctionGroup
		want      FunctionNameGroup
	}{
		{
			name: "nonempty",
			functions: FunctionGroup{
				"add": {
					Arity: 2,
					Handler: func(arguments []float64) (float64, error) {
						return arguments[0] + arguments[1], nil
					},
				},
				"sub": {
					Arity: 2,
					Handler: func(arguments []float64) (float64, error) {
						return arguments[0] - arguments[1], nil
					},
				},
			},
			want: FunctionNameGroup{"add": {}, "sub": {}},
		},
		{
			name:      "empty",
			functions: FunctionGroup{},
			want:      FunctionNameGroup{},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			got := testCase.functions.Names()

			assert.Equal(test, testCase.want, got)
		})
	}
}
