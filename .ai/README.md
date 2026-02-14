# .ai Purpose

`.ai/` is the operational knowledge base for repository development standards and general planning artifacts.

## How to Use
- Treat `.ai/standards/*.md` as the default policy set for all changes.
- Use `.ai/general/*.md` for cross-cutting planning artifacts (for example, roadmap, threat model, glossary).
- Start with architecture, security, coding, and testing standards before implementing.
- If a change conflicts with a standard, update or add an ADR first.
- Keep updates concise, actionable, and consistent across related docs.

## Source of Truth Order
1. ADRs in `.adr/`
2. Repository operating model in `AGENTS.md`
3. Standards in `.ai/standards/`
4. General planning docs in `.ai/general/`

## Cross-References
- Agent roles and review model: [`AGENTS.md`](../AGENTS.md)
- Architectural decisions: [`.adr/`](../.adr/)
- Roadmap (canonical): [`./general/ROADMAP.md`](./general/ROADMAP.md)
- Threat model (canonical): [`./general/THREAT_MODEL.md`](./general/THREAT_MODEL.md)
- Definitions (canonical): [`./general/DEFINITIONS.md`](./general/DEFINITIONS.md)
