# API Conventions

## Protocol and Versioning
- REST-style JSON API.
- Base path: `/api/v1`.
- Backward-incompatible changes require new API version.

## Resource Design
- Nouns for resources, verbs in actions only when unavoidable.
- Consistent filter and sort query conventions.
- UTC timestamps in RFC 3339 format.

## Time-Range Queries
- Standard parameters: `from`, `to`, `step`.
- Defaults must be explicit and documented.
- Reject invalid or unbounded ranges with clear errors.

## Pagination
- Cursor-based pagination preferred for large time-series datasets.
- Fallback offset pagination allowed for small, bounded resources.
- Response includes page metadata and next cursor when applicable.

## Error Format
Standard error envelope:
- `error.code`
- `error.message`
- `error.details` (optional)
- `error.request_id`

## Authentication Model
- Local-token or session auth for non-dev modes.
- Auth required by default; anonymous read-only mode must be explicit opt-in.
- Authorization checks at handler boundary and application service boundary.
