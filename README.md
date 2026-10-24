# ai-order-management-backend

Go backend (Fiber, pgx) for AI order management.

## Project Structure

```
.
├── cmd/                  # app entrypoint (main package, built by air: ./cmd)
├── internal/             # private app code
│   ├── clients/          # external service clients
│   ├── config/           # app config loading (env)
│   ├── dto/              # request/response data transfer objects
│   ├── handlers/         # HTTP handlers
│   ├── routes/           # route registration
│   ├── services/         # business logic
│   └── store/            # database access
├── pkg/                  # reusable packages
│   ├── config/           # shared config helpers
│   ├── logger/           # logging
│   ├── middleware/       # Fiber middleware
│   ├── response/         # standard API response helpers
│   ├── security/         # auth / hashing / token helpers
│   └── utils/            # misc helpers
├── .air.toml             # hot reload config
├── .env.example          # env var template
├── go.mod
└── go.sum
```

## Setup

```sh
cp .env.example .env   # fill DB_USER / DB_PASSWORD
air                    # run with hot reload
```
