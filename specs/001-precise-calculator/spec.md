# Feature Specification: Precise Calculator

## Overview
A GoLang program that performs precise mathematical operations on decimal and hexadecimal numbers while preserving precision.

## Functional Requirements

### Core Functionality
- **Input Processing**: Accept a string containing a mathematical equation
- **Number Support**:
  - Decimal numbers (signed, including exponential notation like "1E3")
  - Hexadecimal numbers (prefixed with "0x" or "-0x", unsigned big integer interpretation)
- **Operations**: Addition (+), Subtraction (-), Multiplication (x or *), Division (/)
- **Operator Precedence**: Standard mathematical precedence rules
- **Precision Preservation**: All calculations must maintain exact precision

### Input Format Specification
Input strings consist of three categories:
1. **Blank characters**: Spaces, tabs, newlines, etc. (ignored)
2. **Math operators**: Single-character: "+", "-", "x", "*", "/"
3. **Numbers**: Continuous non-blank sequences containing only [A-Fa-f0-9xEe.]

#### Number Parsing Rules
- **Hexadecimal**: Starts with "0x" or "-0x"
  - Leftmost "-" indicates negative
  - After prefix: unsigned hexadecimal big integer bytes
- **Decimal**: All other sequences (including exponential notation)
  - Leftmost "-" indicates negative
  - Support for exponential notation (e.g., "1E3", "2.5e-2")
  - Standard decimal number parsing with scientific notation

### Example Input
```
"0.0000000000000001 + 0.1 + -99999999999999 - 0xab91"
"1E3 + 2.5e-2 * 0xff"
```

## Non-Functional Requirements
- **Language**: Go
- **Precision**: Exact arithmetic (no floating point precision loss)
- **Output**: Print the calculation result

## Success Criteria
- Program accepts mathematical expressions as strings
- Correctly parses and computes decimal and hex numbers
- Maintains mathematical operator precedence
- Outputs precise calculation results
- Handles both positive and negative numbers correctly

## User Stories
1. As a developer, I want to perform precise calculations with very small decimals so that I don't lose precision
2. As a developer, I want to mix decimal and hexadecimal numbers in calculations so that I can work with different number formats
3. As a developer, I want standard operator precedence so that calculations follow mathematical rules
4. As a developer, I want to use exponential notation in decimal numbers so that I can work with scientific notation like "1E3" for 1000