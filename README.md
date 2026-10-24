# ai-order-management-backend

Go backend (Fiber, pgx) for AI order management.

## Project Structure

Microservices in one repo, laid out like `queue-backend`: one Go module per service, tied together by `go.work`. Services share only `pkg/`; no service imports another's `internal/`.

```
.
├── go.work               # gateway + pkg + services/*
├── gateway/              # public entry :8080 (skeleton); internal: clients, config, handlers, middleware, routes, services
├── pkg/                  # shared module: server, config, database, events, health, logger, middleware, outbox, response, security, utils
├── contracts/            # openapi/ (HTTP specs), events/ (event schemas)
├── services/
│   ├── user/             # users, roles, permissions :8081 (owns the DB schema)
│   ├── orders/           # order management :8082 (skeleton)
│   ├── product/          # product catalog :8084 (skeleton)
│   ├── supplier/         # suppliers :8085 (skeleton)
│   ├── ai/               # AI features :8083 (skeleton)
│   └── <service>/        # every service has:
│       ├── cmd/main.go
│       ├── internal/     # config, dto, handlers, routes, services, store (user also: schema)
│       ├── migrations/   # SQL migrations (user migrates from store/ instead)
│       ├── Dockerfile    # build context is the repo root
│       └── .air.toml  Makefile  .env.example  go.mod
├── deploy/local/         # docker-compose.yml (postgres + services)
├── Makefile              # make up | down | logs | run | dev | test s=<service>
└── .dockerignore
```

## Setup

```sh
cp gateway/.env.example gateway/.env
cp services/user/.env.example services/user/.env   # fill DB_USER / DB_PASSWORD
cp services/orders/.env.example services/orders/.env
cp services/ai/.env.example services/ai/.env
cp services/product/.env.example services/product/.env
cp services/supplier/.env.example services/supplier/.env
cp deploy/local/.env.example deploy/local/.env     # postgres superuser

make dev s=user        # hot reload one service (gateway | user | orders | product | supplier | ai)
make run s=orders      # run without air
make test s=user
make up                # all services in docker
```
