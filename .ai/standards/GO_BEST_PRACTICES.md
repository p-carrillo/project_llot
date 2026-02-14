# Go Best Practices (Internet-Backed)

Last reviewed: 2026-02-14

## Project Structure
- Use module-first layout with clear package boundaries.
- Keep domain policies independent from transport/storage details.
- Favor small, cohesive packages; avoid cyclic dependencies.

## Error Handling
- Treat errors as values and return early on failure.
- Wrap errors with context and use `errors.Is` / `errors.As` for inspection.
- Avoid exposing low-level wrapped errors as public API contracts unless intentional.

## Context Usage
- Pass `context.Context` as the first parameter for request-scoped operations.
- Do not store context in structs in normal API design.
- Honor cancellation and deadlines in adapters.

## Concurrency
- Make goroutine lifetime and shutdown semantics explicit.
- Use context cancellation, bounded worker pools, and backpressure.
- Avoid shared mutable state unless synchronization semantics are clear.

## Testing
- Prefer table-driven tests for deterministic logic.
- Add fuzz tests for parser and boundary-heavy logic.
- Keep tests hermetic and deterministic by default.

## Logging
- Prefer structured logging with stable key names.
- Use `log/slog` (or compatible structured abstraction) for consistent fields.
- Include request/event correlation IDs in log records.

## Linting
- Enforce `gofmt` and static checks in CI.
- Use `golangci-lint` with a reviewed, versioned configuration.
- Treat lint baseline drift as a quality issue, not optional cleanup.

## Dependency Management
- Use Go modules and commit `go.mod` + `go.sum`.
- Regularly run `go mod tidy` and dependency update checks.
- Pin toolchain/module behavior intentionally and document exceptions.

## Do / Don’t Checklist
Do:
- Pass `context.Context` through call chains.
- Return wrapped, inspectable errors.
- Keep interfaces at boundary seams (ports).
- Write tests and fuzzers for parsers/normalizers.
- Keep package APIs narrow and explicit.

Don’t:
- Store request contexts in long-lived structs.
- Use panic for routine error control flow.
- Leak adapter internals into domain/application layers.
- Launch unmanaged goroutines without shutdown paths.
- Add dependencies without ownership and update policy.

## Sources
- Effective Go: https://go.dev/doc/effective_go
- Go Code Review Comments: https://go.dev/wiki/CodeReviewComments
- Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
- Contexts and structs: https://go.dev/blog/context-and-structs
- `context` package docs: https://pkg.go.dev/context
- Structured logging (`log/slog`): https://go.dev/blog/slog
- Managing dependencies: https://go.dev/doc/modules/managing-dependencies
- Go Modules Reference: https://go.dev/ref/mod
- Go fuzzing tutorial: https://go.dev/doc/tutorial/fuzz
- golangci-lint docs: https://golangci-lint.run/docs/
