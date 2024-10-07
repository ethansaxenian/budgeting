package main

import (
	"embed"
	"log"
	"os"

	"github.com/ethansaxenian/budgeting/database"
	"github.com/pressly/goose/v3"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	direction := os.Args[1]

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := goose.Up(db.DB, "migrations"); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := goose.Down(db.DB, "migrations"); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Invalid migration direction: %s\n", direction)
	}
}
