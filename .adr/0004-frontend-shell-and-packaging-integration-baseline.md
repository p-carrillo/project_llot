# ADR 0004: Frontend Shell and Packaging Integration Baseline

- Status: Accepted
- Date: 2026-02-14

## Context
The project needs a minimal UI shell and deployment wiring path to ensure API/UI integration and Debian delivery concerns are addressed from the start, not deferred.

## Decision
Adopt the following baseline implementation decisions:

Frontend bootstrap:
- React + TypeScript app scaffolded with Vite tooling.
- Initial component-based shell with feature placeholders:
  - Overview
  - Traffic Quality
  - Sessions
  - Hosts/Sites
- Typed API client baseline using `/api/v1` as default base path.

Packaging/infra bootstrap:
- Add hardened `systemd` service template at `packaging/systemd/nginx-ti-agent.service` with least-privilege defaults (`NoNewPrivileges`, `ProtectSystem`, `ProtectHome`, etc.).
- Add Nginx integration template at `packaging/nginx/agent-ui.conf`:
  - static UI serving path
  - reverse proxy for `/api/v1/` to local daemon (`127.0.0.1:8080`)
- Add script baseline `scripts/check-nginx-config.sh` to validate Nginx configuration with `nginx -t`.

## Consequences
Positive:
- UI and packaging concerns are part of the implementation loop from early phases.
- Establishes an operator-facing shape consistent with roadmap milestones.
- Reinforces Nginx safety workflow by codifying validation as a scriptable step.

Trade-offs:
- Frontend currently contains placeholders without production feature depth.
- Nginx and systemd templates require environment-specific adjustments in deployment.

Follow-up:
- Implement real data-driven UI via API v1 endpoints.
- Add Debian maintainer scripts and upgrade/rollback behavior validation.
- Add TLS/domain runbook for external access scenarios.
