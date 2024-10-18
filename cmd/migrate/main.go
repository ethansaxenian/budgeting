package main

import (
	"database/sql"
	"embed"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	direction := os.Args[1]

	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Invalid migration direction: %s\n", direction)
	}
}
