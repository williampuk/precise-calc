package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"precise-calc/pkg/calculator"
)

var writer = bufio.NewWriter(os.Stdout)

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
	ignoreInterrupt()
	defer resetInterrupt()

	scanner := bufio.NewScanner(os.Stdin)
	writer.Flush()

	fmt.Println("Interactive mode. Press Ctrl+C or Ctrl+D to exit.")
	writer.Flush()

	for {
		fmt.Print("> ")
		writer.Flush()
		if !scanner.Scan() {
			return
		}

		line := strings.TrimSpace(scanner.Text())

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
		writer.Flush()
	}
}

var originalTerm *os.File

func ignoreInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()
}

func resetInterrupt() {
	signal.Reset(os.Interrupt, syscall.SIGINT)
}

var validDecimal = regexp.MustCompile(`^[0-9]*\.?[0-9]+(?:[Ee][+-]?[0-9]+)?$`)

func init() {
	_ = validDecimal
}
