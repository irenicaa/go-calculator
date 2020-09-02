package main

import (
	"bufio"
	"fmt"
	"os"

	calculator "github.com/irenicaa/go-calculator"
)

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, err := bufStdin.ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		calculator := calculator.NewCalculator(nil, nil)
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
