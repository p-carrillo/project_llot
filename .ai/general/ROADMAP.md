# Roadmap (General)

Canonical path: `.ai/general/ROADMAP.md`.

## Milestones

### M1: MVP Local Agent
Scope:
- Nginx structured log ingestion.
- Human vs bot classification baseline.
- Session estimation and window aggregation.
- Local API (`/api/v1`) and minimal dashboard.
- Debian package with systemd service.

Definition of Done:
- Installable `.deb` with documented lifecycle.
- Stable local metrics dashboard for at least one host.
- Core testing layers active (unit, integration, replay).

### M2: Hardened Operations
Scope:
- Security hardening, retention controls, redaction policy enforcement.
- Stronger observability and operational diagnostics.
- Load/security test coverage expansion.

Definition of Done:
- Threat model mitigations mapped to tests.
- Reproducible, signed release artifacts.
- Verified Nginx rollback-safe integration.

### M3: Optional Central Panel Mode
Scope:
- Agent -> control-plane transport adapter.
- Policy synchronization and optional remote observability.

Definition of Done:
- Local-first mode remains functional without central dependency.
- Secure transport/auth model documented and tested.

## MVP Explicit Non-Goals
- No mandatory external SaaS dependency.
- No client-side JS tracking injection.
- No multi-cluster orchestration in first release.
