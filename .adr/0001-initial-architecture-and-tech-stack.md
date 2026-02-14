# ADR 0001: Initial Architecture and Tech Stack

- Status: Accepted
- Date: 2026-02-14

## Context
The project targets a Debian-installable local agent that analyzes Nginx structured logs to classify human vs bot traffic, aggregate sessions/metrics, and expose a local API and web UI.

Constraints:
- Must run without adding JavaScript tracking scripts to target websites.
- Must be deployable and operable via `.deb` package and `systemd`.
- Must allow future extension to an optional agent-to-central control plane.
- Must keep strict architectural boundaries to reduce coupling and operational risk.

## Decision
- Backend daemon and local API will be implemented in **Go**.
- Backend architecture will follow **hexagonal architecture** with explicit ports/adapters and strict dependency direction.
- Frontend will be implemented in **React + TypeScript** with a **component-based** architecture.
- Primary data source is **Nginx structured access logs** (log-based analytics, no client-side JS required).
- Packaging as a Debian `.deb` is a **first-class concern**, including `systemd` service and Nginx integration hooks.
- ADRs are stored in the top-level `.adr/` directory.

## Consequences
Positive:
- Clear backend boundaries improve testability, maintainability, and future adapter swaps.
- Go offers strong suitability for daemon/API workloads and system-level packaging targets.
- React + TypeScript supports maintainable UI composition with strong typing.
- Log-based analytics reduces privacy risk and deployment friction on target websites.
- `.deb`-first design aligns engineering work with real operational delivery.

Trade-offs:
- Strict hexagonal boundaries add initial structure overhead.
- Log-only analytics may limit behavioral depth compared to full client instrumentation.
- Two-language stack requires consistent interface contracts and CI discipline.

Follow-ups:
- Define exact repository layout and coding conventions.
- Define security, threat model, and packaging hardening baselines.
- Specify MVP scope before central control-plane implementation.
