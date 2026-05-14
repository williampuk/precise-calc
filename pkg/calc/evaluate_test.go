package calculator

import (
	"math/big"
	"testing"
)

func TestNumber_NewNumber(t *testing.T) {
	tests := []struct {
		input   string
		wantVal string
		wantHex bool
		wantErr bool
	}{
		{"42", "42", false, false},
		{"3.14", "3.14", false, false},
		{"-5", "-5", false, false},
		{"0xFF", "0xff", true, false},
		{"", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			n, err := NewNumber(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if n.isHex != tt.wantHex {
				t.Errorf("isHex = %v, want %v", n.isHex, tt.wantHex)
			}
			if n.String() != tt.wantVal {
				t.Errorf("String() = %v, want %v", n.String(), tt.wantVal)
			}
		})
	}
}

func TestNumber_ToFloat(t *testing.T) {
	n, _ := NewNumber("42")
	f := n.ToFloat()
	expected := new(big.Float).SetInt64(42)
	if f.Cmp(expected) != 0 {
		t.Errorf("ToFloat() = %v, want %v", f, expected)
	}
}

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		input    string
		wantType TokenType
		wantVal  string
	}{
		{"42", NUMBER, "42"},
		{"+", PLUS, "+"},
		{"-", MINUS, "-"},
		{"*", MULTIPLY, "*"},
		{"/", DIVIDE, "/"},
		{"(", LEFT_PAREN, "("},
		{")", RIGHT_PAREN, ")"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			token, err := lexer.NextToken()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if token.Type != tt.wantType {
				t.Errorf("Type = %v, want %v", token.Type, tt.wantType)
			}
			if token.Value != tt.wantVal {
				t.Errorf("Value = %v, want %v", token.Value, tt.wantVal)
			}
		})
	}
}

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"42", false},
		{"2 + 3", false},
		{"2+3", false},
		{"(1 + 2) * 3", false},
		{"1 + 2 * 3", false},
		{"", true},
		{"abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			parser := NewParser(tt.input)
			_, err := parser.Parse()
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr bool
	}{
		{"42", "42", false},
		{"2 + 3", "5", false},
		{"10 - 4", "6", false},
		{"6 * 7", "42", false},
		{"20 / 4", "5", false},
		{"(3 + 4) * 2", "14", false},
		{"2 + 3 * 4", "14", false},
		{"abc", "", true},
		{"", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if tt.wantErr {
				if result.IsSuccess() {
					t.Error("expected error, got success")
				}
				return
			}
			if result.IsError() {
				t.Errorf("unexpected error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.want {
				t.Errorf("EvaluateString(%q) = %v, want %v", tt.input, result.Result.String(), tt.want)
			}
		})
	}
}

func TestEvaluate_BasicOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple addition", "1 + 2", "3"},
		{"simple subtraction", "5 - 3", "2"},
		{"simple multiplication", "4 * 5", "20"},
		{"simple division", "20 / 4", "5"},
		{"combined operations", "2 + 3 + 4", "9"},
		{"all operations", "10 + 5 - 3 * 2", "9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_OperatorPrecedence(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"multiplication before addition", "2 + 3 * 4", "14"},
		{"division before subtraction", "10 - 8 / 2", "6"},
		{"mixed operations with precedence", "5 + 2 * 3 - 4 / 2", "9"},
		{"left-to-right for same precedence", "8 / 2 * 3", "12"},
		{"complex expression", "10 + 5 * 2 - 3 * 4 / 2", "14"},
		{"only multiplication and division", "8 * 4 / 2 * 3", "48"},
		{"full precedence test", "2 + 3 * 4 - 10 / 2", "9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_Parentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"parentheses override", "(2 + 3) * 4", "20"},
		{"nested parentheses", "((1 + 2) * 3) + 4", "13"},
		{"complex parentheses", "(10 + 5) * (3 - 1)", "30"},
		{"deeply nested", "(((1 + 2)))", "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_HexNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple hex", "0xff + 1", "256"},
		{"hex multiplication", "0x10 * 10", "160"},
		{"negative hex", "-0xff + 256", "1"},
		{"hex and decimal", "0xa + 1.5", "11.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_ExponentialNotation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic exponential", "1E3 + 2", "1002"},
		{"lowercase e", "2.5e-2 + 1", "1.025"},
		{"large exponential", "5e6 * 2", "10000000"},
		{"negative exponent", "1e-5 + 0.00001", "0.00002"},
		{"scientific notation", "1.5E10 / 3", "5000000000"},
		{"complex exponential", "2e3 + 3e2", "2300"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_NoSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no spaces addition", "2+3", "5"},
		{"no spaces mixed", "2+3*4", "14"},
		{"no spaces parentheses", "(2+3)*4", "20"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
				return
			}
			if result.Result.String() != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result.Result.String())
			}
		})
	}
}

func TestEvaluate_ErrorCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"incomplete expression", "1 + "},
		{"empty expression", ""},
		{"just operator", "+"},
		{"two operators", "1 + * 2"},
		{"invalid characters", "1 + abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if !result.IsError() {
				t.Errorf("Expected error, got success")
			}
		})
	}
}

func TestEvaluate_DivisionByZero(t *testing.T) {
	result := EvaluateString("10 / 0")
	if !result.IsError() {
		t.Errorf("Expected error for division by zero")
	}
}

func TestEvaluate_PreciseDecimals(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"very small decimals", "0.0000000000000001 + 0.0000000000000002"},
		{"large precision preservation", "1.0000000000000001 * 2"},
		{"decimal subtraction precision", "1.0000000000000001 - 0.0000000000000001"},
		{"nested small decimals", "0.1 + 0.2 + 0.3"},
		{"complex precision", "0.0000000000000001 + 0.1 + -99999999999999"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvaluateString(tt.input)
			if result.IsError() {
				t.Errorf("Expected success, got error: %v", result.Error)
			}
		})
	}
}
