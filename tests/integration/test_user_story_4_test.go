package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestUserStory4_ExponentialNotation validates exponential notation support
// User Story 4: Developer can use scientific notation for large and small numbers
func TestUserStory4_ExponentialNotation(t *testing.T) {
	expression := "1E3 + 2.5e-2 * 100"
	expectedOutput := "1002.5"

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

// TestUserStory4_ExponentialScenarios validates various exponential notation scenarios
func TestUserStory4_ExponentialScenarios(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "uppercase E notation",
			input:    "1E3 + 5",
			expected: "1005",
		},
		{
			name:     "lowercase e notation",
			input:    "2.5e-2 + 1",
			expected: "1.025",
		},
		{
			name:     "large positive exponent",
			input:    "5e6 * 2",
			expected: "10000000",
		},
		{
			name:     "large negative exponent",
			input:    "1e-6 + 0.000001",
			expected: "0.000002",
		},
		{
			name:     "mixed exponential and decimal",
			input:    "1E2 + 50.5",
			expected: "150.5",
		},
		{
			name:     "exponential in complex expression",
			input:    "1E3 + 2.5e-2 * 100 - 5e2",
			expected: "502.5",
		},
		{
			name:     "very small exponential",
			input:    "1e-10 + 1e-10",
			expected: "2.0000000000e-10",
		},
		{
			name:     "scientific notation with hex",
			input:    "1E2 + 0xff",
			expected: "355",
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
