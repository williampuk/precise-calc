# Tasks: Precise Calculator

**Input**: Design documents from `/specs/001-precise-calculator/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → If not found: ERROR "No implementation plan found"
   → Extract: tech stack (Go + math/big), libraries (standard library), structure (cmd/calculator, pkg/calculator, tests/)
2. Load optional design documents:
   → data-model.md: Extract entities (Number, Token, Expression, CalculatorResult) → model tasks
   → contracts/: cli-contract.json + cli_contract_test.go → contract test tasks
   → research.md: Extract decisions (parsing, precision) → setup tasks
3. Generate tasks by category:
   → Setup: project init, dependencies, linting
   → Tests: contract tests, integration tests for user stories
   → Core: types, parser, evaluator, CLI command
   → Integration: CLI error handling, validation
   → Polish: unit tests, performance, docs
4. Apply task rules:
   → Different files = mark [P] for parallel
   → Same file = sequential (no [P])
   → Tests before implementation (TDD)
5. Number tasks sequentially (T001, T002...)
6. Generate dependency graph
7. Create parallel execution examples
8. Validate task completeness:
   → All contracts have tests? ✓ (cli_contract_test.go)
   → All entities have models? T010-T013 create them
   → All user stories tested? T005-T008 create integration tests
9. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
- **Single Go project**: `cmd/calculator/` (CLI), `pkg/calculator/` (library), `tests/` at repository root
- Go standard layout: cmd/ for entry points, pkg/ for libraries

## Phase 3.1: Setup
- [ ] T001 Create project directory structure per implementation plan
- [ ] T002 Initialize Go module and dependencies (math/big standard library)
- [ ] T003 [P] Configure Go linting tools (gofmt, go vet, golint)

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**
- [ ] T004 [P] CLI contract test success cases in tests/contract/cli_contract_success_test.go
- [ ] T005 [P] CLI contract test error cases in tests/contract/cli_contract_error_test.go
- [ ] T006 [P] Integration test user story 1 (precise decimals) in tests/integration/test_user_story_1_test.go
- [ ] T007 [P] Integration test user story 2 (mixed formats) in tests/integration/test_user_story_2_test.go
- [ ] T008 [P] Integration test user story 3 (operator precedence) in tests/integration/test_user_story_3_test.go
- [ ] T009 [P] Integration test user story 4 (exponential notation) in tests/integration/test_user_story_4_test.go

## Phase 3.3: Core Implementation (ONLY after tests are failing)
- [ ] T010 [P] Number, Token, Expression types in pkg/calculator/types.go
- [ ] T011 [P] Parser implementation with lexical analysis in pkg/calculator/parser.go
- [ ] T012 [P] Evaluator implementation with precedence in pkg/calculator/evaluator.go
- [ ] T013 [P] CLI main entry point in cmd/calculator/main.go
- [ ] T014 Input validation and error handling in pkg/calculator/types.go (add to existing file)
- [ ] T015 Result formatting and output in pkg/calculator/evaluator.go (add to existing file)

## Phase 3.4: Integration
- [ ] T016 CLI argument parsing and result output in cmd/calculator/main.go (add to existing file)
- [ ] T017 End-to-end CLI integration test in tests/integration/test_cli_integration_test.go
- [ ] T018 Performance validation (<200ms for reasonable expressions)
- [ ] T019 Memory safety checks and big number handling

## Phase 3.5: Polish
- [ ] T020 [P] Unit tests for parser in pkg/calculator/parser_test.go
- [ ] T021 [P] Unit tests for evaluator in pkg/calculator/evaluator_test.go
- [ ] T022 [P] Unit tests for types in pkg/calculator/types_test.go
- [ ] T023 [P] CLI unit tests in cmd/calculator/main_test.go
- [ ] T024 [P] Update README.md with usage examples
- [ ] T025 Run comprehensive test suite and fix any failures
- [ ] T026 Performance benchmarking and optimization
- [ ] T027 Code documentation and comments

## Dependencies
- Tests (T004-T009) before implementation (T010-T016)
- T010 creates foundation types, blocks T014
- T011 depends on T010 for types
- T012 depends on T010-T011
- T013 depends on T010-T012, T016
- T014 depends on T010
- T015 depends on T012
- T016 depends on T013
- Implementation before polish (T017-T027)
- T017-T019 can run after core implementation complete
- T020-T023 depend on respective implementation files

## Parallel Example
```
# Launch T004-T009 together (all contract and integration tests):
Task: "CLI contract test success cases in tests/contract/cli_contract_success_test.go"
Task: "CLI contract test error cases in tests/contract/cli_contract_error_test.go"
Task: "Integration test user story 1 (precise decimals) in tests/integration/test_user_story_1_test.go"
Task: "Integration test user story 2 (mixed formats) in tests/integration/test_user_story_2_test.go"
Task: "Integration test user story 3 (operator precedence) in tests/integration/test_user_story_3_test.go"
Task: "Integration test user story 4 (exponential notation) in tests/integration/test_user_story_4_test.go"

# Then launch core implementation in parallel where possible:
Task: "Number, Token, Expression types in pkg/calculator/types.go"
Task: "Parser implementation with lexical analysis in pkg/calculator/parser.go"
Task: "Evaluator implementation with precedence in pkg/calculator/evaluator.go"
Task: "CLI main entry point in cmd/calculator/main.go"
```

## Notes
- [P] tasks = different files, no dependencies
- Verify tests fail before implementing (should see compilation errors)
- Commit after each task
- TDD approach: red tests → green implementation → refactor
- Avoid: vague tasks, same file conflicts

## Task Generation Rules
*Applied during main() execution*

1. **From Contracts**:
   - cli-contract.json defines CLI interface → T004-T005 contract tests [P]
   - cli_contract_test.go has examples → split into success/error test files

2. **From Data Model**:
   - Number, Token, Expression, CalculatorResult entities → T010 types.go [P]
   - Validation rules → T014 (add to types.go)

3. **From User Stories**:
   - 4 user stories from quickstart → T006-T009 integration tests [P]

4. **From Research**:
   - Technical decisions guide implementation tasks (parsing, evaluation, CLI)

5. **Ordering**:
   - Setup → Tests → Types/Models → Parser/Evaluator → CLI → Validation → Polish
   - Dependencies prevent true parallel conflicts

## Validation Checklist
*GATE: Checked by main() before returning*

- [x] All contracts have corresponding tests (T004-T005 cover cli-contract.json)
- [x] All entities have model tasks (T010 creates all data model entities)
- [x] All tests come before implementation (T004-T009 before T010-T016)
- [x] Parallel tasks truly independent (different file paths, no shared dependencies)
- [x] Each task specifies exact file path (yes, all tasks have explicit paths)
- [x] No task modifies same file as another [P] task (validated)