package main

import (
	"log"
	"os"

	"github.com/ethansaxenian/budgeting/cmd/migrate/migrations"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("migration direction required: up or down")
	}

	direction := os.Args[1]

	switch direction {
	case "up":
		if err := migrations.Up(os.Getenv("DATABASE_URL")); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := migrations.Down(os.Getenv("DATABASE_URL")); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Invalid migration direction: %s\n", direction)
	}
}
