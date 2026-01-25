# Data Model: Interactive Calculator Mode

## Overview

The calculator is stateless - no persistent data structures between invocations. This document describes the data flow and transient structures used during execution.

## Data Structures

### Expression (Input)

| Field | Type | Validation | Notes |
|-------|------|------------|-------|
| RawInput | string | Max 4096 chars | User-provided input line |
| Trimmed | string | Non-empty after trim | Used for empty line detection |

### Result (Output)

| Field | Type | Format | Notes |
|-------|------|--------|-------|
| Value | string | Decimal representation | Arbitrary precision from math/big |
| Prefix | string | Always "= " | Result prefix per spec |
| OutputLine | string | "= " + Value | Full output line |

### Error

| Field | Type | Format | Notes |
|-------|------|--------|-------|
| Message | string | Human-readable | Descriptive error text |
| HasColor | bool | N/A | Visual distinction flag |
| Output | string | To stderr | Error written to stderr |

## State Transitions

```
Interactive Session State Machine:

    [Idle] --no args--> [WaitingForInput]
    [WaitingForInput] --valid expression--> [DisplayResult] ---> [WaitingForInput]
    [WaitingForInput] --empty line--> [WaitingForInput] (no output)
    [WaitingForInput] --invalid expression--> [DisplayError] ---> [WaitingForInput]
    [WaitingForInput] --Ctrl+C/Ctrl+D--> [Exited]
    [DisplayResult] ---> [WaitingForInput]
    [DisplayError] ---> [WaitingForInput]
```

## Validation Rules

1. Input must not exceed 4096 characters
2. Empty input (whitespace only) triggers silent reprompt
3. Invalid mathematical expressions trigger error with visual distinction
4. Errors go to stderr, results to stdout

## No Persistence

The calculator maintains no state between invocations. Each calculation is independent.
