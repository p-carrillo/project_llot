# ADR 0006: Phase 2 Vertical Slice for Ingestion and Metrics API

- Status: Accepted
- Date: 2026-02-14

## Context
After Phase 1 bootstrap, the backend only exposed health endpoints. The implementation plan requires an executable analytics slice: ingest structured Nginx logs, classify traffic, estimate sessions, and expose API metrics.

## Decision
Implement a Phase 2 backend vertical slice with these boundaries and contracts:

- Ingestion contract:
  - `POST /api/v1/ingest/logs` with JSON body `{ "lines": ["<nginx-json-line>", ...] }`.
  - Supports up to 10,000 lines per request.
  - Body size bounded by `AGENT_MAX_BODY_BYTES`.
- Metrics contracts:
  - `GET /api/v1/metrics/overview?from=<RFC3339>&to=<RFC3339>&host=<optional>`
  - `GET /api/v1/metrics/windows?from=<RFC3339>&to=<RFC3339>&step=<duration>&limit=<n>&cursor=<offset>&host=<optional>`
- Processing behavior:
  - Parse Nginx structured JSON lines.
  - Heuristic bot/human/unknown classification from user-agent.
  - Session estimation with timeout (`AGENT_SESSION_TIMEOUT`, default `30m`).
  - Hash remote IP before storage to avoid plain address persistence.
- Storage for Phase 2:
  - In-memory repository adapter as an explicit temporary implementation.

## Consequences
Positive:
- End-to-end analytics path is now executable for local validation and replay workflows.
- API contracts for overview and windows can be consumed by the frontend immediately.
- Security posture improves by hashing remote addresses and bounding payload sizes.

Trade-offs:
- In-memory storage is non-durable and not suitable for production retention needs.
- Bot classification is heuristic and expected to evolve.
- Window aggregation is currently computed from queried events, with scalability limits.

Follow-up:
- Replace in-memory repository with persistent storage adapter.
- Add stronger bot-scoring model and richer parsing coverage for log format variants.
- Add integration tests for API handlers and pagination semantics.
