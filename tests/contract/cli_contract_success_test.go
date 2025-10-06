package contract

import (
	"os/exec"
	"strings"
	"testing"
)

// TestCLIContractSuccess validates successful CLI execution
func TestCLIContractSuccess(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantCode   int
		wantOutput string
	}{
		{
			name:       "simple addition",
			expression: "1 + 2",
			wantCode:   0,
			wantOutput: "3",
		},
		{
			name:       "hex and decimal",
			expression: "0xa + 1.5",
			wantCode:   0,
			wantOutput: "11.5",
		},
		{
			name:       "negative numbers",
			expression: "-5 + 10",
			wantCode:   0,
			wantOutput: "5",
		},
		{
			name:       "operator precedence",
			expression: "2 + 3 * 4",
			wantCode:   0,
			wantOutput: "14",
		},
		{
			name:       "exponential notation",
			expression: "1E3 + 2",
			wantCode:   0,
			wantOutput: "1002",
		},
		{
			name:       "lowercase exponential",
			expression: "2.5e-2 + 1",
			wantCode:   0,
			wantOutput: "1.025",
		},
		{
			name:       "large exponential",
			expression: "5e6 * 2",
			wantCode:   0,
			wantOutput: "10000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("../../calculator", tt.expression)
			output, err := cmd.CombinedOutput()

			if err != nil {
				if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != tt.wantCode {
					t.Errorf("Expected exit code %d, got error: %v (output: %s)", tt.wantCode, err, string(output))
				}
				return
			}

			// Success case: err == nil means exit code 0
			if tt.wantCode != 0 {
				t.Errorf("Expected exit code %d, but command succeeded", tt.wantCode)
				return
			}

			actualOutput := strings.TrimSpace(string(output))
			if actualOutput != tt.wantOutput {
				t.Errorf("Expected output %q, got %q", tt.wantOutput, actualOutput)
			}
		})
	}
}

// TestCLIContractInputFormatsSuccess validates successful input format handling
func TestCLIContractInputFormatsSuccess(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantOutput string
	}{
		{"decimal numbers", "1.23 + 4.56", "5.79"},
		{"hex numbers", "0xff + 0x10", "271"},
		{"negative hex", "-0xab + 10", "-161"},
		{"mixed formats", "0.5 * 0xff", "127.5"},
		{"whitespace", "  1   +   2  ", "3"},
		{"tabs and newlines", "1\t+\n2", "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("../../calculator", tt.expression)
			output, err := cmd.CombinedOutput()

			if err != nil {
				if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != 0 {
					t.Errorf("Expected exit code 0, got error: %v (output: %s)", err, string(output))
				}
				return
			}

			// Success case: err == nil means exit code 0
			actualOutput := strings.TrimSpace(string(output))
			if actualOutput != tt.wantOutput {
				t.Errorf("Expected output %q, got %q", tt.wantOutput, actualOutput)
			}
		})
	}
}
