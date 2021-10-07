package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	calculator "github.com/irenicaa/go-calculator"
	"github.com/irenicaa/go-calculator/models"
	"github.com/irenicaa/go-calculator/tokenizer"
)

var variables = map[string]float64{"pi": math.Pi, "e": math.E}
var functions = models.FunctionGroup{
	// operators
	"+": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return arguments[0] + arguments[1]
		},
	},
	"-": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return arguments[0] - arguments[1]
		},
	},
	"*": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return arguments[0] * arguments[1]
		},
	},
	"/": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return arguments[0] / arguments[1]
		},
	},
	"%": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return math.Mod(arguments[0], arguments[1])
		},
	},
	"^": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return math.Pow(arguments[0], arguments[1])
		},
	},

	// functions
	"floor": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Floor(arguments[0])
		},
	},
	"ceil": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Ceil(arguments[0])
		},
	},
	"trunc": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Trunc(arguments[0])
		},
	},
	"round": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Round(arguments[0])
		},
	},
	"sin": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Sin(arguments[0])
		},
	},
	"cos": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Cos(arguments[0])
		},
	},
	"tan": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Tan(arguments[0])
		},
	},
	"asin": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Asin(arguments[0])
		},
	},
	"acos": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Acos(arguments[0])
		},
	},
	"atan": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Atan(arguments[0])
		},
	},
	"atan2": {
		Arity: 2,
		Handler: func(arguments []float64) float64 {
			return math.Atan2(arguments[0], arguments[1])
		},
	},
	"sqrt": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Sqrt(arguments[0])
		},
	},
	"exp": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Exp(arguments[0])
		},
	},
	"log": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Log(arguments[0])
		},
	},
	"log10": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Log10(arguments[0])
		},
	},
	"abs": {
		Arity: 1,
		Handler: func(arguments []float64) float64 {
			return math.Abs(arguments[0])
		},
	},
}

func printError(err error) {
	fmt.Printf("error: %s\n", err)
}

func extractVariable(input string) (string, string) {
	variable := ""
	if separatorIndex := strings.IndexRune(input, '='); separatorIndex != -1 {
		variable = strings.TrimSpace(input[:separatorIndex])
		input = input[separatorIndex+1:]
	}

	return variable, input
}

func main() {
	flag.Parse()

	bufStdin := bufio.NewReader(os.Stdin)
	for {
		input, err := bufStdin.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			printError(err)
			continue
		}

		input = tokenizer.RemoveComment(input)

		variable, input := extractVariable(input)
		if variable == "" && strings.TrimSpace(input) == "" {
			continue
		}

		calculator := calculator.NewCalculator(variables, functions)
		if err = calculator.Calculate(input); err != nil {
			printError(err)
			continue
		}

		number, err := calculator.Finalize()
		if err != nil {
			printError(err)
			continue
		}

		if variable != "" {
			variables[variable] = number
		}

		fmt.Println(number)
	}
}
