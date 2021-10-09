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
	for {
		input, err := bufStdin.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			printError(err)
			continue
		}

		number, err := calculator.ProcessLine(
			input,
			calculator.BuiltInVariables,
			calculator.BuiltInFunctions,
		)
		if err != nil {
			if err != calculator.ErrNoCode {
				printError(err)
			}
			continue
		}

		fmt.Println(number)
	}
}
