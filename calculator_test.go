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
		args       args
		fields     fields
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
