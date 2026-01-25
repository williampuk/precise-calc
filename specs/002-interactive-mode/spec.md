# Feature Specification: Interactive Calculator Mode

**Feature Branch**: `002-interactive-mode`  
**Created**: 2026-01-26  
**Status**: Draft  
**Input**: User description: "In addition to the current program usage, I want to add one more. Keep existing: The original usage requires one argument, which is the expression. Please keep this. New usage: However, if there is no argument, e.g. $ precise-calc, the program will go to an interactive mode, which shows an '>' and prompts for a math expression..."

## Clarifications

### Session 2026-01-26

- Q: Error message formatting for interactive mode → A: Use color or highlighting to distinguish errors from results
- Q: Handling empty expression in interactive mode → A: Simply redisplay the prompt with no message
- Q: Maximum expression length handling → A: Enforce a reasonable limit (e.g., 4096 characters) with clear error message
- Q: Ctrl+D (EOF) handling in interactive mode → A: Exit cleanly, same as Ctrl+C
- Q: Input line handling after calculation → A: Keep both the user's input and the result on screen (show all history)

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Single Expression Evaluation (Priority: P1)

As a user, I want to evaluate a math expression by passing it as a command-line argument, so that I can quickly get results without entering interactive mode.

**Why this priority**: This is the existing primary use case that must be preserved. Maintaining backward compatibility is critical for users who rely on the current CLI behavior.

**Independent Test**: Can be fully tested by running `precise-calc "1 + 2"` and verifying the output is `3`. Delivers value by providing quick, non-interactive calculation results.

**Acceptance Scenarios**:

1. **Given** the program is invoked with an expression argument, **When** the user executes `precise-calc "2 + 3"`, **Then** the program outputs `5` and exits.
2. **Given** the program is invoked with a complex expression, **When** the user executes `precise-calc "(3 + 4) * 2"`, **Then** the program outputs `14` and exits.
3. **Given** the program is invoked with invalid input, **When** the user executes `precise-calc "1 +"`, **Then** the program displays an error with visual distinction (color/highlighting) and exits with a non-zero status.

---

### User Story 2 - Interactive Mode Entry (Priority: P1)

As a user, I want to start the calculator without arguments and enter an interactive session, so that I can perform multiple calculations in sequence without restarting the program.

**Why this priority**: This is the new primary feature that enables iterative calculation workflows. Users performing multiple related calculations benefit from not having to repeatedly type the program name.

**Independent Test**: Can be fully tested by running `precise-calc` with no arguments, typing `1 + 2`, pressing Enter, and verifying `= 3` is displayed. Delivers value by enabling an interactive calculation session.

**Acceptance Scenarios**:

1. **Given** the program is invoked with no arguments, **When** the program starts, **Then** it displays a `>` prompt and waits for user input.
2. **Given** the program is in interactive mode, **When** the user types `5 * 5` and presses Enter, **Then** the program outputs `= 25` on a new line and displays another `>` prompt.
3. **Given** the program is in interactive mode, **When** the user presses Ctrl+C or Ctrl+D, **Then** the program exits cleanly without displaying an error.

---

### User Story 3 - Interactive Calculation Loop (Priority: P1)

As a user in interactive mode, I want to enter multiple expressions sequentially and receive immediate results, so that I can perform a series of calculations efficiently.

**Why this priority**: This defines the core interactive experience. Users need reliable, repeatable prompt-response behavior for effective use of the interactive feature.

**Independent Test**: Can be fully tested by running `precise-calc` with no arguments, entering `10 / 2`, then `7 - 3`, then `4 * 2`, and verifying each result is correctly displayed with `=` prefix.

**Acceptance Scenarios**:

1. **Given** the program is in interactive mode, **When** the user enters `10 / 2`, **Then** the program outputs `= 5` and returns to the prompt.
2. **Given** the program is in interactive mode, **When** the user enters `7 - 3`, **Then** the program outputs `= 4` and returns to the prompt.
3. **Given** the program is in interactive mode, **When** the user enters `4 * 2`, **Then** the program outputs `= 8` and returns to the prompt.
4. **Given** the program is in interactive mode, **When** the user enters an invalid expression, **Then** the program displays an error with visual distinction and returns to the prompt without exiting.
5. **Given** the program is in interactive mode, **When** the user presses Enter without entering any expression, **Then** the program redispenses the `>` prompt with no error message.

---

### Edge Cases

- Empty expression entry: Program redispenses the prompt silently without error.
- Non-mathematical text: Program displays error with visual distinction.
- Long expressions: Program enforces a reasonable maximum length and displays an error when exceeded.
- Ctrl+C or Ctrl+D during session: Program exits cleanly without error.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The program MUST accept a single expression as a command-line argument and output the result immediately.
- **FR-002**: The program MUST detect when no arguments are provided and enter interactive mode.
- **FR-003**: Interactive mode MUST display a `>` prompt followed by a space at the beginning of each input line.
- **FR-004**: Interactive mode MUST output results prefixed with `=` followed by a space, preserving the user's input line above the result.
- **FR-005**: Interactive mode MUST remain active after each calculation and continue accepting new expressions.
- **FR-006**: Interactive mode MUST exit cleanly when the user presses Ctrl+C or Ctrl+D (EOF signal).
- **FR-007**: The program MUST handle errors in both argument mode and interactive mode without crashing, displaying errors with visual distinction (color or highlighting) to separate them from calculation results.
- **FR-008**: The program MUST preserve all existing calculation capabilities in both modes.
- **FR-009**: Interactive mode MUST enforce a reasonable maximum expression length and display an error when exceeded.

### Key Entities *(include if feature involves data)*

No persistent data entities required. The calculator is stateless between invocations.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can evaluate a single expression via command-line argument in under 1 second.
- **SC-002**: Interactive mode responds to user input within 500 milliseconds of pressing Enter.
- **SC-003**: 100% of users can successfully enter interactive mode by running the program without arguments.
- **SC-004**: Users can perform unlimited sequential calculations in interactive mode without program restart.
- **SC-005**: Existing single-argument usage continues to work exactly as before, with no behavioral changes.
