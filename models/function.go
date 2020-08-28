package models

// Function ...
type Function struct {
	Arity   int // argument count
	Handler func(arguments []float64) float64
}
