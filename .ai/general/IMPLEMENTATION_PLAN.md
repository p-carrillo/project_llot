# Implementation Plan (General)

Canonical path: `.ai/general/IMPLEMENTATION_PLAN.md`.

## Goal
Deliver an installable local agent that ingests Nginx structured logs, classifies human vs bot, aggregates sessions/metrics, serves a local API/UI, and can be published behind a configured domain.

## Phase 0: Scope Lock and Contracts
Objective:
- Convert current standards into executable backlog and acceptance criteria.

Deliverables:
- MVP contract for ingest, aggregation, API, UI, packaging.
- API v1 endpoint list and response contracts.
- Risk register for security and operational constraints.

Exit Criteria:
- Team agrees on M1 acceptance criteria with explicit non-goals.
- No unresolved architecture or ownership ambiguity.

## Phase 1: Repository Skeleton and Runtime Baseline
Objective:
- Create production-oriented repo structure with minimal runtime wiring.

Deliverables:
- `backend/`, `frontend/`, `packaging/`, `scripts/` scaffolding aligned to standards.
- Backend hexagonal module layout (`domain`, `application`, `ports`, `adapters`).
- Frontend React+TS app shell with component-based folder structure.
- Basic config loading, structured logging, health endpoints stubs.

Exit Criteria:
- Project builds in CI with placeholder wiring and coding/lint checks enabled.
- Architecture boundaries are enforceable by import/package conventions.

## Phase 2: Backend Vertical Slice (MVP Core)
Objective:
- Ship end-to-end analytics core from logs to API.

Deliverables:
- Nginx structured log ingestion adapter with strict parsing and validation.
- Enrichment + baseline bot scoring flow.
- Session estimation and time-window aggregations.
- Storage adapter for raw/derived data with TTL controls.
- `/api/v1` endpoints for overview and time-range queries.

Exit Criteria:
- Deterministic replay dataset produces expected metrics.
- API latency and correctness meet initial SLO targets for MVP loads.

## Phase 3: Minimal UI and Operator Workflows
Objective:
- Provide a fast, clear dashboard over API v1.

Deliverables:
- React+TS dashboard pages: Overview, Traffic Quality, Sessions, Hosts/Sites.
- Default filters (24h, all hosts, bot toggle) and typed API client.
- Accessibility baseline (semantic structure, keyboard navigation, focus visibility).

Exit Criteria:
- Operators can answer MVP questions without raw log inspection.
- UI performance acceptable on realistic local datasets.

## Phase 4: Debian Packaging and Nginx Integration
Objective:
- Make installation and operations repeatable on Debian-based systems.

Deliverables:
- `.deb` package layout, install/upgrade/uninstall scripts.
- `systemd` unit with hardening defaults.
- Nginx drop-in templates and validation workflow (`nginx -t` before reload).
- Rollback logic for invalid Nginx config or failed upgrade.

Exit Criteria:
- Fresh install and upgrade are reproducible.
- Nginx never reloads invalid config; rollback path is verified.

## Phase 5: Hardening and Quality Gates (M2)
Objective:
- Raise security and operational quality to release-ready.

Deliverables:
- Authn/authz model for non-dev API/UI modes.
- Rate limiting, input bounds, redaction and anonymization enforcement.
- Expanded tests: unit, integration, replay, load, security/fuzz.
- Signed artifacts and reproducibility checks in release pipeline.

Exit Criteria:
- Threat model mitigations map to passing tests.
- Release candidate passes security and QA gates.

## Phase 6: Domain Exposure and Production Readiness
Objective:
- Publish panel safely behind a configured domain.

Deliverables:
- DNS and Nginx vhost configuration guide.
- TLS setup (recommended Let's Encrypt flow) and renewal checks.
- Reverse proxy/static serving mode selection documented.
- Production runbook: deploy, rollback, backup, incident basics.

Exit Criteria:
- Dashboard/API reachable via configured domain over TLS.
- Access controls and rate limits validated from external network path.

## Suggested Execution Order
1. Phase 0 and 1
2. Phase 2
3. Phase 3
4. Phase 4
5. Phase 5
6. Phase 6

## Dependencies and Risks
- Log format variability across Nginx deployments can delay parser stability.
- Packaging/release quality depends on early CI reproducibility setup.
- Domain exposure must stay blocked until auth, TLS, and rate limits are active.
