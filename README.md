# ai-order-management-backend

Go backend (Fiber, pgx) for AI order management.

## Project Structure

Microservices in one repo, laid out like `queue-backend`: one Go module per service, tied together by `go.work`. Services share only `pkg/`; no service imports another's `internal/`.

```
.
├── go.work               # pkg + services/*
├── pkg/                  # shared module (server bootstrap)
├── services/
│   ├── user/             # users, roles, permissions :8081 (owns the DB schema)
│   ├── orders/           # order management :8082 (skeleton)
│   ├── product/          # product catalog :8084 (skeleton)
│   ├── supplier/         # suppliers :8085 (skeleton)
│   ├── ai/               # AI features :8083 (skeleton)
│   └── <service>/        # every service has:
│       ├── cmd/main.go
│       ├── internal/     # per-service code (user: schema, services, store)
│       ├── Dockerfile    # build context is the repo root
│       └── .air.toml  Makefile  .env.example  go.mod
├── deploy/local/         # docker-compose.yml (postgres + services)
├── Makefile              # make up | down | logs | run | dev | test s=<service>
└── .dockerignore
```

## Setup

```sh
cp services/user/.env.example services/user/.env   # fill DB_USER / DB_PASSWORD
cp services/orders/.env.example services/orders/.env
cp services/ai/.env.example services/ai/.env
cp services/product/.env.example services/product/.env
cp services/supplier/.env.example services/supplier/.env
cp deploy/local/.env.example deploy/local/.env     # postgres superuser

make dev s=user        # hot reload one service (user | orders | product | supplier | ai)
make run s=orders      # run without air
make test s=user
make up                # all services in docker
```
