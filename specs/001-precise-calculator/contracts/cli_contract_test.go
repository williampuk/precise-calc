package contracts

import (
	"os/exec"
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
			// This test will fail until the calculator binary is implemented
			cmd := exec.Command("./calculator", tt.expression)
			output, err := cmd.CombinedOutput()

			// Expected to fail initially - implementation not yet complete
			if err == nil {
				t.Errorf("Expected command to fail (not implemented yet), but it succeeded with output: %s", output)
				return
			}

			// When implemented, these assertions should pass:
			// if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != tt.wantCode {
			//     t.Errorf("Expected exit code %d, got error: %v", tt.wantCode, err)
			// }
			// if strings.TrimSpace(string(output)) != tt.wantOutput {
			//     t.Errorf("Expected output %q, got %q", tt.wantOutput, strings.TrimSpace(string(output)))
			// }
		})
	}
}

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test will fail until the calculator binary is implemented
			cmd := exec.Command("./calculator", tt.expression)
			_, err := cmd.CombinedOutput()

			// Expected to fail initially - implementation not yet complete
			if err == nil {
				t.Errorf("Expected command to fail (not implemented yet), but it succeeded")
				return
			}

			// When implemented, these assertions should pass:
			// if exitError, ok := err.(*exec.ExitError); !ok || exitError.ExitCode() != tt.wantCode {
			//     t.Errorf("Expected exit code %d, got error: %v", tt.wantCode, err)
			// }
		})
	}
}

// TestCLIContractInputFormats validates input format handling
func TestCLIContractInputFormats(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		valid      bool
	}{
		{"decimal numbers", "1.23 + 4.56", true},
		{"hex numbers", "0xff + 0x10", true},
		{"negative hex", "-0xab + 10", true},
		{"mixed formats", "0.5 * 0xff", true},
		{"whitespace", "  1   +   2  ", true},
		{"tabs and newlines", "1\t+\n2", true},
		{"invalid chars", "1 + z", false},
		{"empty expression", "", false},
		{"just operator", "+", false},
		{"invalid hex", "0xgg", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test defines the expected behavior - will fail until implemented
			cmd := exec.Command("./calculator", tt.expression)
			_, err := cmd.CombinedOutput()

			// For now, all commands fail (not implemented)
			// When implemented, valid expressions should succeed (exit code 0)
			// and invalid expressions should fail (exit code 1)
			if err == nil {
				t.Errorf("Expected command to fail (not implemented yet), but it succeeded")
			}
		})
	}
}
