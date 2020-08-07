package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(test *testing.T) {
	type args struct {
		commands  []Command
		variables map[string]float64
		functions map[string]Function
	}

	testsCases := []struct {
		name       string
		args       args
		wantNumber float64
		wantErr    string
	}{
		{
			name:       "empty",
			args:       args{commands: []Command{}, variables: nil, functions: nil},
			wantNumber: 0,
			wantErr:    "number stack is empty",
		},
		{
			name: "with the push number command (success)",
			args: args{
				commands:  []Command{{Kind: PushNumberCommand, Operand: "2.3"}},
				variables: nil,
				functions: nil,
			},
			wantNumber: 2.3,
			wantErr:    "",
		},
		{
			name: "with the push number command (error)",
			args: args{
				commands:  []Command{{Kind: PushNumberCommand, Operand: "incorrect"}},
				variables: nil,
				functions: nil,
			},
			wantNumber: 0,
			wantErr: "incorrect number for command {Kind:0 Operand:incorrect} " +
				"with number #0: strconv.ParseFloat: parsing \"incorrect\": " +
				"invalid syntax",
		},
		{
			name: "with the push variable command (success)",
			args: args{
				commands:  []Command{{Kind: PushVariableCommand, Operand: "test"}},
				variables: map[string]float64{"test": 2.3},
				functions: nil,
			},
			wantNumber: 2.3,
			wantErr:    "",
		},
		{
			name: "with the push variable command (error)",
			args: args{
				commands:  []Command{{Kind: PushVariableCommand, Operand: "unknown"}},
				variables: map[string]float64{"test": 2.3},
				functions: nil,
			},
			wantNumber: 0,
			wantErr: "unknown variable in command {Kind:1 Operand:unknown} " +
				"with number #0",
		},
		{
			name: "with the push call function command (success)",
			args: args{
				commands: []Command{
					{Kind: PushNumberCommand, Operand: "2"},
					{Kind: PushNumberCommand, Operand: "3"},
					{Kind: CallFunctionCommand, Operand: "sub"},
				},
				variables: nil,
				functions: map[string]Function{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) float64 {
							return arguments[0] - arguments[1]
						},
					},
				},
			},
			wantNumber: -1,
			wantErr:    "",
		},
		{
			name: "with the push call function command (error)",
			args: args{
				commands: []Command{
					{Kind: PushNumberCommand, Operand: "2"},
					{Kind: PushNumberCommand, Operand: "3"},
					{Kind: CallFunctionCommand, Operand: "unknown"},
				},
				variables: nil,
				functions: map[string]Function{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) float64 {
							return arguments[0] - arguments[1]
						},
					},
				},
			},
			wantNumber: 0,
			wantErr: "unknown function in command {Kind:2 Operand:unknown} " +
				"with number #2",
		},
		{
			name: "with the push call function command (error)",
			args: args{
				commands: []Command{
					{Kind: PushNumberCommand, Operand: "2"},
					{Kind: CallFunctionCommand, Operand: "sub"},
				},
				variables: nil,
				functions: map[string]Function{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) float64 {
							return arguments[0] - arguments[1]
						},
					},
				},
			},
			wantNumber: 0,
			wantErr: "number stack is empty for argument #1 in command " +
				"{Kind:2 Operand:sub} with number #1",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotNumber, gotErr := Evaluate(
				testCase.args.commands,
				testCase.args.variables,
				testCase.args.functions,
			)

			assert.Equal(test, testCase.wantNumber, gotNumber)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}
