# Quickstart: Interactive Calculator Mode

## Building

```bash
go build -o precise-calc .
```

## Quick Examples

### Single Expression (existing behavior)
```bash
./precise-calc "2 + 3"
# Output: 5

./precise-calc "(10 + 5) * 2"
# Output: 30
```

### Interactive Mode (new feature)
```bash
./precise-calc
> 1 + 2
= 3
> 100 / 4
= 25
> 3 ^ 2
= 9
# Press Ctrl+C or Ctrl+D to exit
```

## Development

### Run Tests
```bash
go test ./...
```

### Project Structure
```
src/
├── main.go           # Entry point
├── calculator/       # Evaluation logic
└── interactive/      # Interactive session handling
```

### Key Files
| File | Purpose |
|------|---------|
| `src/main.go` | Mode dispatch (arg vs interactive) |
| `src/calculator/evaluate.go` | Expression parsing and evaluation |
| `src/interactive/prompt.go` | Interactive loop and prompt handling |
| `src/interactive/terminal.go` | Signal handling and terminal config |
