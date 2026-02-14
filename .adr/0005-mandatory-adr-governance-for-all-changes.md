# ADR 0005: Mandatory ADR Governance for All Changes

- Status: Accepted
- Date: 2026-02-14

## Context
During bootstrap and Phase 1 scaffolding, multiple technical decisions were implemented before documenting all of them as ADRs. This creates governance drift and reduces traceability.

## Decision
Adopt a mandatory ADR policy for this repository:
- Every modification or decision must be reflected in `.adr/`.
- Compliance can be done by:
  - creating a new ADR, or
  - updating an existing ADR when the decision is a direct continuation.
- The rule applies to documentation, code, infrastructure, packaging, and operational workflows.

## Consequences
Positive:
- Full decision traceability over time.
- Better review quality and explicit change rationale.
- Easier audits for architecture, security, and release operations.

Trade-offs:
- Additional documentation overhead per change.
- Requires discipline to keep ADR scope concise and avoid duplication.

Follow-up:
- Keep ADRs short, concrete, and linked to affected paths.
- Reject changes in review when ADR coverage is missing.
