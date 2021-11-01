package calculator

import (
	"math"
	"testing"

	"github.com/irenicaa/go-calculator/v2/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculator(test *testing.T) {
	type fields struct {
		variables models.VariableGroup
		functions models.FunctionGroup
	}
	type args struct {
		code string
	}

	testsCases := []struct {
		name       string
		fields     fields
		args       args
		wantNumber float64
		wantErr    string
	}{
		{
			name: "success with numbers",
			fields: fields{
				variables: nil,
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:       args{code: "2 + 3"},
			wantNumber: 5,
			wantErr:    "",
		},
		{
			name: "success with variables",
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
			args:       args{code: "x + y"},
			wantNumber: 5,
			wantErr:    "",
		},
		{
			name: "success with function calls",
			fields: fields{
				variables: nil,
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},

					"floor": {
						Arity: 1,
						Handler: func(arguments []float64) (float64, error) {
							return math.Floor(arguments[0]), nil
						},
					},
					"ceil": {
						Arity: 1,
						Handler: func(arguments []float64) (float64, error) {
							return math.Ceil(arguments[0]), nil
						},
					},
				},
			},
			args:       args{code: "floor(2.3) + ceil(2.3)"},
			wantNumber: 5,
			wantErr:    "",
		},

		// errors
		{
			name: "error with tokenization",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "2 @ 3"},
			wantNumber: 0,
			wantErr:    "unable to tokenize the code: unknown symbol '@' at position 2",
		},
		{
			name: "error with translation",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "2 + 3)"},
			wantNumber: 0,
			wantErr: "unable to translate the tokens: " +
				"missed pair for token {Kind:9 Value:)} with number #3",
		},
		{
			name: "error with evaluation",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "x + 3"},
			wantNumber: 0,
			wantErr: "unable to evaluate the commands: " +
				"unknown variable in command {Kind:1 Operand:x} with number #0",
		},
		{
			name: "error with finalizing of tokenization",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "2 + ."},
			wantNumber: 0,
			wantErr: "unable to finalize the tokenizer: " +
				"both integer and fractional parts are empty at EOI",
		},
		{
			name: "error with finalizing of translation",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "(2 + 3"},
			wantNumber: 0,
			wantErr: "unable to finalize the translator: " +
				"missed pair for token {Kind:8 Value:(}",
		},
		{
			name: "error with evaluation during finalizing",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: "2 + x"},
			wantNumber: 0,
			wantErr: "unable to evaluate the commands: " +
				"unknown variable in command {Kind:1 Operand:x} with number #0",
		},
		{
			name: "error with finalizing of evaluation",
			fields: fields{
				variables: nil,
				functions: nil,
			},
			args:       args{code: ""},
			wantNumber: 0,
			wantErr:    "unable to finalize the evaluator: number stack is empty",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotNumber := 0.0

			calculator := NewCalculator(
				testCase.fields.variables,
				testCase.fields.functions,
			)
			gotErr := calculator.Calculate(testCase.args.code)
			if gotErr == nil {
				gotNumber, gotErr = calculator.Finalize()
			}

			assert.Equal(test, testCase.wantNumber, gotNumber)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}

func TestCalculator_withSequentialCalls(test *testing.T) {
	type fields struct {
		variables models.VariableGroup
		functions models.FunctionGroup
	}
	type args struct {
		codeParts []string
	}

	testsCases := []struct {
		name       string
		fields     fields
		args       args
		wantNumber float64
	}{
		{
			name: "success",
			fields: fields{
				variables: models.VariableGroup{"number2": 2, "number3": 3},
				functions: models.FunctionGroup{
					"+": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] + arguments[1], nil
						},
					},
				},
			},
			args:       args{codeParts: []string{"(number", "2", "+", "number", "3)"}},
			wantNumber: 5,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotNumber, gotErr := 0.0, error(nil)

			calculator := NewCalculator(
				testCase.fields.variables,
				testCase.fields.functions,
			)
			for _, codePart := range testCase.args.codeParts {
				gotErr = calculator.Calculate(codePart)
				if gotErr != nil {
					break
				}
			}
			if gotErr == nil {
				gotNumber, gotErr = calculator.Finalize()
			}

			assert.Equal(test, testCase.wantNumber, gotNumber)
			assert.NoError(test, gotErr)
		})
	}
}
