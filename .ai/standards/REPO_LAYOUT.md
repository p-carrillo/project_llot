# Repository Layout Standard

## Canonical Layout

```text
.
├── .adr/
├── .ai/
│   ├── general/
│   └── standards/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── domain/
│   │   ├── application/
│   │   ├── ports/
│   │   └── adapters/
│   └── test/
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   ├── features/
│   │   ├── hooks/
│   │   └── api/
│   └── test/
├── packaging/
│   ├── debian/
│   ├── systemd/
│   └── nginx/
├── docs/
└── scripts/
```

## Ownership Map
- `backend/`: Backend / Daemon (primary), Security and QA (required reviewers).
- `frontend/`: Product / UX + Frontend owners (primary), QA and Security (review).
- `packaging/`: Packaging / Release + Nginx / Infra (primary), Security and QA (review).
- `.ai/` and `.adr/`: cross-team architecture/governance ownership.
