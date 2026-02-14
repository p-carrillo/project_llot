# Threat Model (General)

Canonical path: `.ai/general/THREAT_MODEL.md`.

## Actors
- External attacker over network.
- Internal authenticated operator.
- Compromised local agent process.
- Supply-chain attacker (dependency/build path).
- Misconfigured infrastructure actor.

## Assets
- Raw request events and derived metrics.
- Session estimation and bot classification outputs.
- API credentials and local secrets.
- Agent binary/package and runtime config.
- Nginx availability and configuration integrity.

## Attack Vectors
- API auth bypass/bruteforce.
- Log poisoning or malformed log injection.
- Config tampering and unsafe reload workflows.
- Privilege escalation via service account/capabilities.
- Data exfiltration from local APIs or files.
- Dependency compromise in build pipeline.

## Impact
- Incorrect analytics and operator decisions.
- Service interruption (Nginx or agent downtime).
- Sensitive metadata exposure.
- Privilege escalation on host.

## Mitigations
- Strong auth + rate limiting + audit logging.
- Input validation and strict parsing with bounded resources.
- `nginx -t` preflight + atomic rollback.
- Systemd sandboxing and least privilege.
- Secret handling with restricted file permissions.
- Signed artifacts and reproducible build checks.

## Validation Tests
- AuthN/AuthZ negative-path tests.
- Parser fuzz/replay tests for malformed logs.
- Nginx config validation/rollback integration tests.
- Privilege boundary checks in packaged runtime.
- Dependency and artifact integrity verification in CI.
