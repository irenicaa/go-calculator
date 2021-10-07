package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuiltInVariables(test *testing.T) {
	type args struct {
		name string
	}

	testsCases := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{
			name: "pi",
			args: args{
				name: "pi",
			},
			wantResult: math.Pi,
		},
		{
			name: "e",
			args: args{
				name: "e",
			},
			wantResult: math.E,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotResult := BuiltInVariables[testCase.args.name]

			assert.InDelta(test, testCase.wantResult, gotResult, 1e-6)
		})
	}
}

func TestBuiltInFunctions(test *testing.T) {
	type args struct {
		name      string
		arguments []float64
	}

	testsCases := []struct {
		name       string
		args       args
		wantArity  int
		wantResult float64
		wantErr    string
	}{
		// operators
		{
			name: "+",
			args: args{
				name:      "+",
				arguments: []float64{2, 3},
			},
			wantArity:  2,
			wantResult: 5,
			wantErr:    "",
		},
		{
			name: "-",
			args: args{
				name:      "-",
				arguments: []float64{2, 3},
			},
			wantArity:  2,
			wantResult: -1,
			wantErr:    "",
		},
		{
			name: "*",
			args: args{
				name:      "*",
				arguments: []float64{2, 3},
			},
			wantArity:  2,
			wantResult: 6,
			wantErr:    "",
		},
		{
			name: "/",
			args: args{
				name:      "/",
				arguments: []float64{5, 2},
			},
			wantArity:  2,
			wantResult: 2.5,
			wantErr:    "",
		},
		{
			name: "%",
			args: args{
				name:      "%",
				arguments: []float64{5, 2},
			},
			wantArity:  2,
			wantResult: 1,
			wantErr:    "",
		},
		{
			name: "^",
			args: args{
				name:      "^",
				arguments: []float64{2, 3},
			},
			wantArity:  2,
			wantResult: 8,
			wantErr:    "",
		},

		// functions
		{
			name: "floor/positive",
			args: args{
				name:      "floor",
				arguments: []float64{2.5},
			},
			wantArity:  1,
			wantResult: 2,
			wantErr:    "",
		},
		{
			name: "floor/negative",
			args: args{
				name:      "floor",
				arguments: []float64{-2.5},
			},
			wantArity:  1,
			wantResult: -3,
			wantErr:    "",
		},
		{
			name: "ceil/positive",
			args: args{
				name:      "ceil",
				arguments: []float64{2.5},
			},
			wantArity:  1,
			wantResult: 3,
			wantErr:    "",
		},
		{
			name: "ceil/negative",
			args: args{
				name:      "ceil",
				arguments: []float64{-2.5},
			},
			wantArity:  1,
			wantResult: -2,
			wantErr:    "",
		},
		{
			name: "trunc/positive",
			args: args{
				name:      "trunc",
				arguments: []float64{2.5},
			},
			wantArity:  1,
			wantResult: 2,
			wantErr:    "",
		},
		{
			name: "trunc/negative",
			args: args{
				name:      "trunc",
				arguments: []float64{-2.5},
			},
			wantArity:  1,
			wantResult: -2,
			wantErr:    "",
		},
		{
			name: "round/positive/less than half",
			args: args{
				name:      "round",
				arguments: []float64{2.4},
			},
			wantArity:  1,
			wantResult: 2,
			wantErr:    "",
		},
		{
			name: "round/positive/greater than half",
			args: args{
				name:      "round",
				arguments: []float64{2.6},
			},
			wantArity:  1,
			wantResult: 3,
			wantErr:    "",
		},
		{
			name: "round/negative/less than half",
			args: args{
				name:      "round",
				arguments: []float64{-2.4},
			},
			wantArity:  1,
			wantResult: -2,
			wantErr:    "",
		},
		{
			name: "round/negative/greater than half",
			args: args{
				name:      "round",
				arguments: []float64{-2.6},
			},
			wantArity:  1,
			wantResult: -3,
			wantErr:    "",
		},
		{
			name: "sin",
			args: args{
				name:      "sin",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 0.909297,
			wantErr:    "",
		},
		{
			name: "cos",
			args: args{
				name:      "cos",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: -0.416147,
			wantErr:    "",
		},
		{
			name: "tan",
			args: args{
				name:      "tan",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: -2.18504,
			wantErr:    "",
		},
		{
			name: "asin",
			args: args{
				name:      "asin",
				arguments: []float64{0.2},
			},
			wantArity:  1,
			wantResult: 0.201358,
			wantErr:    "",
		},
		{
			name: "acos",
			args: args{
				name:      "acos",
				arguments: []float64{0.2},
			},
			wantArity:  1,
			wantResult: 1.369438,
			wantErr:    "",
		},
		{
			name: "atan",
			args: args{
				name:      "atan",
				arguments: []float64{0.2},
			},
			wantArity:  1,
			wantResult: 0.197396,
			wantErr:    "",
		},
		{
			name: "atan2",
			args: args{
				name:      "atan2",
				arguments: []float64{2, 3},
			},
			wantArity:  2,
			wantResult: 0.588003,
			wantErr:    "",
		},
		{
			name: "sqrt",
			args: args{
				name:      "sqrt",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 1.414214,
			wantErr:    "",
		},
		{
			name: "exp",
			args: args{
				name:      "exp",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 7.389056,
			wantErr:    "",
		},
		{
			name: "log",
			args: args{
				name:      "log",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 0.693147,
			wantErr:    "",
		},
		{
			name: "log10",
			args: args{
				name:      "log10",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 0.30103,
			wantErr:    "",
		},
		{
			name: "abs/positive",
			args: args{
				name:      "abs",
				arguments: []float64{2},
			},
			wantArity:  1,
			wantResult: 2,
			wantErr:    "",
		},
		{
			name: "abs/negative",
			args: args{
				name:      "abs",
				arguments: []float64{-2},
			},
			wantArity:  1,
			wantResult: 2,
			wantErr:    "",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotFunction, gotOk := BuiltInFunctions[testCase.args.name]
			require.True(test, gotOk)

			gotResult, gotErr := gotFunction.Handler(testCase.args.arguments)

			assert.Equal(test, testCase.wantArity, gotFunction.Arity)
			assert.InDelta(test, testCase.wantResult, gotResult, 1e-6)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}
