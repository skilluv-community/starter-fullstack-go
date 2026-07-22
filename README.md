# starter-fullstack-go

> A Skilluv starter — Go 1.23 (Gin + sqlx) backend, SvelteKit 5 frontend, PostgreSQL 18.

[![CI](https://github.com/skilluv-community/starter-fullstack-go/actions/workflows/ci.yml/badge.svg)](https://github.com/skilluv-community/starter-fullstack-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Skilluv](https://img.shields.io/badge/skilluv-community-emerald)](https://skilluv.io)

## English

### What this is

- **Backend**: Go 1.23 + Gin 1.10 + sqlx + pgx/v5 + `golang-migrate` for migrations + `slog` structured logs
- **Frontend**: SvelteKit 5 (runes) + Tailwind v4
- **Orchestration**: Docker Compose (postgres + backend + frontend)
- **Tests**: `go test` + Vitest + Playwright
- **CI**: GitHub Actions

### Quickstart

```bash
git clone git@github.com:skilluv-community/starter-fullstack-go.git
cd starter-fullstack-go
cp .env.example .env
docker compose up --build
```

- Frontend: <http://localhost:5173>
- Backend: <http://localhost:3001/health>

### What's inside

- `GET /health` — liveness
- `GET /api/hello?name=Ada` — greeting
- `GET/POST/DELETE /api/notes` — CRUD (sqlx + PostgreSQL, `golang-migrate` on boot)
- SvelteKit pages `/` and `/notes`

### Docs

- [`docs/en/getting-started.md`](./docs/en/getting-started.md)
- [`docs/en/architecture.md`](./docs/en/architecture.md)

---

## Français

Starter fullstack Go (Gin + sqlx) + SvelteKit. Voir [`docs/fr/getting-started.md`](./docs/fr/getting-started.md).

```bash
git clone git@github.com:skilluv-community/starter-fullstack-go.git
cd starter-fullstack-go
cp .env.example .env
docker compose up --build
```

---

## License

MIT — see [LICENSE](./LICENSE).
