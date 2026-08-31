package main

import (
	"log"
	"os"

	"github.com/ethansaxenian/budgeting/migrate"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("migration direction required: up or down")
	}

	direction := os.Args[1]

	databaseURL := os.Getenv("DATABASE_URL")

	switch direction {
	case "up":
		if err := migrate.Up(databaseURL); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := migrate.Down(databaseURL); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Invalid migration direction: %s\n", direction)
	}
}
