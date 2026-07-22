# Démarrage — starter-fullstack-go

## Prérequis

- Docker + Docker Compose 2.24+
- (Optionnel) Go 1.23+, Node 22/24 LTS, PostgreSQL 18

## Premier lancement

```bash
git clone git@github.com:skilluv-community/starter-fullstack-go.git
cd starter-fullstack-go
cp .env.example .env
docker compose up --build
```

- Frontend : <http://localhost:5173>
- Health backend : <http://localhost:3001/health>

## Scripts

- `make test`, `make lint`, `make build`

## Déploiement

En prod : `GIN_MODE=release`, restreindre `CORS_ORIGIN`, vrai `POSTGRES_PASSWORD`.
