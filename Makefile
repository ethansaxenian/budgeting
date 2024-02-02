include .env

.PHONY: run run-exe tailwind templ build clean dev backup help

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  run        - Run the application"
	@echo "  run-exe    - Run the compiled application"
	@echo "  dev        - Run the application with hot reloading"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean the application"
	@echo "  tailwind   - Compile tailwindcss"
	@echo "  templ      - Generate templates"
	@echo "  backup     - Backup the database"
	@echo "  help       - Show this help message"

run: tailwind templ
	@go run cmd/main.go

dev:
	@air -c .air.toml

run-exe: build
	@./bin/main

tailwind:
	@tailwindcss -i ./assets/main.css -o ./assets/dist/tailwind.css

templ:
	@templ generate

build: tailwind templ
	@go build -o ./bin/main cmd/main.go

backup:
	@pg_dump postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME} -a > "backups/backup_$(shell date --iso="seconds").sql"

clean:
	@rm -rfv ./bin
	@rm -rfv ./assets/dist
	@rm -rfv ./components/**/*_templ.go
