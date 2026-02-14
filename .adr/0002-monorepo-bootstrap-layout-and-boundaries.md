# ADR 0002: Monorepo Bootstrap Layout and Boundaries

- Status: Accepted
- Date: 2026-02-14

## Context
ADR 0001 defined stack and architecture direction, but repository execution needed a concrete structure for implementation and ownership. Without an agreed layout, dependency boundaries and review responsibilities become inconsistent.

## Decision
Adopt the initial monorepo skeleton and boundaries:
- Top-level implementation areas: `backend/`, `frontend/`, `packaging/`, `docs/`, `scripts/`.
- Backend structure aligned with hexagonal architecture:
  - `backend/cmd/`
  - `backend/internal/domain/`
  - `backend/internal/application/`
  - `backend/internal/ports/`
  - `backend/internal/adapters/`
  - `backend/test/`
- Frontend structure aligned with component-based design:
  - `frontend/src/app/`
  - `frontend/src/components/`
  - `frontend/src/features/`
  - `frontend/src/hooks/`
  - `frontend/src/api/`
  - `frontend/test/`
- Packaging split by concern:
  - `packaging/debian/`
  - `packaging/systemd/`
  - `packaging/nginx/`

## Consequences
Positive:
- Consistent import boundaries and clearer ownership/review paths.
- Reduced coupling risk between domain, adapters, UI features, and packaging artifacts.
- Faster onboarding and execution against the implementation plan.

Trade-offs:
- Upfront structural overhead before feature delivery.
- Some directories intentionally start with placeholders until Phase 2+ implementation.

Follow-up:
- Enforce dependency rules and quality gates once toolchain is available in CI/runtime environment.
