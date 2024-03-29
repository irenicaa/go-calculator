package evaluator

import (
	"testing"
	"testing/iotest"

	"github.com/irenicaa/go-calculator/v2/models"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator(test *testing.T) {
	type args struct {
		commands  []models.Command
		variables models.VariableGroup
		functions models.FunctionGroup
	}

	testsCases := []struct {
		name       string
		args       args
		wantNumber float64
		wantErr    string
	}{
		{
			name: "without commands",
			args: args{
				commands:  []models.Command{},
				variables: nil,
				functions: nil,
			},
			wantNumber: 0,
			wantErr:    "number stack is empty",
		},
		{
			name: "with the push number command (success)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "2.3"},
				},
				variables: nil,
				functions: nil,
			},
			wantNumber: 2.3,
			wantErr:    "",
		},
		{
			name: "with the push number command (error)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "incorrect"},
				},
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
				commands: []models.Command{
					{Kind: models.PushVariableCommand, Operand: "test"},
				},
				variables: models.VariableGroup{"test": 2.3},
				functions: nil,
			},
			wantNumber: 2.3,
			wantErr:    "",
		},
		{
			name: "with the push variable command (error)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushVariableCommand, Operand: "unknown"},
				},
				variables: models.VariableGroup{"test": 2.3},
				functions: nil,
			},
			wantNumber: 0,
			wantErr: "unknown variable in command {Kind:1 Operand:unknown} " +
				"with number #0",
		},
		{
			name: "with the call function command (success)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "2"},
					{Kind: models.PushNumberCommand, Operand: "3"},
					{Kind: models.CallFunctionCommand, Operand: "sub"},
				},
				variables: nil,
				functions: models.FunctionGroup{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] - arguments[1], nil
						},
					},
				},
			},
			wantNumber: -1,
			wantErr:    "",
		},
		{
			name: "with the call function command (error with an unknown function)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "2"},
					{Kind: models.PushNumberCommand, Operand: "3"},
					{Kind: models.CallFunctionCommand, Operand: "unknown"},
				},
				variables: nil,
				functions: models.FunctionGroup{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] - arguments[1], nil
						},
					},
				},
			},
			wantNumber: 0,
			wantErr: "unknown function in command {Kind:2 Operand:unknown} " +
				"with number #2",
		},
		{
			name: "with the call function command (error with lack of arguments)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "2"},
					{Kind: models.CallFunctionCommand, Operand: "sub"},
				},
				variables: nil,
				functions: models.FunctionGroup{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] - arguments[1], nil
						},
					},
				},
			},
			wantNumber: 0,
			wantErr: "number stack is empty for argument #1 in command " +
				"{Kind:2 Operand:sub} with number #1",
		},
		{
			name: "with the call function command (error with the function call)",
			args: args{
				commands: []models.Command{
					{Kind: models.PushNumberCommand, Operand: "2"},
					{Kind: models.PushNumberCommand, Operand: "3"},
					{Kind: models.CallFunctionCommand, Operand: "sub"},
				},
				variables: nil,
				functions: models.FunctionGroup{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return 0, iotest.ErrTimeout
						},
					},
				},
			},
			wantNumber: 0,
			wantErr: "unable to call the function from command " +
				"{Kind:2 Operand:sub} with number #2: " + iotest.ErrTimeout.Error(),
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotNumber := 0.0

			evaluator := Evaluator{}
			gotErr := evaluator.Evaluate(
				testCase.args.commands,
				testCase.args.variables,
				testCase.args.functions,
			)
			if gotErr == nil {
				gotNumber, gotErr = evaluator.Finalize()
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

func TestEvaluator_withSequentialCalls(test *testing.T) {
	type args struct {
		commandGroups [][]models.Command
		variables     models.VariableGroup
		functions     models.FunctionGroup
	}

	testsCases := []struct {
		name       string
		args       args
		wantNumber float64
	}{
		{
			name: "with the call function command",
			args: args{
				commandGroups: [][]models.Command{
					{
						{Kind: models.PushNumberCommand, Operand: "2"},
						{Kind: models.PushNumberCommand, Operand: "3"},
					},
					{
						{Kind: models.CallFunctionCommand, Operand: "sub"},
					},
				},
				variables: nil,
				functions: models.FunctionGroup{
					"sub": {
						Arity: 2,
						Handler: func(arguments []float64) (float64, error) {
							return arguments[0] - arguments[1], nil
						},
					},
				},
			},
			wantNumber: -1,
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotNumber, gotErr := 0.0, error(nil)

			evaluator := Evaluator{}
			for _, commandGroup := range testCase.args.commandGroups {
				gotErr = evaluator.Evaluate(
					commandGroup,
					testCase.args.variables,
					testCase.args.functions,
				)
				if gotErr != nil {
					break
				}
			}
			if gotErr == nil {
				gotNumber, gotErr = evaluator.Finalize()
			}

			assert.Equal(test, testCase.wantNumber, gotNumber)
			assert.NoError(test, gotErr)
		})
	}
}
