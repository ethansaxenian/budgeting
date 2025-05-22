include .env

DATABASE_URL = postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/postgres

.PHONY: start
start:
	docker compose up -d --wait

.PHONY: stop
stop:
	docker compose down

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: rebuild
rebuild:
	docker compose up --build -d --wait

.PHONY: shell
shell:
	docker exec -it server sh

.PHONY: migrate
migrate:
	docker exec -i server sh -c "go run cmd/migrate/main.go up"

.PHONY: migrate-create
migrate-create:
	GOOSE_MIGRATION_DIR=cmd/migrate/migrations go tool goose create $(name) sql

.PHONY: migrate-rollback
migrate-rollback:
	docker exec -i server sh -c "go run cmd/migrate/main.go down"

.PHONY: backup
backup:
	docker exec -i db sh -c "pg_dump --data-only --exclude-table=goose_db_version ${DATABASE_URL}" > "backups/backup_$(shell date --iso="seconds").sql"

.PHONY: backup-restore
backup-restore:
	docker exec -i db sh -c "psql -U ${DB_USER} -d postgres" < "$(backup_file)"

.PHONY: sql
sql:
	go tool sqlc generate

.PHONY: lint
lint:
	go tool golangci-lint run
