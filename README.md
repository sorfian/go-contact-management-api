# Go Todo List API

A RESTful Todo List API built with Go, Fiber, GORM (MySQL), and Google Wire for dependency injection. It provides user authentication plus CRUD for contacts and nested addresses. An OpenAPI specification is included.

Note: This README consolidates and updates information from QUICKSTART.md, DEPLOYMENT.md, and integration_test.md. See those files for deeper, task‑specific details.

## Overview

- Language: Go (Go modules)
- Web framework: Fiber v2
- ORM/DB: GORM with MySQL driver
- DI codegen: Google Wire
- Validation: go-playground/validator
- Testing: Go test + Testify
- API spec: OpenAPI 3.1 (`apispec.yaml`)

Base API path: `/api`

Key resources and endpoints (see router and OpenAPI for full details):
- Users: `POST /api/users/register`, `POST /api/users/login`, `GET|PATCH|DELETE /api/users/current`
- Contacts: `POST|GET /api/contacts`, `GET|PATCH|DELETE /api/contacts/:contactId`
- Addresses (nested under contacts): `POST|GET /api/contacts/:contactId/addresses`, `GET|PATCH|DELETE /api/contacts/:contactId/addresses/:addressId`

## Requirements

- Go toolchain (version declared in `go.mod`: 1.25.3)
- MySQL server (or compatible, e.g., MariaDB)
- Make (optional but recommended for convenience)

Optional tools:
- curl or a REST client for testing endpoints
- A migration tool (see Migrations section below)

## Project Stack and Entry Points

- Package manager: Go modules (`go.mod`, `go.sum`)
- Application entry point: `main.go`
- Dependency graph wired via Google Wire:
  - Build-time provider set files: `wire.go`, `wire_gen.go` (generated)
  - Test DI: `test/wire.go`, `test/wire_gen.go` (generated)
- HTTP routes are defined in `app/router.go`; Fiber app setup in `app_setup.go`.

## Getting Started

### 1) Clone and prepare

```bash
# clone your fork or the repository
# cd go-todo-list-app

# Install dependencies
make install       # or: go mod download && go mod tidy

# Generate DI code (Wire)
make wire          # or: go run github.com/google/wire/cmd/wire && (cd test && go run github.com/google/wire/cmd/wire)

# Copy and edit environment file
cp .env.example .env   # If present; otherwise create .env based on the variables documented below
```

### 2) Database

- Ensure a MySQL instance is running and accessible based on your `.env` configuration.
- Create the database schema if it does not exist (e.g., `CREATE DATABASE go_todo_list;`).
- Migrations are provided under `db/migrations`. See Migrations section.

### 3) Run in development

Using Make:
```bash
make dev           # runs: go run .
```

Without Make:
```bash
go run .
```

Server starts on the configured port (default `3000`).

### 4) Build and run a binary

Using Make:
```bash
make build         # produces bin/app (or bin/app.exe on Windows)
make run           # builds and runs
```

Manual:
```bash
go build -o bin/app .
./bin/app          # Linux/macOS
# .\bin\app.exe    # Windows
```

## Environment Variables

The application reads configuration from environment variables (optionally via `.env`, loaded by `github.com/joho/godotenv`). Defaults are shown in parentheses.

Application:
- `APP_ENV` ("development")
- `APP_PORT` ("3000")
- `LOG_LEVEL` ("info")

Database:
- `DB_HOST` ("localhost")
- `DB_PORT` ("3306")
- `DB_USER` ("root")
- `DB_PASSWORD` ("")
- `DB_NAME` ("go_todo_list")
- `DB_MAX_IDLE_CONNS` (10)
- `DB_MAX_OPEN_CONNS` (100)
- `DB_CONN_MAX_LIFETIME` ("30m")
- `DB_CONN_MAX_IDLE_TIME` ("10m")

See `app/config.go` for authoritative defaults and DSN construction; database connection is initialized in `app/database.go`.

## Scripts and Common Commands

Make targets (from `Makefile`):
- `make install` — Install/tidy dependencies
- `make wire` — Generate Wire DI code (app and test)
- `make dev` — Run in development (`go run .`)
- `make build` — Build binary to `bin/app`
- `make run` — Build and run
- `make test` — Run all tests (`go test -v ./...`)
- `make test-unit` — Run unit tests (`-short`)
- `make test-integration` — Run tests under `./test`
- `make clean` — Remove build artifacts

Equivalent Go commands are available for environments without Make (see QUICKSTART.md).

## Tests

- Unit tests: run with `make test-unit` or `go test -v -short ./...`.
- Integration tests: see `integration_test.md` and use `make test-integration` or `cd test && go test -v`.
- All tests: `make test`.

Test wiring uses Google Wire in `test/wire.go` (generated code in `test/wire_gen.go`).

## API Documentation

- OpenAPI spec: `apispec.yaml` (OpenAPI 3.1). It defines endpoints and example schemas.
- Servers listed in spec include `http://localhost:3000/api` and `https://todo.signal.id/api`.
  - TODO: Confirm the production base URL and deployment domain.

You can use `curl` or a REST client (Insomnia/Postman) to hit endpoints. Example:
```bash
# Register
target=http://localhost:3000/api
curl -sS -X POST "$target/users/register" \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","username":"alice","password":"secret"}'
```

## Migrations

SQL migration files are under `db/migrations`:
- `*_create_table_users.up.sql` / `.down.sql`
- `*_create_table_contacts.up.sql` / `.down.sql`
- `*_create_table_addresses.up.sql` / `.down.sql`

A dedicated migration tool is not bundled/configured in this repository.
- You can apply these SQL files manually using your MySQL client.
- Or integrate a tool like `golang-migrate` or `goose` in your environment.
- TODO: Document and/or add a unified migration command.

## Project Structure

Top-level overview:
```
.
├─ app/                 # App config, DB connection, HTTP router, Wire providers
├─ controller/          # HTTP controllers (interfaces + implementations)
├─ middleware/          # Auth middleware, etc.
├─ model/               # Domain and web (request/response) models
│  ├─ domain/
│  └─ web/
├─ repository/          # Data access layer (interfaces + implementations)
├─ service/             # Business logic services
├─ db/migrations/       # SQL migration files
├─ test/                # Integration tests and test DI wiring
├─ apispec.yaml         # OpenAPI 3.1 specification
├─ Makefile             # Developer convenience commands
├─ main.go              # Program entry point
├─ wire.go              # Wire build description (generate wire_gen.go)
├─ QUICKSTART.md        # Quick start guide
├─ DEPLOYMENT.md        # Deployment guide
└─ README.md            # This file
```

## Deployment

See `DEPLOYMENT.md` for a full deployment guide, including environment preparation and running the compiled binary. Typical steps:
- Build: `make build`
- Export environment variables for production
- Run: `./bin/app` (Linux/macOS) or `./bin/app.exe` (Windows)

## License

No license file detected in this repository.
- TODO: Add a `LICENSE` file (e.g., MIT/Apache-2.0) and state the chosen license here.

## Notes & Caveats

- Google Wire must be run before building when provider sets change: `make wire`.
- Default port is `3000`. Change via `APP_PORT`.
- Authentication middleware protects `GET|PATCH|DELETE /api/users/current` and all `/api/contacts` routes; include proper auth headers per your implementation.
- Windows users: if `make` is unavailable, use the raw `go` commands shown above.
