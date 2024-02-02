package main

import (
	"embed"
	"log"

	"github.com/ethansaxenian/budgeting/database"
	"github.com/pressly/goose/v3"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.Fatal(err)
	}
}
