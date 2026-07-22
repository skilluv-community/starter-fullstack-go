# Getting started — starter-fullstack-go

## Prerequisites

- Docker + Docker Compose 2.24+
- (Optional) Go 1.23+, Node 22/24 LTS, PostgreSQL 18

## First run

```bash
git clone git@github.com:skilluv-community/starter-fullstack-go.git
cd starter-fullstack-go
cp .env.example .env
docker compose up --build
```

- Frontend at <http://localhost:5173>
- Backend health at <http://localhost:3001/health>

## Scripts

- `make test` — run backend + frontend tests
- `make lint` — `go vet`, `gofmt`, ESLint, `svelte-check`
- `make build` — compile the Go binary and the frontend

## Deploying

Point Coolify at the repo. In production:

- Set `GIN_MODE=release`
- Restrict `CORS_ORIGIN`
- Set a strong `POSTGRES_PASSWORD`
