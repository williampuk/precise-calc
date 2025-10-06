package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestUserStory1_PreciseDecimals validates precise decimal calculations
// User Story 1: Developer needs exact precision with very small decimals
func TestUserStory1_PreciseDecimals(t *testing.T) {
	expression := "0.0000000000000001 + 0.1 + -99999999999999"
	expectedOutput := "-99999999999998.8999999999999999"

	cmd := exec.Command("../../calculator", expression)
	output, err := cmd.CombinedOutput()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != 0 {
			t.Errorf("Expected exit code 0, got error: %v (output: %s)", err, string(output))
		}
		return
	}

	// Success case: err == nil means exit code 0
	actualOutput := strings.TrimSpace(string(output))
	if actualOutput != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, actualOutput)
	}
}

// TestUserStory1_AdditionalScenarios validates additional precise decimal scenarios
func TestUserStory1_AdditionalScenarios(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "very small decimals",
			input:    "0.0000000000000001 + 0.0000000000000002",
			expected: "3.0000000000e-16",
		},
		{
			name:     "large precision preservation",
			input:    "1.0000000000000001 * 2",
			expected: "2.0000000000000002",
		},
		{
			name:     "decimal subtraction precision",
			input:    "1.0000000000000001 - 0.0000000000000001",
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("../../calculator", tt.input)
			output, err := cmd.CombinedOutput()

			if err != nil {
				if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != 0 {
					t.Errorf("Expected exit code 0, got error: %v (output: %s)", err, string(output))
				}
				return
			}

			// Success case: err == nil means exit code 0
			actualOutput := strings.TrimSpace(string(output))
			if actualOutput != tt.expected {
				t.Errorf("Expected output %q, got %q", tt.expected, actualOutput)
			}
		})
	}
}
