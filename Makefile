include .env

.PHONY: run run-exe tailwind templ build clean dev backup help db migrate migrate-create stop

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run                       - Run the application"
	@echo "  run-exe                   - Run the compiled application"
	@echo "  dev                       - Run the application with hot reloading"
	@echo "  build                     - Build the application"
	@echo "  clean                     - Clean the application"
	@echo "  tailwind                  - Compile tailwindcss"
	@echo "  templ                     - Generate templates"
	@echo "  db                        - Start the database"
	@echo "  migrate                   - Migrate the database"
	@echo "  migrate-create name=NAME  - Create a new migration with the given name"
	@echo "  stop                      - Stop the database"
	@echo "  backup                    - Backup the database"
	@echo "  help                      - Show this help message"

run: db tailwind templ
	@go run cmd/app/main.go

dev:
	@air -c .air.toml

run-exe: db build
	@./bin/main

tailwind:
	@tailwindcss -i ./assets/main.css -o ./assets/dist/tailwind.css

templ:
	@templ generate

build: tailwind templ
	@go build -o ./bin/main cmd/app/main.go

db:
	@docker-compose up -d

migrate: db
	@go run cmd/migrate/main.go

migrate-create: db
	@GOOSE_DBSTRING="user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} host=${DB_HOST} sslmode=disable" GOOSE_DRIVER=postgres goose -dir cmd/migrate/migrations create $(name) sql

stop:
	@docker-compose down

backup: db
	@pg_dump postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME} -a > "backups/backup_$(shell date --iso="seconds").sql"

clean:
	@rm -rfv ./bin
	@rm -rfv ./assets/dist
	@rm -rfv ./components/**/*_templ.go
