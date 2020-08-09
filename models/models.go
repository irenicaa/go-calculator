package models

// Token ...
type Token struct {
	Kind  TokenKind
	Value string
}

// CommandKind ...
type CommandKind int

// ...
const (
	PushNumberCommand CommandKind = iota
	PushVariableCommand
	CallFunctionCommand
)

// Command ...
type Command struct {
	Kind    CommandKind
	Operand string
}

// Function ...
type Function struct {
	Arity   int
	Handler func(arguments []float64) float64
}
