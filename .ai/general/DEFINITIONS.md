# Definitions (General)

Canonical path: `.ai/general/DEFINITIONS.md`.

## Glossary
- `vhost`: Nginx virtual host configuration block serving one or more hostnames.
- `host`: Effective hostname value associated with a request.
- `session estimation`: Heuristic grouping of requests into inferred user/bot sessions.
- `bot score`: Numeric or categorical confidence that traffic is automated.
- `window`: Fixed or rolling time bucket used for aggregations.
- `upstream time`: Nginx-reported upstream response timing metric.

## Out of Scope (Current Phase)
- Full user identity analytics across devices.
- Third-party ad/conversion attribution.
- Replacing Nginx itself or acting as reverse proxy.
- Frontend-heavy BI features beyond operational dashboarding.
