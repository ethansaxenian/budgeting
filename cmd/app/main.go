package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/migrate"
	"github.com/ethansaxenian/budgeting/server"
	_ "github.com/joho/godotenv/autoload"
)

func envBool(key string, defaultValue bool) (bool, error) {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return defaultValue, nil
	}

	return strconv.ParseBool(value)
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("No DATABASE_URL provided, exiting...")
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("No APP_PORT provided, exiting...")
	}

	runMigrations, err := envBool("RUN_MIGRATIONS", false)
	if err != nil {
		log.Fatalf("invalid RUN_MIGRATIONS: %v", err)
	}

	if runMigrations {
		if err := migrate.Up(databaseURL); err != nil {
			log.Fatalf("running database migrations: %v", err)
		}
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
