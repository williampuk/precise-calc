# CLI Contract: precise-calc

## Usage Modes

### Mode 1: Single Expression (Argument Mode)

```bash
precise-calc "<expression>"
```

**Example:**
```bash
$ precise-calc "2 + 3"
5
$ precise-calc "(3 + 4"
14
$ precise-calc ") * 2invalid"
Error: invalid expression
# Exit code: 1
```

### Mode 2: Interactive Mode

```bash
precise-calc
```

**Example Session:**
```bash
$ precise-calc
> 1 + 2
= 3
> 5 * 5
= 25
> (10 + 5) / 3
= 5
>
# Ctrl+C pressed - exits cleanly
```

## Input/Output Contract

| Scenario | stdin | stdout | stderr | Exit Code |
|----------|-------|--------|--------|-----------|
| Valid expression (arg) | N/A | Result only | N/A | 0 |
| Invalid expression (arg) | N/A | N/A | Error message | 1 |
| Interactive valid | Expression line | "= result" | N/A | 0 |
| Interactive invalid | Expression line | N/A | Error (colored) | 0 |
| Interactive empty | Empty line | N/A | N/A | 0 |
| Ctrl+C / Ctrl+D | N/A | N/A | N/A | 0 |

## Output Formats

### Result (stdout)
```
<value>
```
Example: `3`

### Result with History (interactive stdout)
```
> <expression>
= <value>
```

### Error (stderr)
```
Error: <descriptive message>
```
With color highlighting for interactive mode.

## Signal Handling

| Signal | Behavior |
|--------|----------|
| Ctrl+C (SIGINT) | Clean exit, no output |
| Ctrl+D (EOF) | Clean exit, no output |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success (including interrupt exit) |
| 1 | Invalid expression in argument mode |

## Constraints

- Maximum input length: 4096 characters
- Expression syntax: Supports existing calculator grammar
- Precision: Arbitrary precision (math/big)
