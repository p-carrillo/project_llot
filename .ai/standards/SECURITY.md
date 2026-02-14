# Security Standard

## Baseline Principles
- Least privilege by default for process, filesystem, and network.
- Secure-by-default configuration; opt-in for risky capabilities.
- Minimize data collection and retention.

## Runtime Hardening
Systemd baseline expectations:
- Dedicated service user/group.
- `NoNewPrivileges=true`.
- `PrivateTmp=true`.
- `ProtectSystem=strict` and explicit writable paths.
- `ProtectHome=true` unless explicitly required.
- Restrict address families/capabilities to minimum required.

## API and UI Authentication
- Local API must require authentication in non-dev modes.
- Session/token scope must be narrow and revocable.
- UI and API permissions follow least-privilege roles.

## Abuse Protection
- Rate limiting on API endpoints.
- Bounded request payload sizes and query limits.
- Explicit timeouts for all inbound/outbound operations.

## Data Handling
- Default TTLs for raw and aggregated datasets.
- Configurable retention caps with safe minimum/maximum bounds.
- IP hashing/anonymization options with keyed hashing support.

## Nginx Integration Safety
- Install as drop-in config snippets; never rewrite user configs wholesale.
- Validate configuration with `nginx -t` before any reload.
- Rollback strategy:
  - Keep previous known-good config snapshot.
  - If validation fails, restore snapshot and skip reload.
  - Emit actionable error with rollback status.
