package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(test *testing.T) {
	type args struct {
		tokens    []Token
		functions map[string]struct{}
	}

	testsCases := []struct {
		name         string
		args         args
		wantCommands []Command
		wantErr      string
	}{
		{
			name: "number",
			args: args{
				tokens:    []Token{{Kind: NumberToken, Value: "23"}},
				functions: nil,
			},
			wantCommands: []Command{{Kind: PushNumberCommand, Operand: "23"}},
			wantErr:      "",
		},
		{
			name: "identifier",
			args: args{
				tokens:    []Token{{Kind: IdentifierToken, Value: "test"}},
				functions: nil,
			},
			wantCommands: []Command{{Kind: PushVariableCommand, Operand: "test"}},
			wantErr:      "",
		},
		{
			name: "few operators with the same precedence",
			args: args{
				tokens: []Token{
					{Kind: NumberToken, Value: "12"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "23"},
					{Kind: MinusToken, Value: "-"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "-"},
			},
			wantErr: "",
		},
		{
			name: "few operators with different precedences (ascending)",
			args: args{
				tokens: []Token{
					{Kind: NumberToken, Value: "12"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "23"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "*"},
				{Kind: CallFunctionCommand, Operand: "+"},
			},
			wantErr: "",
		},
		{
			name: "few operators with different precedences (descending)",
			args: args{
				tokens: []Token{
					{Kind: NumberToken, Value: "12"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "23"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "*"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "+"},
			},
			wantErr: "",
		},
		{
			name: "few operators with one pair of parentheses",
			args: args{
				tokens: []Token{
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "12"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "23"},
					{Kind: RightParenthesisToken, Value: ")"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "*"},
			},
			wantErr: "",
		},
		{
			name: "few operators with few pairs of parentheses",
			args: args{
				tokens: []Token{
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "5"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "12"},
					{Kind: RightParenthesisToken, Value: ")"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "23"},
					{Kind: RightParenthesisToken, Value: ")"},
					{Kind: ExponentiationToken, Value: "^"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "5"},
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "*"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "^"},
			},
			wantErr: "",
		},
		{
			name: "function call without arguments",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: []Command{{Kind: CallFunctionCommand, Operand: "test"}},
			wantErr:      "",
		},
		{
			name: "function call with one simple argument",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "23"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with one complex argument",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "23"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "42"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with few simple arguments",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "23"},
					{Kind: CommaToken, Value: ","},
					{Kind: NumberToken, Value: "42"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},
		{
			name: "function call with few complex arguments",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "5"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "12"},
					{Kind: CommaToken, Value: ","},
					{Kind: NumberToken, Value: "23"},
					{Kind: MinusToken, Value: "-"},
					{Kind: NumberToken, Value: "42"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "5"},
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "-"},
				{Kind: CallFunctionCommand, Operand: "test"},
			},
			wantErr: "",
		},

		// errors
		{
			name: "missed left parenthesis",
			args: args{
				tokens: []Token{
					{Kind: NumberToken, Value: "12"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "23"},
					{Kind: RightParenthesisToken, Value: ")"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:9 Value:)} with number #3",
		},
		{
			name: "missed left parenthesis in a function call",
			args: args{
				tokens: []Token{
					{Kind: IdentifierToken, Value: "test"},
					{Kind: NumberToken, Value: "23"},
					{Kind: CommaToken, Value: ","},
					{Kind: NumberToken, Value: "42"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: map[string]struct{}{"test": {}},
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:9 Value:)} with number #4",
		},
		{
			name: "missed function name and left parenthesis in a function call",
			args: args{
				tokens: []Token{
					{Kind: NumberToken, Value: "23"},
					{Kind: CommaToken, Value: ","},
					{Kind: NumberToken, Value: "42"},
					{Kind: RightParenthesisToken, Value: ")"},
				},
				functions: nil,
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:10 Value:,} with number #1",
		},
		{
			name: "missed right parenthesis",
			args: args{
				tokens: []Token{
					{Kind: LeftParenthesisToken, Value: "("},
					{Kind: NumberToken, Value: "12"},
					{Kind: PlusToken, Value: "+"},
					{Kind: NumberToken, Value: "23"},
					{Kind: AsteriskToken, Value: "*"},
					{Kind: NumberToken, Value: "42"},
				},
				functions: nil,
			},
			wantCommands: nil,
			wantErr:      "missed pair for token {Kind:8 Value:(}",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotCommands := []Command(nil)

			translator := Translator{}
			gotErr := translator.Translate(testCase.args.tokens, testCase.args.functions)
			if gotErr == nil {
				gotCommands, gotErr = translator.Finalize()
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
		tokenGroups [][]Token
		functions   map[string]struct{}
	}

	testsCases := []struct {
		name         string
		args         args
		wantCommands []Command
		wantErr      string
	}{
		{
			name: "few operators with different precedences (ascending)",
			args: args{
				tokenGroups: [][]Token{
					{
						{Kind: LeftParenthesisToken, Value: "("},
						{Kind: NumberToken, Value: "12"},
						{Kind: PlusToken, Value: "+"},
						{Kind: NumberToken, Value: "23"},
					},
					{
						{Kind: RightParenthesisToken, Value: ")"},
						{Kind: AsteriskToken, Value: "*"},
						{Kind: NumberToken, Value: "42"},
					},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushNumberCommand, Operand: "12"},
				{Kind: PushNumberCommand, Operand: "23"},
				{Kind: CallFunctionCommand, Operand: "+"},
				{Kind: PushNumberCommand, Operand: "42"},
				{Kind: CallFunctionCommand, Operand: "*"},
			},
			wantErr: "",
		},
		{
			name: "function call with one simple argument",
			args: args{
				tokenGroups: [][]Token{
					{
						{Kind: IdentifierToken, Value: "test"},
					},
					{
						{Kind: LeftParenthesisToken, Value: "()"},
						{Kind: NumberToken, Value: "23"},
						{Kind: RightParenthesisToken, Value: ")"},
					},
				},
				functions: nil,
			},
			wantCommands: []Command{
				{Kind: PushVariableCommand, Operand: "test"},
				{Kind: PushNumberCommand, Operand: "23"},
			},
			wantErr: "",
		},
	}
	for _, testCase := range testsCases {
		test.Run(testCase.name, func(test *testing.T) {
			gotCommand, gotErr := []Command(nil), error(nil)
			translator := Translator{}
			for _, token := range testCase.args.tokenGroups {
				gotErr = translator.Translate(token, testCase.args.functions)
				if gotErr != nil {
					break
				}
			}
			if gotErr == nil {
				gotCommand, gotErr = translator.Finalize()
			}

			assert.Equal(test, testCase.wantCommands, gotCommand)
			assert.NoError(test, gotErr)
		})
	}
}
