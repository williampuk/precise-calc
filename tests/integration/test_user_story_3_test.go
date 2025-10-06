package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestUserStory3_OperatorPrecedence validates operator precedence rules
// User Story 3: Calculations follow standard mathematical precedence
func TestUserStory3_OperatorPrecedence(t *testing.T) {
	expression := "2 + 3 * 4 - 10 / 2"
	expectedOutput := "9"

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

// TestUserStory3_PrecedenceScenarios validates various precedence scenarios
func TestUserStory3_PrecedenceScenarios(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "multiplication before addition",
			input:    "2 + 3 * 4",
			expected: "14",
		},
		{
			name:     "division before subtraction",
			input:    "10 - 8 / 2",
			expected: "6",
		},
		{
			name:     "mixed operations with precedence",
			input:    "5 + 2 * 3 - 4 / 2",
			expected: "9",
		},
		{
			name:     "left-to-right for same precedence",
			input:    "8 / 2 * 3",
			expected: "12",
		},
		{
			name:     "complex expression",
			input:    "10 + 5 * 2 - 3 * 4 / 2",
			expected: "14",
		},
		{
			name:     "only multiplication and division",
			input:    "8 * 4 / 2 * 3",
			expected: "48",
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
