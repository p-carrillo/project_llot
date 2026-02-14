# ADR 0003: Backend Runtime Bootstrap and Health Contract

- Status: Accepted
- Date: 2026-02-14

## Context
Phase 1 required a runnable backend baseline to validate process lifecycle, boundary wiring, and health semantics before ingestion/aggregation logic is implemented.

## Decision
Implement an initial Go runtime baseline with explicit operational contracts:
- Daemon entrypoint at `backend/cmd/agent/main.go`.
- Environment-driven config in `backend/internal/adapters/config/config.go`:
  - `AGENT_HTTP_ADDR` (default `:8080`)
  - `AGENT_LOG_LEVEL` (default `info`)
- Structured JSON logging via `log/slog`.
- Graceful shutdown behavior for `SIGINT`/`SIGTERM` with bounded timeout.
- Health contract endpoints:
  - `GET /health/live`
  - `GET /health/ready`
  - `GET /api/v1/health`
- Health logic wired through hexagonal seams:
  - port: `ReadinessChecker`
  - application service: health snapshot generation
  - adapter: static readiness checker (bootstrap placeholder)

## Consequences
Positive:
- Early executable runtime surface for operations and packaging integration.
- Health/readiness shape established for local orchestration and monitoring.
- Maintains architecture boundaries from day one.

Trade-offs:
- Readiness checker is static placeholder until real dependencies are connected.
- API surface is minimal and not representative of final analytics endpoints.

Follow-up:
- Replace static readiness adapter with dependency-aware checks (storage, ingestion pipeline, etc.).
- Add integration tests for health semantics and shutdown guarantees.
