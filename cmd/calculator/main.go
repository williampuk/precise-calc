package main

import (
	"fmt"
	"os"

	"precise-calc/pkg/calculator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <expression>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s \"1 + 2 * 3\"\n", os.Args[0])
		os.Exit(1)
	}

	expression := os.Args[1]

	// Evaluate the expression
	result := calculator.EvaluateString(expression)

	if result.IsError() {
		fmt.Fprintf(os.Stderr, "Error: %s\n", result.Error.Error())
		os.Exit(1)
	}

	// Print the result to stdout
	fmt.Println(result.Result.String())
}
