package calculator

import (
	"errors"
	"math/big"
	"regexp"
	"strings"
)

// TokenType represents the type of a parsed token
type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LEFT_PAREN
	RIGHT_PAREN
	EOF
)

// String returns the string representation of a TokenType
func (t TokenType) String() string {
	switch t {
	case NUMBER:
		return "NUMBER"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULTIPLY:
		return "MULTIPLY"
	case DIVIDE:
		return "DIVIDE"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

// Number represents a mathematical number that can be either decimal or hexadecimal
type Number struct {
	value      interface{} // *big.Int for hex, *big.Float for decimal
	isHex      bool
	isNegative bool
}

// NewNumber creates a new Number from a string representation
func NewNumber(s string) (*Number, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty number string")
	}

	// Check for negative
	isNegative := false
	if strings.HasPrefix(s, "-") {
		isNegative = true
		s = s[1:]
	}

	// Check if hexadecimal (starts with 0x or 0X)
	if strings.HasPrefix(strings.ToLower(s), "0x") {
		// Hexadecimal number
		value, ok := new(big.Int).SetString(s, 0) // 0 = detect base automatically
		if !ok {
			return nil, errors.New("invalid hexadecimal number format")
		}
		return &Number{
			value:      value,
			isHex:      true,
			isNegative: isNegative,
		}, nil
	}

	// Decimal number (including exponential notation)
	value, _, err := big.ParseFloat(s, 10, 1000, big.ToNearestEven)
	if err != nil {
		return nil, errors.New("invalid decimal number format")
	}

	if isNegative {
		value.Neg(value)
	}

	return &Number{
		value:      value,
		isHex:      false,
		isNegative: isNegative,
	}, nil
}

// String returns the string representation of the number
func (n *Number) String() string {
	if n.isHex {
		if n.isNegative {
			return "-0x" + n.value.(*big.Int).Text(16)
		}
		return "0x" + n.value.(*big.Int).Text(16)
	}

	// Handle both big.Float and big.Int in the value field
	switch v := n.value.(type) {
	case *big.Float:
		if v.IsInt() {
			i, _ := v.Int(nil)
			return i.String()
		}
		s := v.Text('f', 50)
		if idx := strings.Index(s, "."); idx != -1 {
			s = strings.TrimRight(s, "0")
			s = strings.TrimRight(s, ".")
		}
		return s
	case *big.Int:
		return v.String()
	default:
		return "0"
	}
}

// ToFloat converts the number to big.Float for arithmetic operations
func (n *Number) ToFloat() *big.Float {
	if n.isHex {
		// Convert hex big.Int to big.Float
		i := n.value.(*big.Int)
		f := new(big.Float).SetInt(i)
		return f
	}
	return new(big.Float).Copy(n.value.(*big.Float))
}

// ValidateNumberFormat validates number format according to specification
func ValidateNumberFormat(s string) error {
	s = strings.TrimSpace(s)

	// Remove leading minus
	if strings.HasPrefix(s, "-") {
		s = s[1:]
	}

	if strings.HasPrefix(strings.ToLower(s), "0x") {
		// Hex validation: must contain only valid hex characters after 0x
		hexPart := s[2:]
		if hexPart == "" {
			return errors.New("hex number cannot be empty after 0x")
		}
		validHex := regexp.MustCompile(`^[A-Fa-f0-9]+$`)
		if !validHex.MatchString(hexPart) {
			return errors.New("invalid characters in hexadecimal number")
		}
	} else {
		// Decimal validation including exponential notation
		// Allow: digits, decimal point, E/e, +, -
		validDecimal := regexp.MustCompile(`^[0-9]*\.?[0-9]+(?:[Ee][+-]?[0-9]+)?$`)
		if !validDecimal.MatchString(s) {
			return errors.New("invalid decimal number format")
		}

		// Try to parse to ensure it's valid
		_, _, err := big.ParseFloat(s, 10, 1000, big.ToNearestEven)
		if err != nil {
			return errors.New("invalid exponential notation format")
		}
	}

	return nil
}

// Token represents a parsed element from the input expression
type Token struct {
	Type     TokenType
	Value    string
	Position int
}

// NewToken creates a new Token
func NewToken(tokenType TokenType, value string, position int) *Token {
	return &Token{
		Type:     tokenType,
		Value:    value,
		Position: position,
	}
}

// Expression represents a parsed mathematical expression tree
type Expression struct {
	Left     *Expression
	Right    *Expression
	Operator TokenType
	Value    *Number
}

// NewNumberExpression creates a leaf expression with a number value
func NewNumberExpression(number *Number) *Expression {
	return &Expression{
		Value: number,
	}
}

// NewBinaryExpression creates a binary operation expression
func NewBinaryExpression(left *Expression, operator TokenType, right *Expression) *Expression {
	return &Expression{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

// CalculatorResult represents the final result of a calculation
type CalculatorResult struct {
	Result *Number
	Error  error
}

// NewSuccessResult creates a successful calculation result
func NewSuccessResult(result *Number) *CalculatorResult {
	return &CalculatorResult{
		Result: result,
		Error:  nil,
	}
}

// NewErrorResult creates an error result
func NewErrorResult(err error) *CalculatorResult {
	return &CalculatorResult{
		Result: nil,
		Error:  err,
	}
}

// IsSuccess returns true if the result represents a successful calculation
func (cr *CalculatorResult) IsSuccess() bool {
	return cr.Error == nil
}

// IsError returns true if the result represents an error
func (cr *CalculatorResult) IsError() bool {
	return cr.Error != nil
}
