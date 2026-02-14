# Backend (Phase 2 Slice)

Implemented vertical slice:
- Ingest Nginx structured JSON logs.
- Classify human vs bot from user-agent heuristics.
- Estimate sessions with timeout-based grouping.
- Expose overview/window metrics on `/api/v1`.

## Environment
- `AGENT_HTTP_ADDR` (default `:8080`)
- `AGENT_LOG_LEVEL` (default `info`)
- `AGENT_MAX_BODY_BYTES` (default `1048576`)
- `AGENT_SESSION_TIMEOUT` (default `30m`)
- `AGENT_DATA_FILE` (default `./data/events.jsonl`)

## Endpoints
- `GET /health/live`
- `GET /health/ready`
- `POST /api/v1/ingest/logs`
- `GET /api/v1/metrics/overview`
- `GET /api/v1/metrics/windows`

## Example: ingest

```bash
curl -X POST http://127.0.0.1:8080/api/v1/ingest/logs \
  -H 'Content-Type: application/json' \
  -d '{
    "lines": [
      "{\"time_iso8601\":\"2026-02-14T18:00:00Z\",\"host\":\"site.local\",\"request_method\":\"GET\",\"uri\":\"/\",\"status\":200,\"remote_addr\":\"1.2.3.4\",\"http_user_agent\":\"Mozilla/5.0\"}",
      "{\"time_iso8601\":\"2026-02-14T18:01:00Z\",\"host\":\"site.local\",\"request_method\":\"GET\",\"uri\":\"/pricing\",\"status\":200,\"remote_addr\":\"5.6.7.8\",\"http_user_agent\":\"Googlebot\"}"
    ]
  }'
```

## Example: overview

```bash
curl "http://127.0.0.1:8080/api/v1/metrics/overview?from=2026-02-14T17:00:00Z&to=2026-02-14T19:00:00Z"
```

## Example: windows

```bash
curl "http://127.0.0.1:8080/api/v1/metrics/windows?from=2026-02-14T17:00:00Z&to=2026-02-14T19:00:00Z&step=1m&limit=100"
```
