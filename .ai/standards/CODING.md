# Coding Standard

## Monorepo Conventions
- Keep backend, frontend, packaging, and docs concerns separated by directory.
- Prefer explicit boundaries and small interfaces over shared global utilities.
- No cross-layer imports that violate architecture rules.

## Naming
- Use clear, domain-driven names.
- Avoid ambiguous abbreviations unless industry-standard.
- Keep package/module names concise and intention-revealing.

## Error Handling
- Return explicit errors with contextual wrapping.
- Avoid panic for expected runtime failures.
- Map internal errors to stable API error envelopes.

## Logging
- Structured, leveled logs only.
- Include correlation IDs on request/event paths.
- Enforce redaction of sensitive fields.

## Configuration Management
- Config via file/env with explicit precedence rules.
- Validate config at startup; fail fast on invalid critical settings.
- Provide safe defaults for all security-sensitive options.

## Go-Specific Rules
- Module layout aligns with hexagonal architecture (`domain`, `application`, `ports`, `adapters`).
- Define interfaces at port boundaries, not deep in adapters.
- Dependency injection via constructor wiring; avoid hidden singletons.

## Frontend-Specific Rules
- Use component-based folder structure by feature.
- Prefer function components and hooks.
- Keep hooks focused and side-effect boundaries explicit.
- API client conventions:
  - typed request/response contracts
  - centralized error mapping
  - no ad-hoc fetch logic scattered across UI components

## Security Coding Rules
- Validate and normalize all external inputs.
- Use safe, parameterized query building.
- Enforce payload/time/range limits.
- Redact or hash sensitive values before persistence/logging.
