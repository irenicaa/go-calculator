package models

// Function ...
type Function struct {
	Arity   int // argument count
	Handler func(arguments []float64) float64
}

// FunctionNameGroup ...
type FunctionNameGroup map[string]struct{}

// FunctionGroup ...
type FunctionGroup map[string]Function

// Names ...
func (functions FunctionGroup) Names() FunctionNameGroup {
	functionsNames := FunctionNameGroup{}
	for name := range functions {
		functionsNames[name] = struct{}{}
	}

	return functionsNames
}
