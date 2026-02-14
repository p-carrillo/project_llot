# Observability Standard

## Logging
- Structured logs (JSON) by default.
- Required fields: timestamp, level, component, message, correlation_id.
- Never log secrets, full tokens, or raw sensitive payloads.

## Metrics
- Expose local metrics endpoint (machine-readable).
- Core metric families:
  - ingestion throughput/lag/errors
  - classification counts (human/bot/unknown)
  - session aggregation latency
  - API latency/error rates
  - storage size and retention actions

## Health Endpoints
- Liveness endpoint for process availability.
- Readiness endpoint for adapter dependencies and config validity.

## Debug Mode
- Explicit opt-in only.
- Time-limited and clearly marked in logs.
- Adds diagnostic depth without disabling redaction rules.

## Correlation IDs
- Generate or propagate per request/event flow.
- Carry through ingest -> enrichment -> aggregation -> API responses.

## Redaction Rules
- Redact credentials, secrets, and personally identifying fields by default.
- Hash/anonymize IP fields when anonymization mode is enabled.
- Log redaction policy version for traceability.
