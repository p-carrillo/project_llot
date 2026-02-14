# Data Model Standard

## Core Entities
- `request_event`: parsed Nginx log record with normalized dimensions and timings.
- `session_estimation`: grouped request behavior for inferred visitor sessions.
- `window_aggregation`: metrics rolled up by time window and dimensions.
- `host_site`: logical site/vhost context.
- `bot_score`: classification output with score, label, and rationale tags.

## Identifiers
- `request_event_id`: deterministic hash over stable event fields.
- `session_id`: deterministic session key with configurable timeout rules.
- `host_site_id`: stable key derived from host/vhost identity.
- `aggregation_key`: compound key of window + dimensions.

## Retention and TTL
- Raw request events: short TTL by default.
- Session estimation and aggregations: longer TTL with bounded upper limit.
- TTL configuration must support compliance-driven minimization.

## Schema Evolution
- Additive-first changes preferred.
- Breaking schema changes require ADR + migration plan.
- Version schema and transformation logic explicitly.
- Backfill/rebuild strategy must be documented before release.
