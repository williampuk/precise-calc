package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestUserStory2_MixedFormats validates mixed decimal and hex number calculations
// User Story 2: Developer wants to combine different number formats
func TestUserStory2_MixedFormats(t *testing.T) {
	expression := "0xab91 + 100.5 - 0xff"
	expectedOutput := "43766.5"

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

// TestUserStory2_AdditionalScenarios validates additional mixed format scenarios
func TestUserStory2_AdditionalScenarios(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hex multiplication with decimal",
			input:    "0xff * 1.5",
			expected: "382.5",
		},
		{
			name:     "decimal division by hex",
			input:    "1000 / 0x10",
			expected: "62.5",
		},
		{
			name:     "negative hex with decimal",
			input:    "-0xab + 171.5",
			expected: "0.5",
		},
		{
			name:     "large hex with small decimal",
			input:    "0xffff + 0.001",
			expected: "65535.001",
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
