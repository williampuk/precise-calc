package main

import (
	calculator "calc/pkg/calc"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/peterh/liner"
)

const MaxInputLength = 4096

type InputLengthError struct {
	Message   string
	MaxLength int
	ActualLen int
}

func (e *InputLengthError) Error() string {
	return e.Message
}

func NewLengthError(actualLen, maxLen int) *InputLengthError {
	return &InputLengthError{
		Message:   fmt.Sprintf("expression exceeds maximum length of %d characters (got %d)", maxLen, actualLen),
		MaxLength: maxLen,
		ActualLen: actualLen,
	}
}

func ValidateInput(line string) (string, *InputLengthError) {
	if len(line) > MaxInputLength {
		return "", NewLengthError(len(line), MaxInputLength)
	}
	return line, nil
}

func main() {
	if len(os.Args) < 2 {
		runInteractiveMode()
		return
	}

	expression := os.Args[1]

	_, lengthErr := ValidateInput(expression)
	if lengthErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", lengthErr.Error())
		os.Exit(1)
	}

	result := calculator.EvaluateString(expression)
	if result.IsError() {
		fmt.Fprintf(os.Stderr, "Error: %s\n", result.Error.Error())
		os.Exit(1)
	}
	fmt.Println(result.Result.String())
}

func runInteractiveMode() {
	state := liner.NewLiner()
	defer state.Close()

	fmt.Println("Interactive mode. Press Ctrl+C or Ctrl+D to exit.")

	for {
		line, err := state.Prompt("> ")
		if err == io.EOF {
			fmt.Println()
			return
		}
		if err != nil {
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "exit" || line == "quit" {
			return
		}

		_, lengthErr := ValidateInput(line)
		if lengthErr != nil {
			fmt.Fprintf(os.Stderr, "\033[31mError: %s\033[0m\n", lengthErr.Error())
			continue
		}

		result := calculator.EvaluateString(line)
		if result.IsError() {
			fmt.Fprintf(os.Stderr, "\033[31mError: %s\033[0m\n", result.Error.Error())
			continue
		}

		fmt.Printf("= %s\n", result.Result.String())
		state.AppendHistory(line)
	}
}
