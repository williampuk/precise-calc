package contract

import (
	"os/exec"
	"strings"
	"testing"
)

// TestCLIContractErrors validates error handling
func TestCLIContractErrors(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantCode   int
	}{
		{
			name:       "invalid syntax",
			expression: "1 + ",
			wantCode:   1,
		},
		{
			name:       "division by zero",
			expression: "10 / 0",
			wantCode:   1,
		},
		{
			name:       "invalid characters",
			expression: "1 + abc",
			wantCode:   1,
		},
		{
			name:       "invalid hex format",
			expression: "0xgg + 1",
			wantCode:   1,
		},
		{
			name:       "empty expression",
			expression: "",
			wantCode:   1,
		},
		{
			name:       "just operator",
			expression: "+",
			wantCode:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("../../calculator", tt.expression)
			output, err := cmd.CombinedOutput()

			exitError, ok := err.(*exec.ExitError)
			if !ok || exitError.ExitCode() != tt.wantCode {
				t.Errorf("Expected exit code %d, got error: %v (output: %s)", tt.wantCode, err, string(output))
			}
			// Check stderr has error message
			if ok && exitError.ExitCode() == 1 && len(strings.TrimSpace(string(output))) == 0 {
				t.Errorf("Expected error message on stderr, but got no output")
			}
		})
	}
}

// TestCLIContractInputFormatsErrors validates error input format handling
func TestCLIContractInputFormatsErrors(t *testing.T) {
	tests := []struct {
		name       string
		expression string
	}{
		{"invalid characters", "1 + z"},
		{"invalid hex", "0xgg"},
		{"malformed exponential", "1Ee3"},
		{"incomplete exponential", "1E"},
		{"invalid operator sequence", "1 + * 2"},
		{"unmatched parentheses", "(1 + 2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("../../calculator", tt.expression)
			output, err := cmd.CombinedOutput()

			exitError, ok := err.(*exec.ExitError)
			if !ok || exitError.ExitCode() != 1 {
				t.Errorf("Expected exit code 1 for invalid input, got error: %v (output: %s)", err, string(output))
			}
			// Check stderr has error message
			if ok && exitError.ExitCode() == 1 && len(strings.TrimSpace(string(output))) == 0 {
				t.Errorf("Expected error message on stderr, but got no output")
			}
		})
	}
}
