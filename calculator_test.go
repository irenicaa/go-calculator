package calculator

import (
	"testing"

	"github.com/irenicaa/go-calculator/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculator(test *testing.T) {
	type fields struct {
		variables map[string]float64
		functions map[string]models.Function
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
	}{}
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
		variables map[string]float64
		functions map[string]models.Function
	}
	type args struct {
		codeParts []string
	}

	testsCases := []struct {
		name       string
		fields     fields
		args       args
		wantNumber float64
	}{}
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
