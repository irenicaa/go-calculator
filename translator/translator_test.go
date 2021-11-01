package translator

import (
	"testing"

	"github.com/irenicaa/go-calculator/v2/models"
	"github.com/stretchr/testify/assert"
)

func TestTranslator(test *testing.T) {
	type args struct {
		tokens    []models.Token
		functions models.FunctionNameGroup
	}

	testsCases := []struct {
		name         string
		args         args
		wantCommands []models.Command
		wantErr      string
	}{
		{
			name: "number",
			args: args{
				tokens:    []models.Token{{Kind: models.NumberToken, Value: "23"}},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "23"},
			},
			wantErr: "",
		},
		{
			name: "identifier",
			args: args{
				tokens:    []models.Token{{Kind: models.IdentifierToken, Value: "test"}},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushVariableCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "few operators with the same precedence",
			args: args{
				tokens: []models.Token{
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.MinusToken, Value: "-"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "-"},
			},
			wantErr: "",
		},
		{
			name: "few operators with different precedences (ascending)",
			args: args{
				tokens: []models.Token{
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "*"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
			},
			wantErr: "",
		},
		{
			name: "few operators with different precedences (descending)",
			args: args{
				tokens: []models.Token{
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "*"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
			},
			wantErr: "",
		},
		{
			name: "few operators with one pair of parentheses",
			args: args{
				tokens: []models.Token{
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.RightParenthesisToken, Value: ")"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "*"},
			},
			wantErr: "",
		},
		{
			name: "few operators with few pairs of parentheses",
			args: args{
				tokens: []models.Token{
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "5"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.RightParenthesisToken, Value: ")"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.RightParenthesisToken, Value: ")"},
					{Kind: models.ExponentiationToken, Value: "^"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "5"},
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "*"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "^"},
			},
			wantErr: "",
		},
		{
			name: "function call without arguments",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with one simple argument",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with one complex argument",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "42"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with few simple arguments",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.CommaToken, Value: ","},
					{Kind: models.NumberToken, Value: "42"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with few complex arguments",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "5"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.CommaToken, Value: ","},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.MinusToken, Value: "-"},
					{Kind: models.NumberToken, Value: "42"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "5"},
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "-"},
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},

		// errors
		{
			name: "missed left parenthesis",
			args: args{
				tokens: []models.Token{
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.RightParenthesisToken, Value: ")"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:9 Value:)} with number #3",
		},
		{
			name: "missed left parenthesis in a function call",
			args: args{
				tokens: []models.Token{
					{Kind: models.IdentifierToken, Value: "test"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.CommaToken, Value: ","},
					{Kind: models.NumberToken, Value: "42"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:9 Value:)} with number #4",
		},
		{
			name: "missed function name and left parenthesis in a function call",
			args: args{
				tokens: []models.Token{
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.CommaToken, Value: ","},
					{Kind: models.NumberToken, Value: "42"},
					{Kind: models.RightParenthesisToken, Value: ")"},
				},
				functions: nil,
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:10 Value:,} with number #1",
		},
		{
			name: "missed right parenthesis",
			args: args{
				tokens: []models.Token{
					{Kind: models.LeftParenthesisToken, Value: "("},
					{Kind: models.NumberToken, Value: "12"},
					{Kind: models.PlusToken, Value: "+"},
					{Kind: models.NumberToken, Value: "23"},
					{Kind: models.AsteriskToken, Value: "*"},
					{Kind: models.NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.PushNumberCommand, Operand: "42"},
			},
			wantErr: "missed pair for token {Kind:8 Value:(}",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotCommands := []models.Command(nil)

			translator := Translator{}
			commands, gotErr := translator.Translate(
				testCase.args.tokens,
				testCase.args.functions,
			)
			gotCommands = append(gotCommands, commands...)

			if gotErr == nil {
				commands, gotErr = translator.Finalize()
				gotCommands = append(gotCommands, commands...)
			}

			assert.Equal(test, testCase.wantCommands, gotCommands)
			if testCase.wantErr == "" {
				assert.NoError(test, gotErr)
			} else {
				assert.EqualError(test, gotErr, testCase.wantErr)
			}
		})
	}
}

func TestTranslator_withSequentialCalls(test *testing.T) {
	type args struct {
		tokenGroups [][]models.Token
		functions   models.FunctionNameGroup
	}

	testsCases := []struct {
		name         string
		args         args
		wantCommands []models.Command
	}{
		{
			name: "few operators with one pair of parentheses",
			args: args{
				tokenGroups: [][]models.Token{
					{
						{Kind: models.LeftParenthesisToken, Value: "("},
						{Kind: models.NumberToken, Value: "12"},
						{Kind: models.PlusToken, Value: "+"},
						{Kind: models.NumberToken, Value: "23"},
					},
					{
						{Kind: models.RightParenthesisToken, Value: ")"},
						{Kind: models.AsteriskToken, Value: "*"},
						{Kind: models.NumberToken, Value: "42"},
					},
				},
				functions: nil,
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "12"},
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "+"},
				{Kind: models.PushNumberCommand, Operand: "42"},
				{Kind: models.CallFunctionCommand, Operand: "*"},
			},
		},
		{
			name: "function call with one simple argument",
			args: args{
				tokenGroups: [][]models.Token{
					{
						{Kind: models.IdentifierToken, Value: "test"},
					},
					{
						{Kind: models.LeftParenthesisToken, Value: "("},
						{Kind: models.NumberToken, Value: "23"},
						{Kind: models.RightParenthesisToken, Value: ")"},
					},
				},
				functions: models.FunctionNameGroup{"test": {}},
			},
			wantCommands: []models.Command{
				{Kind: models.PushNumberCommand, Operand: "23"},
				{Kind: models.CallFunctionCommand, Operand: "test"},
			},
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotCommands, gotErr := []models.Command(nil), error(nil)

			translator := Translator{}
			for _, tokenGroup := range testCase.args.tokenGroups {
				commands, err := translator.Translate(
					tokenGroup,
					testCase.args.functions,
				)

				gotCommands = append(gotCommands, commands...)
				gotErr = err
				if gotErr != nil {
					break
				}
			}
			if gotErr == nil {
				commands, err := translator.Finalize()

				gotCommands = append(gotCommands, commands...)
				gotErr = err
			}

			assert.Equal(test, testCase.wantCommands, gotCommands)
			assert.NoError(test, gotErr)
		})
	}
}
