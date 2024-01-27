include .env

.PHONY: all run run-exe tailwind templ build clean dev backup

all: build run-exe

run: tailwind templ
	@go run cmd/main.go

dev:
	@air -c .air.toml

run-exe:
	@./bin/main

tailwind:
	@tailwindcss -i ./assets/main.css -o ./assets/dist/tailwind.css

templ:
	@templ generate

build: tailwind templ
	@go build -o ./bin/main cmd/main.go

backup:
	@pg_dump postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE} -a > "backups/backup_$(shell date --iso="seconds").sql"

clean:
	@rm -rfv ./bin
	@rm -rfv ./assets/dist
	@rm -rfv ./components/**/*_templ.go
