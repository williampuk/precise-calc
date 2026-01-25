# precise-calc

A high-precision CLI calculator written in Go, featuring arbitrary-precision arithmetic, hexadecimal support, and an interactive mode.

## Features

- **Arbitrary Precision**: Uses Go's `math/big` package for precise decimal calculations without floating-point limitations
- **Multiple Number Formats**: Supports decimal and hexadecimal (`0x`) numbers
- **Exponential Notation**: Handles scientific notation (e.g., `1.5e10`, `5e-3`)
- **Operator Precedence**: Correctly evaluates `*` and `/` before `+` and `-`
- **Parentheses Support**: Group expressions with `(...)` for explicit precedence
- **Interactive Mode**: REPL-style calculator with `>` prompt
- **Clean Error Handling**: Descriptive error messages with appropriate exit codes

## Installation

### From Source

```bash
git clone https://github.com/yourusername/precise-calc.git
cd precise-calc
make build
```

### Pre-built Binary

Download the appropriate binary for your platform from the [releases page](https://github.com/yourusername/precise-calc/releases).

## Usage

### Single Expression Mode

Pass an expression as a command-line argument:

```bash
$ ./calculator "2 + 3"
5
$ ./calculator "(10 + 5) * 2"
30
$ ./calculator "0xff + 1"
256
$ ./calculator "1.5e10 / 1000"
15000000
```

### Interactive Mode

Run without arguments to enter interactive mode:

```bash
$ ./calculator
Interactive mode. Press Ctrl+C or Ctrl+D to exit.
> 2 + 3
= 5
> 10 / 4
= 2.5
> 0xff * 2
= 510
> (1 + 2) * (3 + 4)
= 21
>
```

Press `Ctrl+C` or `Ctrl+D` to exit interactive mode.

## Supported Operations

| Operator | Description     |
|----------|-----------------|
| `+`      | Addition        |
| `-`      | Subtraction     |
| `*`      | Multiplication  |
| `/`      | Division        |
| `(` `)`  | Grouping        |

### Number Formats

- **Decimal**: `123`, `-456`, `3.14159`
- **Hexadecimal**: `0xff`, `0xFF`, `-0xab`
- **Exponential**: `1e3`, `2.5e-2`, `1.5E10`

### Examples

```bash
# Basic arithmetic
$ ./calculator "10 + 5"
15

# Operator precedence
$ ./calculator "2 + 3 * 4"
14

# Parentheses override precedence
$ ./calculator "(2 + 3) * 4"
20

# Hexadecimal calculations
$ ./calculator "0xff + 0x1"
256

# Mixed formats
$ ./calculator "0xa + 1.5"
11.5

# Exponential notation
$ ./calculator "1E6 / 1000"
1000

# Precise decimals (no floating-point errors)
$ ./calculator "0.1 + 0.2"
0.3
```

## Architecture

```
precise-calc/
├── cmd/
│   └── calculator/
│       └── main.go            # Entry point, CLI argument dispatch, interactive mode
├── pkg/
│   └── calculator/
│       ├── types.go           # Number, Token, Expression, CalculatorResult types
│       ├── parser.go          # Lexer and Parser for expression parsing
│       ├── evaluator.go       # Expression evaluation with math/big
│       └── evaluate_test.go   # Unit tests
└── tests/
    └── integration/
        └── cli_test.go        # Integration tests
```

### Components

1. **Lexer** (`pkg/calculator/parser.go`): Tokenizes input strings into meaningful tokens (numbers, operators, parentheses)

2. **Parser** (`pkg/calculator/parser.go`): Uses recursive descent parsing to build an Abstract Syntax Tree (AST) from tokens, respecting operator precedence

3. **Evaluator** (`pkg/calculator/evaluator.go`): Traverses the AST and computes the result using `math/big` for arbitrary-precision arithmetic

4. **Number Type** (`pkg/calculator/types.go`): Wrapper around `big.Float` and `big.Int` supporting decimal and hexadecimal formats

### Error Handling

- Exit code `0`: Successful evaluation
- Exit code `1`: Error (invalid expression, division by zero, etc.)

Error messages are printed to stderr with ANSI color highlighting in interactive mode.

## Testing

Run all tests:

```bash
make test
```

Or manually:

```bash
go test ./pkg/calculator/...
go test ./tests/integration/...
```

### Test Coverage

- **Unit Tests**: Core functionality (lexer, parser, evaluator, number parsing)
- **Integration Tests**: CLI workflow, exit codes, interactive mode behavior

## Limitations

- Maximum input length: 4096 characters
- No modulus (`%`) operator
- No power/exponentiation operator (`^`)
- No functions (sin, cos, sqrt, etc.)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Go standard library's `math/big` package for arbitrary-precision arithmetic
- Inspired by `bc` and other command-line calculators
