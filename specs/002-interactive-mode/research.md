# Research: Interactive Calculator Mode

## Interactive Input Handling

**Decision**: Use Go's `bufio.NewReader` with `ReadString('\n')` for line-by-line input

**Rationale**:
- `bufio.Scanner` has a 64KB default limit but can be configured for longer lines
- `ReadString` provides more control over input parsing and EOF detection
- Handles Ctrl+D (EOF) cleanly by returning io.EOF
- Compatible with signal handling for Ctrl+C (interrupt)

**Alternatives considered**:
- `fmt.Scanln`: Limited error handling, doesn't handle empty lines well
- `os.Stdin.Read`: Too low-level, requires manual buffer management
- `bufio.Scanner`: Default token length of 64K is sufficient for 4096 char limit

## Terminal Color/Highlighting for Errors

**Decision**: Use ANSI escape codes for color output (cross-platform compatible)

**Rationale**:
- Go standard library supports basic ANSI codes
- ANSI codes work on most modern terminals (Linux, macOS, Windows 10+)
- Can detect terminal capability if needed for Windows older versions

**Implementation approach**:
```go
const (
    red   = "\033[31m"
    reset = "\033[0m"
)
fmt.Fprintf(stderr, "%sError: %s%s\n", red, message, reset)
```

**Alternatives considered**:
- `fatih/color` library: Adds dependency, not needed for simple use case
- `termbox-go`: Full terminal UI library, overkill for simple highlighting

## Signal Handling (Ctrl+C, Ctrl+D)

**Decision**: Use Go's `os/signal` package with `Notify` for interrupt and quit signals

**Rationale**:
- Clean exit handling without leaving terminal in broken state
- `signal.Ignore(os.Interrupt)` in child processes to prevent signal propagation
- Ctrl+D (EOF) handled naturally by ReadString returning empty string + io.EOF

**Implementation approach**:
```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, os.Signal(syscall.SIGQUIT))
```

## Expression Length Limit

**Decision**: 4096 character limit for input validation

**Rationale**:
- Sufficient for complex expressions (nesting, multiple operations)
- Prevents memory issues from extremely long inputs
- Early validation before parsing/evaluation

## Key Implementation Patterns

1. **Mode Dispatch**: Check `len(os.Args)` - if 1, enter interactive; if 2+, evaluate and exit
2. **Input Loop**: `for { prompt(); line, err := reader.ReadString('\n'); process(line); }`
3. **Error Handling**: Separate stdout (results) from stderr (errors) per Unix conventions
4. **Result Format**: `> expression` input line preserved, `= result` on separate line

## References

- Go bufio package documentation
- Go signal handling: https://pkg.go.dev/os/signal
- ANSI color codes: https://en.wikipedia.org/wiki/ANSI_escape_code
