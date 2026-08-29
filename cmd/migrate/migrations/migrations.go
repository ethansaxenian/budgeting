package migrations

import (
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var files embed.FS

func Up(databaseURL string) error {
	return run(databaseURL, func(db *sql.DB) error {
		return goose.Up(db, ".")
	})
}

func Down(databaseURL string) error {
	return run(databaseURL, func(db *sql.DB) error {
		return goose.Down(db, ".")
	})
}

func run(databaseURL string, migrate func(*sql.DB) error) error {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(files)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return migrate(db)
}
