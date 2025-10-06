package calculator

import (
	"errors"
	"math/big"
	"unicode"
)

// Lexer performs lexical analysis on input expressions
type Lexer struct {
	input    string
	position int
}

// NewLexer creates a new lexer for the given input
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
	}
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() (*Token, error) {
	// Skip whitespace
	l.skipWhitespace()

	if l.position >= len(l.input) {
		return NewToken(EOF, "", l.position), nil
	}

	char := l.input[l.position]

	switch char {
	case '+':
		l.position++
		return NewToken(PLUS, "+", l.position-1), nil
	case '-':
		l.position++
		return NewToken(MINUS, "-", l.position-1), nil
	case '*':
		l.position++
		return NewToken(MULTIPLY, "*", l.position-1), nil
	case '/':
		l.position++
		return NewToken(DIVIDE, "/", l.position-1), nil
	case 'x', 'X':
		l.position++
		return NewToken(MULTIPLY, "x", l.position-1), nil
	default:
		// Check if it's a number
		if unicode.IsDigit(rune(char)) || char == '.' || char == '-' || (char == '0' && l.peekNext() == 'x') {
			return l.readNumber()
		}

		return nil, errors.New("unexpected character: " + string(char))
	}
}

// skipWhitespace skips over whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.position < len(l.input) && unicode.IsSpace(rune(l.input[l.position])) {
		l.position++
	}
}

// peekNext returns the next character without advancing position
func (l *Lexer) peekNext() byte {
	if l.position+1 >= len(l.input) {
		return 0
	}
	return l.input[l.position+1]
}

// readNumber reads a complete number token
func (l *Lexer) readNumber() (*Token, error) {
	start := l.position

	// Handle negative sign
	if l.input[l.position] == '-' {
		l.position++
	}

	// Check for hex number
	if l.position+1 < len(l.input) && l.input[l.position] == '0' &&
		(l.input[l.position+1] == 'x' || l.input[l.position+1] == 'X') {
		// Hex number: read until non-hex character
		l.position += 2 // skip "0x"
		for l.position < len(l.input) {
			char := l.input[l.position]
			if !isHexDigit(char) {
				break
			}
			l.position++
		}
	} else {
		// Decimal number: read until non-number character (including exponential notation)
		for l.position < len(l.input) {
			char := l.input[l.position]
			if !unicode.IsDigit(rune(char)) && char != '.' && char != 'E' && char != 'e' && char != '+' && char != '-' {
				break
			}
			// Handle signs in exponential notation
			if (char == '+' || char == '-') && l.position > start {
				prevChar := l.input[l.position-1]
				if prevChar != 'E' && prevChar != 'e' {
					break
				}
			}
			l.position++
		}
	}

	numberStr := l.input[start:l.position]

	// Validate the number format
	if err := ValidateNumberFormat(numberStr); err != nil {
		return nil, errors.New("invalid number format: " + numberStr + " - " + err.Error())
	}

	return NewToken(NUMBER, numberStr, start), nil
}

// isHexDigit checks if a character is a valid hexadecimal digit
func isHexDigit(char byte) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'A' && char <= 'F') ||
		(char >= 'a' && char <= 'f')
}

// Parser converts a token stream into an expression tree
type Parser struct {
	lexer     *Lexer
	current   *Token
	peekToken *Token
}

// NewParser creates a new parser for the given input
func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	parser := &Parser{lexer: lexer}

	// Initialize current and peek tokens
	parser.nextToken()
	parser.nextToken()

	return parser
}

// Parse parses the entire expression
func (p *Parser) Parse() (*Expression, error) {
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// Should be at end of input
	if p.current.Type != EOF {
		return nil, errors.New("unexpected token after expression: " + p.current.Value)
	}

	return expr, nil
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.current = p.peekToken

	token, err := p.lexer.NextToken()
	if err != nil {
		// For error handling, create an error token
		p.peekToken = NewToken(EOF, "", p.lexer.position)
	} else {
		p.peekToken = token
	}
}

// parseExpression parses an expression (handles precedence with recursive descent)
func (p *Parser) parseExpression() (*Expression, error) {
	return p.parseAddition()
}

// parseAddition parses addition and subtraction (lowest precedence)
func (p *Parser) parseAddition() (*Expression, error) {
	left, err := p.parseMultiplication()
	if err != nil {
		return nil, err
	}

	for p.current.Type == PLUS || p.current.Type == MINUS {
		operator := p.current.Type
		p.nextToken()

		right, err := p.parseMultiplication()
		if err != nil {
			return nil, err
		}

		left = NewBinaryExpression(left, operator, right)
	}

	return left, nil
}

// parseUnary parses unary expressions (-, +)
func (p *Parser) parseUnary() (*Expression, error) {
	if p.current.Type == MINUS {
		p.nextToken()
		expr, err := p.parseUnary() // Allow nested unary, like --5
		if err != nil {
			return nil, err
		}
		// Create a unary minus expression as 0 - expr
		zero := &Number{value: new(big.Float), isHex: false, isNegative: false}
		zeroExpr := NewNumberExpression(zero)
		return NewBinaryExpression(zeroExpr, MINUS, expr), nil
	}

	if p.current.Type == PLUS {
		p.nextToken()
		return p.parseUnary() // Unary plus is essentially no-op
	}

	return p.parsePrimary()
}

// parseMultiplication parses multiplication and division (higher precedence)
func (p *Parser) parseMultiplication() (*Expression, error) {
	left, err := p.parseUnary()
	if err != nil {
		return nil, err
	}

	for p.current.Type == MULTIPLY || p.current.Type == DIVIDE {
		operator := p.current.Type
		p.nextToken()

		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}

		left = NewBinaryExpression(left, operator, right)
	}

	return left, nil
}

// parsePrimary parses primary expressions (numbers, parentheses)
func (p *Parser) parsePrimary() (*Expression, error) {
	switch p.current.Type {
	case NUMBER:
		number, err := NewNumber(p.current.Value)
		if err != nil {
			return nil, err
		}

		expr := NewNumberExpression(number)
		p.nextToken()
		return expr, nil

	case EOF:
		return nil, errors.New("unexpected end of input")

	default:
		return nil, errors.New("unexpected token: " + p.current.Type.String() + " (" + p.current.Value + ")")
	}
}

// ParseExpression is the main entry point for parsing an expression string
func ParseExpression(input string) (*Expression, error) {
	parser := NewParser(input)
	return parser.Parse()
}
