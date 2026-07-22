# Architecture — starter-fullstack-go

## Choix opinionated

### 1. Gin plutôt que stdlib

Assez opinionated (binding, groups, middleware) pour que les débutants produisent du code correct vite. Si Gin devient trop, `net/http` + `chi` est la suite.

### 2. sqlx plutôt qu'ORM

SQL brut proche du Go, tags struct pour le scan. Pas de N+1 caché. Utiliser `sqlc` pour vérification à la compilation.

### 3. pgx/v5 comme driver

Driver Postgres moderne. Utilisé via sa couche `stdlib` compatible avec `database/sql`.

### 4. `golang-migrate` au boot

Le serveur exécute `migrate up` au démarrage. Fichiers sous `backend/migrations/000X_*.up.sql` / `.down.sql`.

### 5. `slog` pour les logs

Logs structurés standard (Go 1.21+). Zéro dépendance.

### 6. `caarlos0/env/v11` pour la config

Env parsing par tags struct. Simple, testable, sans viper.

## Hors scope

- Auth — middleware + JWT lib de ton choix.
- Rate limiting — `ulule/limiter` ou `gin-contrib/limits`.
- Observabilité — voir `starter-devops`.
