.PHONY: run
run:
	@tailwindcss -i ./assets/main.css -o ./assets/dist/tailwind.css
	@templ generate
	@go run cmd/main.go
