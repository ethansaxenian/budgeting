package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/cmd/migrate/migrations"
	"github.com/ethansaxenian/budgeting/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("No DATABASE_URL provided, exiting...")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("No APP_PORT provided, exiting...")
	}

	if err := migrations.Up(databaseURL); err != nil {
		log.Fatalf("running database migrations: %v", err)
	}

	server, err := server.NewServer(port, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
