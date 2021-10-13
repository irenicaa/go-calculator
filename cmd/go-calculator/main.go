package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/irenicaa/go-calculator"
)

func printError(err error) {
	fmt.Printf("error: %s\n", err)
}

func main() {
	flag.Parse()

	bufStdin := bufio.NewReader(os.Stdin)
	interpreter := calculator.NewInterpreter(
		calculator.BuiltInVariables,
		calculator.BuiltInFunctions,
	)
	for {
		input, err := bufStdin.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			printError(err)
			continue
		}

		number, err := interpreter.Interpret(input)
		if err != nil {
			if err != calculator.ErrNoCode {
				printError(err)
			}
			continue
		}

		fmt.Println(number)
	}
}
