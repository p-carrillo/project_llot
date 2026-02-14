# Architecture Standard

## Backend: Hexagonal Architecture
Layers and responsibilities:
- `domain`: core entities, value objects, domain policies.
- `application`: use cases/application services orchestrating domain behavior.
- `ports`: interfaces for inbound (commands/queries) and outbound dependencies.
- `adapters`: concrete implementations for ingress/egress.

Planned adapters:
- Nginx log ingestion adapter.
- Storage adapter.
- Local REST API adapter.
- UI serving adapter (Nginx static hosting or embedded delivery path).

## Component Model
Core components:
- `ingest`
- `enrichment`
- `aggregation`
- `storage`
- `api`
- `ui`

## Dependency Rules
- Domain has zero infrastructure dependencies.
- Application depends only on domain + port abstractions.
- Adapters depend on ports/application; never the reverse.
- No adapter-to-adapter coupling without an explicit application orchestration boundary.

## Future Extension
A future control-plane mode (agent -> central panel) is an additive adapter set. It must not violate local-first operation or existing port boundaries.
