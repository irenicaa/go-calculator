package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	calculator "github.com/irenicaa/go-calculator"
	"github.com/irenicaa/go-calculator/models"
)

var functions = models.FunctionGroup{
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
}

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, err := bufStdin.ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		calculator := calculator.NewCalculator(nil, functions)
		err = calculator.Calculate(input)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		number, err := calculator.Finalize()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		fmt.Println(number)
	}
}
