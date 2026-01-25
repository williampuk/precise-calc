# Implementation Plan: Interactive Calculator Mode

**Branch**: `002-interactive-mode` | **Date**: 2026-01-26 | **Spec**: [link](spec.md)
**Input**: Feature specification from `/specs/002-interactive-mode/spec.md`

## Summary

Add interactive mode to the precise-calc CLI calculator while preserving existing single-argument usage. Users can run `precise-calc` without arguments to enter an interactive session with `>` prompts, calculating expressions sequentially until Ctrl+C or Ctrl+D is pressed.

## Technical Context

**Language/Version**: Go 1.x (standard library only, math/big for arbitrary precision)  
**Primary Dependencies**: Standard library only (math/big, strconv, bufio for interactive input)  
**Storage**: N/A (stateless CLI, no persistence)  
**Testing**: Go testing (builtin testing package)  
**Target Platform**: Cross-platform CLI (Linux, macOS, Windows)  
**Project Type**: Single CLI tool  
**Performance Goals**: Response within 500ms (SC-002), single expression under 1s (SC-001)  
**Constraints**: Max expression length ~4096 chars, must handle EOF/interrupt signals cleanly  
**Scale**: Single-user CLI, unlimited sequential calculations per session  

## Constitution Check

*No constitution rules defined - proceeding with default project standards*

## Project Structure

### Documentation (this feature)

```text
specs/002-interactive-mode/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
└── contracts/           # Phase 1 output (CLI contract)
```

### Source Code (repository root)

```text
src/
├── main.go              # Entry point, argument parsing, mode dispatch
├── calculator/
│   ├── evaluate.go      # Expression evaluation logic
│   └── errors.go        # Error types and formatting
└── interactive/
    ├── prompt.go        # Interactive session loop
    └── terminal.go      # Terminal handling, signal management

tests/
├── unit/
│   ├── evaluate_test.go
│   └── prompt_test.go
└── integration/
    └── cli_test.go      # Full CLI workflow tests
```

**Structure Decision**: Single Go project with calculator and interactive packages under src/

## Complexity Tracking

N/A - No constitution violations
