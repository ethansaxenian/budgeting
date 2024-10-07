include .env

.PHONY: start stop logs migrate migrate-create migrate-rollback

start:
	docker-compose up -d --wait

stop:
	docker-compose down

logs:
	docker logs budgeting-web -f

rebuild:
	docker-compose up --build -d --wait

shell:
	docker exec -it budgeting-web sh

migrate:
	docker exec -i budgeting-web sh -c "go run cmd/migrate/main.go up"

migrate-create:
	docker exec -i budgeting-web sh -c "goose create $(name) sql"

migrate-rollback:
	docker exec -i budgeting-web sh -c "go run cmd/migrate/main.go down"

backup:
	docker exec -i budgeting-db sh -c "pg_dump --data-only --exclude-table=goose_db_version postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" > "backups/backup_$(shell date --iso="seconds").sql"

backup-restore:
	docker exec -i budgeting-db sh -c "psql -U ${DB_USER} -d ${DB_NAME}" < "$(backup_file)"
