include .env

.PHONY: start stop logs migrate migrate-create migrate-rollback sql

start:
	docker-compose up -d --wait

stop:
	docker-compose down

logs:
	docker compose logs -f

rebuild:
	docker-compose up --build -d --wait

shell:
	docker exec -it server sh

migrate:
	docker exec -i server sh -c "go run cmd/migrate/main.go up"

migrate-create:
	docker exec -i server sh -c "go tool goose create $(name) sql"

migrate-rollback:
	docker exec -i server sh -c "go run cmd/migrate/main.go down"

backup:
	docker exec -i db sh -c "pg_dump --data-only --exclude-table=goose_db_version ${DATABASE_URL}" > "backups/backup_$(shell date --iso="seconds").sql"

backup-restore:
	docker exec -i db sh -c "psql -U ${DB_USER} -d ${DB_NAME}" < "$(backup_file)"

sql:
	docker exec -i server sh -c "go tool sqlc generate"
