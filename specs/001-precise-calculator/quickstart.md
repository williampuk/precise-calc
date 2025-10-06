# Quickstart: Precise Calculator

## Overview
The precise calculator is a command-line tool that evaluates mathematical expressions with exact precision, supporting decimal numbers (including exponential notation), hexadecimal numbers, and standard mathematical operators.

## Installation & Setup

```bash
# Build the calculator
go build -o calculator ./cmd/calculator

# Make executable
chmod +x calculator
```

## Basic Usage

```bash
# Simple arithmetic
./calculator "1 + 2"
# Output: 3

# Decimal precision preservation
./calculator "0.0000000000000001 + 0.1"
# Output: 0.1000000000000001

# Hexadecimal numbers
./calculator "0xff + 1"
# Output: 256

# Mixed number formats
./calculator "0.5 * 0xff"
# Output: 127.5

# Negative numbers
./calculator "-5 + 10"
# Output: 5

# Exponential notation
./calculator "1E3 + 2"
# Output: 1002

# Scientific notation
./calculator "2.5e-2 + 1"
# Output: 1.025
```

## Validation Scenarios

### User Story 1: Precise decimal calculations
**Scenario**: Developer needs exact precision with very small decimals
```bash
./calculator "0.0000000000000001 + 0.1 + -99999999999999"
# Expected: -99999999999898.8999999999999999
```

### User Story 2: Mixed decimal and hex numbers
**Scenario**: Developer wants to combine different number formats
```bash
./calculator "0xab91 + 100.5 - 0xff"
# Expected: 45404.5
```

### User Story 3: Operator precedence
**Scenario**: Calculations follow standard mathematical precedence
```bash
./calculator "2 + 3 * 4 - 10 / 2"
# Expected: 12 (not 16, showing precedence works)
```

### User Story 4: Exponential notation support
**Scenario**: Developer can use scientific notation for large and small numbers
```bash
./calculator "1E3 + 2.5e-2 * 100"
# Expected: 1002.5
```

## Error Cases

```bash
# Invalid syntax
./calculator "1 + "
# Error: Invalid expression syntax

# Division by zero
./calculator "10 / 0"
# Error: Division by zero

# Invalid number format
./calculator "0xgg + 1"
# Error: Invalid number format
```

## Expression Format

- **Numbers**: Decimal (e.g., `1.23`) or hex (e.g., `0xff`, `-0xab`)
- **Operators**: `+` (add), `-` (subtract), `*` or `x` (multiply), `/` (divide)
- **Whitespace**: Spaces, tabs, and newlines are ignored
- **Precedence**: Standard mathematical order (multiplication/division before addition/subtraction)

## Testing

```bash
# Run contract tests
go test ./specs/001-precise-calculator/contracts/...

# Run all tests
go test ./...

# Build and test integration
go build -o calculator ./cmd/calculator
echo "1 + 2" | xargs ./calculator
```