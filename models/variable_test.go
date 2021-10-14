package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableGroup_Copy(test *testing.T) {
	testsCases := []struct {
		name      string
		variables VariableGroup
		want      VariableGroup
	}{
		{
			name:      "nonempty",
			variables: VariableGroup{"one": 23, "two": 42},
			want:      VariableGroup{"one": 23, "two": 42},
		},
		{
			name:      "empty",
			variables: VariableGroup{},
			want:      VariableGroup{},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			got := testCase.variables.Copy()

			assert.Equal(test, testCase.want, got)
		})
	}
}

func TestVariableGroup_Copy_withModifiedCopy(test *testing.T) {
	variables := VariableGroup{"one": 5, "two": 12}

	copyOfVariables := variables.Copy()
	copyOfVariables["two"] = 23
	copyOfVariables["three"] = 42

	assert.Equal(test, VariableGroup{"one": 5, "two": 12}, variables)
	assert.Equal(
		test,
		VariableGroup{"one": 5, "two": 23, "three": 42},
		copyOfVariables,
	)
}
