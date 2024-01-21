.PHONY: all run run-exe tailwind templ build clean dev

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

clean:
	@rm -rfv ./bin
	@rm -rfv ./assets/dist
	@rm -rfv ./components/**/*_templ.go
