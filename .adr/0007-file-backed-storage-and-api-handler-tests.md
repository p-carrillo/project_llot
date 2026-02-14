# ADR 0007: File-Backed Storage and API Handler Test Baseline

- Status: Accepted
- Date: 2026-02-14

## Context
ADR 0006 delivered an end-to-end vertical slice but used in-memory storage, which loses data on restart and is not suitable for realistic local-agent operation.

## Decision
Adopt a file-backed persistence adapter for Phase 2 runtime:
- New storage adapter: JSON Lines repository at `backend/internal/adapters/storage/jsonl/`.
- Configured by environment variable:
  - `AGENT_DATA_FILE` (default `./data/events.jsonl`)
- Runtime wiring in daemon now uses JSONL repository by default.
- Debian/systemd baseline sets `AGENT_DATA_FILE=/var/lib/nginx-ti-agent/events.jsonl`.

Also establish API handler test baseline:
- Add HTTP handler tests for ingest + overview and invalid payload behavior.
- Keep tests focused on handler contracts and status codes.

## Consequences
Positive:
- Data survives daemon restarts in local deployments.
- Storage behavior remains simple and transparent for debugging/replay.
- API behavior has immediate test coverage at handler boundary.

Trade-offs:
- Query performance is linear scan over file contents.
- Corrupted line in data file currently fails query path and requires operator action.
- JSONL backend is still a transitional persistence layer for MVP.

Follow-up:
- Add compaction/retention strategy and corruption recovery policy.
- Introduce production-grade persistent adapter when scale targets require it.
