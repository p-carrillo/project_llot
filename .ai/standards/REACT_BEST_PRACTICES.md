# React + TypeScript Best Practices (Internet-Backed)

Last reviewed: 2026-02-14

## Component Patterns
- Build small, composable components with explicit props.
- Keep render logic pure and side effects outside render.
- Lift shared state to the nearest common owner only when coordination is required.

## Hooks Rules
- Follow Rules of Hooks strictly (top-level only, React functions only).
- Keep custom hooks focused on one concern.
- Prefer event handlers and render-time derivation over unnecessary Effects.

## State Management
- Start with local state; lift state only when multiple components must coordinate.
- Avoid global state for transient UI concerns.
- Keep a single source of truth per state concern.

## Performance
- Use memoization (`memo`, `useMemo`) only when profiling shows value.
- Treat memoization as optimization, not correctness.
- Minimize prop churn and unnecessary Effect-triggered rerenders.

## Accessibility
- Use semantic HTML first.
- Use ARIA attributes correctly when semantic elements are insufficient.
- Preserve keyboard navigation and visible focus behavior.

## Testing
- Prefer user-centric component tests over implementation-detail tests.
- Validate behavior via rendered output and interactions.
- Keep tests resilient to internal refactors.

## Folder Conventions
- Organize by feature/domain first, then shared components.
- Keep API access in dedicated clients/hooks, not scattered in view components.
- Co-locate component tests with component/features where practical.

## TypeScript Practices
- Use `.tsx` for JSX files.
- Configure `tsconfig` JSX mode intentionally (`preserve`/`react-jsx` depending on build pipeline).
- Type component props explicitly and keep exported types stable.

## Do / Don’t Checklist
Do:
- Keep components pure and predictable.
- Use Hooks according to official rules.
- Prefer composition and clear boundaries.
- Profile before adding memoization.
- Design for keyboard and screen-reader compatibility.

Don’t:
- Mutate props/state during render.
- Use Effects for derivations that can happen during render.
- Add global state for local UI behavior.
- Memoize everything by default.
- Hide accessibility issues behind custom widgets without semantics.

## Sources
- Rules of React: https://react.dev/reference/rules
- Rules of Hooks: https://react.dev/reference/rules/rules-of-hooks
- Keeping Components Pure: https://react.dev/learn/keeping-components-pure
- You Might Not Need an Effect: https://react.dev/learn/you-might-not-need-an-effect
- `memo` reference: https://react.dev/reference/react/memo
- `useMemo` reference: https://react.dev/reference/react/useMemo
- Using TypeScript with React: https://react.dev/learn/typescript
- TypeScript JSX handbook: https://www.typescriptlang.org/docs/handbook/jsx
- React DOM common components and ARIA props: https://react.dev/reference/react-dom/components/common
- React Testing Library intro: https://testing-library.com/docs/react-testing-library/intro/
- WCAG 2.2 quick reference: https://www.w3.org/WAI/WCAG22/quickref/
