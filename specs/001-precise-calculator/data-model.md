# Data Model: Precise Calculator

## Core Entities

### Number
Represents a mathematical number that can be either decimal or hexadecimal.

**Fields**:
- `value`: interface{} - The actual numeric value (big.Int for hex, big.Float for decimal)
- `isHex`: bool - True if number was parsed from hex format
- `isNegative`: bool - True if number is negative

**Validation Rules**:
- Hex numbers: Must start with "0x" or "-0x", contain only [A-Fa-f0-9] after prefix
- Decimal numbers: Standard floating point format validation including exponential notation (e.g., "1E3", "2.5e-2")
- Exponential notation: Supports both uppercase "E" and lowercase "e"
- No empty or invalid character sequences

### Token
Represents a parsed element from the input expression.

**Fields**:
- `type`: TokenType - NUMBER, PLUS, MINUS, MULTIPLY, DIVIDE, EOF
- `value`: string - The raw token text
- `position`: int - Position in original input string

**Token Types**:
- NUMBER: Numeric value (decimal or hex)
- OPERATOR: Mathematical operator (+, -, x, *, /)
- EOF: End of input

### Expression
Represents a parsed mathematical expression tree.

**Fields**:
- `left`: Expression - Left operand (can be another expression)
- `right`: Expression - Right operand
- `operator`: TokenType - The operator for this node
- `value`: Number - Leaf node value (when left/right are nil)

**Validation Rules**:
- Binary operations require both left and right operands
- Leaf nodes (numbers) have operator = NONE
- Tree structure represents operator precedence

### CalculatorResult
Represents the final result of a calculation.

**Fields**:
- `result`: Number - The computed result
- `error`: error - Any calculation error (nil on success)

**State Transitions**:
- Valid expression → Success with result
- Invalid expression → Error state
- Division by zero → Error state

## Relationships

```
Input String → [Lexer] → []Token → [Parser] → Expression Tree → [Evaluator] → CalculatorResult
```

## Data Flow

1. **Input Processing**: Raw string accepted via CLI
2. **Lexical Analysis**: String tokenized into Token stream
3. **Parsing**: Token stream converted to Expression tree with precedence
4. **Evaluation**: Expression tree evaluated using arbitrary precision arithmetic
5. **Output**: Result formatted and printed to stdout

## Type System

### Custom Types
```go
type TokenType int
type Number struct {
    value     interface{}
    isHex     bool
    isNegative bool
}
type Token struct {
    type     TokenType
    value    string
    position int
}
type Expression struct {
    left     *Expression
    right    *Expression
    operator TokenType
    value    *Number
}
```

### Precision Handling
- All decimal calculations use `big.Float` with unlimited precision
- Hexadecimal numbers parsed as `big.Int` then converted to `big.Float` for computation
- Results maintain exact precision until final output formatting