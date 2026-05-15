package calculator

import (
	"errors"
	"math/big"
)

// Evaluate evaluates an expression tree and returns the result
func Evaluate(expr *Expression) *CalculatorResult {
	result, err := evaluateExpression(expr)
	if err != nil {
		return NewErrorResult(err)
	}

	// Convert back to Number format
	number := floatToNumber(result)
	return NewSuccessResult(number)
}

// evaluateExpression recursively evaluates an expression tree
func evaluateExpression(expr *Expression) (*big.Float, error) {
	if expr == nil {
		return nil, errors.New("nil expression")
	}

	// Leaf node (number)
	if expr.Value != nil {
		return expr.Value.ToFloat(), nil
	}

	// Binary operation
	if expr.Left == nil || expr.Right == nil {
		return nil, errors.New("invalid binary expression: missing operand")
	}

	left, err := evaluateExpression(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := evaluateExpression(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator {
	case PLUS:
		return new(big.Float).SetPrec(FloatPrecision).Add(left, right), nil
	case MINUS:
		return new(big.Float).SetPrec(FloatPrecision).Sub(left, right), nil
	case MULTIPLY:
		return new(big.Float).SetPrec(FloatPrecision).Mul(left, right), nil
	case DIVIDE:
		// Check for division by zero
		if right.Sign() == 0 {
			return nil, errors.New("division by zero")
		}
		return new(big.Float).SetPrec(FloatPrecision).Quo(left, right), nil
	default:
		return nil, errors.New("unsupported operator: " + expr.Operator.String())
	}
}

// floatToNumber converts a big.Float back to a Number
// For precision preservation, we keep it as a big.Float unless it's an integer
func floatToNumber(f *big.Float) *Number {
	// Check if it's an integer (no fractional part)
	if f.IsInt() {
		i, _ := f.Int(nil)
		return &Number{
			value:      i,
			isHex:      false,
			isNegative: i.Sign() < 0,
		}
	}

	// Keep as float
	return &Number{
		value:      f,
		isHex:      false,
		isNegative: f.Sign() < 0,
	}
}

// EvaluateString is a convenience function that parses and evaluates an expression string
func EvaluateString(input string) *CalculatorResult {
	expr, err := ParseExpression(input)
	if err != nil {
		return NewErrorResult(err)
	}

	return Evaluate(expr)
}
