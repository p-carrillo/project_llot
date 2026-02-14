# AGENTS

This repository uses role-based agents to plan and deliver a Debian-installable Nginx traffic intelligence agent.

## Tech Stack
- Backend daemon and API: Go, hexagonal architecture (ports/adapters, strict dependency boundaries).
- Frontend UI: React + TypeScript, component-based design.
- Packaging: Debian `.deb` + `systemd` + Nginx drop-in integration.

## Standards Index
- [.ai/README.md](.ai/README.md)
- [.ai/standards/CONTRIBUTING.md](.ai/standards/CONTRIBUTING.md)
- [.ai/standards/ARCHITECTURE.md](.ai/standards/ARCHITECTURE.md)
- [.ai/standards/SECURITY.md](.ai/standards/SECURITY.md)
- [.ai/standards/OBSERVABILITY.md](.ai/standards/OBSERVABILITY.md)
- [.ai/standards/DATA_MODEL.md](.ai/standards/DATA_MODEL.md)
- [.ai/standards/API_CONVENTIONS.md](.ai/standards/API_CONVENTIONS.md)
- [.ai/standards/UI_UX.md](.ai/standards/UI_UX.md)
- [.ai/standards/REPO_LAYOUT.md](.ai/standards/REPO_LAYOUT.md)
- [.ai/standards/RELEASE_PACKAGING.md](.ai/standards/RELEASE_PACKAGING.md)
- [.ai/standards/TESTING.md](.ai/standards/TESTING.md)
- [.ai/standards/CODING.md](.ai/standards/CODING.md)
- [.ai/standards/GO_BEST_PRACTICES.md](.ai/standards/GO_BEST_PRACTICES.md)
- [.ai/standards/REACT_BEST_PRACTICES.md](.ai/standards/REACT_BEST_PRACTICES.md)

## General Docs
- [.ai/general/ROADMAP.md](.ai/general/ROADMAP.md)
- [.ai/general/THREAT_MODEL.md](.ai/general/THREAT_MODEL.md)
- [.ai/general/DEFINITIONS.md](.ai/general/DEFINITIONS.md)

## Agent Roles

### Product / UX Agent
Responsibilities:
- Define operator workflows and dashboard information hierarchy.
- Keep UI minimal, fast, and decision-oriented.
- Maintain MVP scope discipline for local-first usage.

Non-goals:
- Backend transport, storage internals, or packaging mechanics.
- Visual experimentation that harms readability or performance.

When to call:
- New dashboard views, filters, naming, interaction flow, and UX acceptance criteria.

### Backend / Daemon Agent
Responsibilities:
- Design and maintain domain, application services, ports, and adapters.
- Implement ingestion, classification, aggregation, storage, and local API contracts.
- Enforce strict dependency direction and failure handling.

Non-goals:
- Direct Nginx config ownership and release artifact policy.
- UI micro-interactions.

When to call:
- Domain model updates, API behavior, performance bottlenecks, and daemon lifecycle logic.

### Nginx / Infra Agent
Responsibilities:
- Define safe Nginx log format requirements and drop-in snippets.
- Own reload-safe integration procedures and config validation steps.
- Provide operational guidance for local deployment and rollback.

Non-goals:
- Business metrics semantics or bot classification policy.
- UI concerns beyond serving/static integration boundaries.

When to call:
- Log ingestion assumptions, config changes, deployment scripts, reload safety.

### Data / Analytics Agent
Responsibilities:
- Define metric semantics, session estimation, and bot score interpretation.
- Own aggregations, windows, cardinality limits, and retention behavior.
- Validate metric correctness and explainability.

Non-goals:
- Authentication policy and `.deb` packaging internals.
- UI component implementation details.

When to call:
- New KPIs, aggregation definitions, schema evolution, and analytical correctness reviews.

### Security Agent
Responsibilities:
- Enforce least privilege, secure defaults, and threat-model alignment.
- Define redaction, retention, auth, rate-limiting, and hardening baselines.
- Review high-risk code and infra changes before release.

Non-goals:
- Product prioritization and visual design decisions.

When to call:
- Any auth/data-handling changes, network exposure changes, packaging hardening, incident response prep.

### Packaging / Release Agent
Responsibilities:
- Own Debian packaging layout, install/upgrade/uninstall scripts, and `systemd` unit policy.
- Maintain signing/reproducibility expectations and rollback behavior.
- Ensure packaged defaults stay secure and operationally predictable.

Non-goals:
- Domain model design and frontend component architecture.

When to call:
- Build/release pipeline changes, artifact metadata, runtime service wiring.

### QA Agent
Responsibilities:
- Define and execute test strategy across unit/integration/replay/load/security levels.
- Enforce regression coverage for API, aggregation, and packaging operations.
- Gate release readiness against objective quality criteria.

Non-goals:
- Architectural ownership and roadmap prioritization.

When to call:
- Test plan updates, flaky behavior triage, release readiness checks.

## Task Taxonomy
- `doc-only`: Standards, ADRs, and plans without behavior changes.
- `code-only`: Implementation changes without infra/package shape changes.
- `infra-only`: Runtime/deployment/packaging configuration changes.
- `cross-cutting`: Any task spanning domain, API/UI, security, or packaging.

## Ownership & Review Map
- API contract changes: Lead `Backend / Daemon`; reviews `Data / Analytics`, `Security`, `QA`.
- Aggregation/session/bot logic: Lead `Data / Analytics`; reviews `Backend / Daemon`, `QA`.
- Nginx integration changes: Lead `Nginx / Infra`; reviews `Security`, `Packaging / Release`, `QA`.
- Auth, retention, redaction changes: Lead `Security`; reviews `Backend / Daemon`, `Data / Analytics`, `QA`.
- Debian/systemd/release changes: Lead `Packaging / Release`; reviews `Security`, `Nginx / Infra`, `QA`.
- Dashboard UX/navigation changes: Lead `Product / UX`; reviews `Frontend owners`, `QA`, `Security` for data exposure.
- Cross-cutting migrations: Lead depends on dominant risk area; mandatory reviews `Security` and `QA`.

## Working Agreement
- Always follow `.ai` standards and active ADRs.
- Propose minimal diffs with explicit intent and rollback path.
- Use deterministic, repeatable steps for build/test/release workflows.
- Never break Nginx reload safety; validate with `nginx -t` before reload.
- Default to least privilege for processes, tokens, files, and network access.
- No breaking external behavior without an ADR update.
- Document assumptions and unresolved risks in each change.
