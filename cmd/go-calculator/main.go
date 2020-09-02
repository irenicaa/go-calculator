package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	calculator "github.com/irenicaa/go-calculator"
)

func main() {
	bufStdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, err := bufStdin.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}

		calculator := calculator.NewCalculator(nil, nil)
		err = calculator.Calculate(input)
		if err != nil {
			log.Print(err)
			continue
		}

		number, err := calculator.Finalize()
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Println(number)
	}
}
