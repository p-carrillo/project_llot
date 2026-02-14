# UI/UX Standard

## Product Philosophy
- Build a Plausible-like dashboard: minimal, fast, low cognitive load.
- Prioritize clarity over configurability for MVP.

## Navigation
- Primary sections:
  - Overview
  - Traffic Quality (human vs bot)
  - Sessions
  - Hosts/Sites
  - Settings

## Default Filters
- Last 24h as default time window.
- Host/site selector defaults to all.
- Bot-classification filter defaults to all with quick toggles.

## Performance
- Avoid blocking renders on large payloads.
- Prefer server-side aggregation over client-heavy processing.
- Keep initial load small; defer non-critical data.

## Accessibility Basics
- Semantic HTML and keyboard-first navigation.
- Visible focus states and logical tab order.
- Sufficient color contrast and clear labels.
- Use ARIA only when semantic elements are insufficient.

## Component-Based Design
- Boundaries: page -> feature section -> reusable UI primitives.
- Naming: `FeatureThing` for feature components, `UiThing` for shared primitives.
- State guidance:
  - local UI state in components
  - shared page state in feature container hooks
  - server state via dedicated API client hooks
