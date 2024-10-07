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

migrate:
	docker exec -i budgeting-web sh -c "go run cmd/migrate/main.go"

migrate-create:
	docker exec -i budgeting-web sh -c "GOOSE_DBSTRING=\"user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} host=${DB_HOST} sslmode=disable\" GOOSE_DRIVER=postgres goose -dir cmd/migrate/migrations create $(name) sql"

migrate-rollback:
	docker exec -i budgeting-web sh -c "GOOSE_DBSTRING=\"user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} host=${DB_HOST} sslmode=disable\" GOOSE_DRIVER=postgres goose -dir cmd/migrate/migrations down"

backup:
	docker exec -i budgeting-db sh -c "pg_dump --data-only postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" > "backups/backup_$(shell date --iso="seconds").sql"
