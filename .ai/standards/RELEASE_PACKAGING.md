# Release & Packaging Standard

## `.deb` Layout (High Level)
- Binary: `/usr/bin/<agent>`.
- Config: `/etc/<agent>/`.
- Runtime state: `/var/lib/<agent>/`.
- Logs: `/var/log/<agent>/` or journal-first strategy.
- Nginx drop-in snippets: `/etc/nginx/conf.d/` or `/etc/nginx/snippets/` (configurable policy).

## systemd Unit Expectations
- Service starts after network target and required dependencies.
- Hardening directives from security standard are mandatory.
- Restart policy uses bounded backoff.

## Install / Upgrade / Uninstall / Rollback
- Install: create user/group, directories, permissions, unit registration.
- Upgrade: preserve config, perform compatibility checks, restart safely.
- Uninstall: remove unit and binaries; preserve data unless purge requested.
- Rollback: keep previous package and config snapshot for quick recovery.

## Nginx Config Validation
- Validate with `nginx -t` before reload/restart.
- Never reload on failed validation.
- Emit clear failure diagnostics and rollback result.

## Signing and Reproducibility
- Sign release artifacts.
- Keep deterministic build inputs/version metadata.
- Record build provenance (toolchain version, commit, timestamp policy).
