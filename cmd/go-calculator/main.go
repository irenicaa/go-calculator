package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	calculatorpkg "github.com/irenicaa/go-calculator"
	"github.com/irenicaa/go-calculator/tokenizer"
)

func printError(err error) {
	fmt.Printf("error: %s\n", err)
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

		variable, input := tokenizer.ExtractVariable(input)
		if variable == "" && strings.TrimSpace(input) == "" {
			continue
		}

		calculator := calculatorpkg.NewCalculator(
			calculatorpkg.BuiltInVariables,
			calculatorpkg.BuiltInFunctions,
		)
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
			calculatorpkg.BuiltInVariables[variable] = number
		}

		fmt.Println(number)
	}
}
