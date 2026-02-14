# Testing Standard

## Unit Tests
- Domain and application logic coverage is mandatory.
- Table-driven tests for classification and aggregation rules.
- Deterministic fixtures for parser behavior.

## Integration Tests
- End-to-end ingest -> storage -> API flow.
- Adapter contract tests for storage and API boundaries.

## Nginx Config Tests
- Validate generated/installable snippets with `nginx -t`.
- Negative tests for invalid config and rollback path.

## Log Replay Tests
- Replay real-like structured logs with expected metric assertions.
- Include malformed and adversarial log samples.

## Load Tests
- Measure ingestion throughput and API latency under realistic volume.
- Validate backpressure behavior and bounded resource usage.

## Security Tests
- AuthN/AuthZ negative-path coverage.
- Input validation and parser fuzzing.
- Redaction verification in logs and API errors.
