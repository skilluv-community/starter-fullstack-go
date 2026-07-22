# Architecture — starter-fullstack-go

## Opinionated choices

### 1. Gin over stdlib

For a starter, Gin is opinionated enough (binding, groups, middleware) that beginners can produce reasonable code fast. If you outgrow Gin, `net/http` + `chi` is the natural next step.

### 2. sqlx over an ORM

Raw SQL kept close to Go, with struct tags for scanning. No hidden N+1s, no query DSL to learn. Use `sqlc` if you want compile-time query checking.

### 3. pgx/v5 as the driver

Modern, feature-complete Postgres driver. Used through its `stdlib` compatibility layer so `sqlx` sees it as a normal `database/sql` driver.

### 4. `golang-migrate` on boot

The server runs `migrate up` at startup. Migration files live under `backend/migrations/000X_*.up.sql` / `.down.sql`.

### 5. `slog` for logs

Standard library structured logging (Go 1.21+). No third-party log package.

### 6. `caarlos0/env/v11` for config

Struct-tagged env parsing. Simple, testable, no viper heaviness.

## What's out of scope

- Auth — add a middleware and a JWT lib of your choice.
- Rate limiting — `ulule/limiter` or `gin-contrib/limits`.
- Observability — see `starter-devops`.
