package calculator

import (
	"testing"

	"github.com/irenicaa/go-calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestInterpreter(test *testing.T) {
	type fields struct {
		variables models.VariableGroup
		functions models.FunctionGroup
	}
	type args struct {
		input string
	}

	testsCases := []struct {
		name          string
		fields        fields
		args          args
		wantVariables models.VariableGroup
		wantNumber    float64
		wantErr       string
	}{
		{
			name: "success with numbers",
			fields: fields{
				variables: models.VariableGroup{},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:          args{input: "2 + 3"},
			wantVariables: models.VariableGroup{},
			wantNumber:    5,
			wantErr:       "",
		},
		{
			name: "success with the use of variables",
			fields: fields{
				variables: models.VariableGroup{"x": 2, "y": 3},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:          args{input: "x + y"},
			wantVariables: models.VariableGroup{"x": 2, "y": 3},
			wantNumber:    5,
			wantErr:       "",
		},
		{
			name: "success with the definition of variables",
			fields: fields{
				variables: models.VariableGroup{"x": 2, "y": 3},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:          args{input: "z = x + y"},
			wantVariables: models.VariableGroup{"x": 2, "y": 3, "z": 5},
			wantNumber:    5,
			wantErr:       "",
		},
		{
			name: "success with the comment",
			fields: fields{
				variables: models.VariableGroup{},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:          args{input: "2 + 3 // test"},
			wantVariables: models.VariableGroup{},
			wantNumber:    5,
			wantErr:       "",
		},

		// errors
		{
			name: "error with the empty input",
			fields: fields{
				variables: models.VariableGroup{},
				functions: nil,
			},
			args:          args{input: ""},
			wantVariables: models.VariableGroup{},
			wantNumber:    0,
			wantErr:       ErrNoCode.Error(),
		},
		{
			name: "error with the comment only",
			fields: fields{
				variables: models.VariableGroup{},
				functions: nil,
			},
			args:          args{input: "// test"},
			wantVariables: models.VariableGroup{},
			wantNumber:    0,
			wantErr:       ErrNoCode.Error(),
		},
		{
			name: "error with calculation",
			fields: fields{
				variables: models.VariableGroup{},
				functions: nil,
			},
			args:          args{input: "2 @ 3"},
			wantVariables: models.VariableGroup{},
			wantNumber:    0,
			wantErr: "unable to calculate the code: " +
				"unable to tokenize the code: " +
				"unknown symbol '@' at position 2",
		},
		{
			name: "error with finalizing of calculation",
			fields: fields{
				variables: models.VariableGroup{},
				functions: nil,
			},
			args:          args{input: "2 + ."},
			wantVariables: models.VariableGroup{},
			wantNumber:    0,
			wantErr: "unable to finalize the calculator: " +
				"unable to finalize the tokenizer: " +
				"both integer and fractional parts are empty at EOI",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			copyOfVariables := testCase.fields.variables.Copy()

			interpreter := NewInterpreter(
				testCase.fields.variables,
				testCase.fields.functions,
			)
			gotNumber, gotErr := interpreter.Interpret(testCase.args.input)

			assert.Equal(test, copyOfVariables, testCase.fields.variables)
			assert.Equal(test, testCase.wantVariables, interpreter.Variables())
			assert.Equal(test, testCase.wantNumber, gotNumber)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}

func TestInterpreter_withSequentialCalls(test *testing.T) {
	type fields struct {
		variables models.VariableGroup
		functions models.FunctionGroup
	}
	type args struct {
		inputs []string
	}

	testsCases := []struct {
		name          string
		fields        fields
		args          args
		wantVariables models.VariableGroup
		wantNumber    float64
	}{
		{
			name: "success",
			fields: fields{
				variables: models.VariableGroup{},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args: args{
				inputs: []string{"x = 5 + 12", "y = x + 23", "z = y + 42"},
			},
			wantVariables: models.VariableGroup{"x": 17, "y": 40, "z": 82},
			wantNumber:    82,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			copyOfVariables := testCase.fields.variables.Copy()
			gotNumber, gotErr := 0.0, error(nil)

			interpreter := NewInterpreter(
				testCase.fields.variables,
				testCase.fields.functions,
			)
			for _, input := range testCase.args.inputs {
				gotNumber, gotErr = interpreter.Interpret(input)
				if gotErr != nil {
					break
				}
			}

			assert.Equal(test, copyOfVariables, testCase.fields.variables)
			assert.Equal(test, testCase.wantVariables, interpreter.Variables())
			assert.Equal(test, testCase.wantNumber, gotNumber)
			assert.NoError(test, gotErr)
		})
	}
}
