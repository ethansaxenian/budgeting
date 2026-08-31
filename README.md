# Budgeting

A simple web app to track monthly expenses.

## Requirements

- [Mise](https://mise.jdx.dev/)
- Docker
- Docker Compose
- Go 1.27.0 or later

Mise installs the project tools, including air, templ, sqlc, goose, and tailwind.

## Local development

Create a `mise.local.toml` file:

```toml
[env]
DB_USER = "postgres"
DB_PASSWORD = "password"
DB_HOST = "localhost"
DB_PORT = 5432
DB_NAME = "budgeting"
APP_PORT = 8000
```

Start the local development environment:

```sh
mise run dev
```

This starts PostgreSQL in Docker and runs the Go server with air, the templ proxy, the tailwind watcher, and the sqlc watcher locally. Open the app at:

```
http://localhost:8000
```

The templ development proxy listens on `APP_PORT` and forwards requests to the Go server on an internal port managed by Mise.

Useful commands:

```sh
mise run db                  # Start PostgreSQL
mise run db-down             # Stop PostgreSQL
mise run migrate             # Run pending migrations
mise run migrate-rollback   # Roll back the latest migration
mise run migrate-create example
```

The application runs migrations automatically when `RUN_MIGRATIONS` is truthy.
