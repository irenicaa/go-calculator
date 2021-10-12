package calculator

import (
	"math"

	"github.com/irenicaa/go-calculator/models"
)

// ...
var (
	BuiltInVariables = map[string]float64{"pi": math.Pi, "e": math.E}
	BuiltInFunctions = models.FunctionGroup{
		// operators
		"+": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return arguments[0] + arguments[1], nil
			},
		},
		"-": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return arguments[0] - arguments[1], nil
			},
		},
		"*": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return arguments[0] * arguments[1], nil
			},
		},
		"/": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return arguments[0] / arguments[1], nil
			},
		},
		"%": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return math.Mod(arguments[0], arguments[1]), nil
			},
		},
		"^": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return math.Pow(arguments[0], arguments[1]), nil
			},
		},

		// functions
		"floor": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Floor(arguments[0]), nil
			},
		},
		"ceil": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Ceil(arguments[0]), nil
			},
		},
		"trunc": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Trunc(arguments[0]), nil
			},
		},
		"round": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Round(arguments[0]), nil
			},
		},
		"sin": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Sin(arguments[0]), nil
			},
		},
		"cos": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Cos(arguments[0]), nil
			},
		},
		"tan": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Tan(arguments[0]), nil
			},
		},
		"asin": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Asin(arguments[0]), nil
			},
		},
		"acos": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Acos(arguments[0]), nil
			},
		},
		"atan": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Atan(arguments[0]), nil
			},
		},
		"atan2": {
			Arity: 2,
			Handler: func(arguments []float64) (float64, error) {
				return math.Atan2(arguments[0], arguments[1]), nil
			},
		},
		"sqrt": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Sqrt(arguments[0]), nil
			},
		},
		"exp": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Exp(arguments[0]), nil
			},
		},
		"log": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Log(arguments[0]), nil
			},
		},
		"log10": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Log10(arguments[0]), nil
			},
		},
		"abs": {
			Arity: 1,
			Handler: func(arguments []float64) (float64, error) {
				return math.Abs(arguments[0]), nil
			},
		},
	}
)
