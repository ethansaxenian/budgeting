package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	DB *sql.DB
}

var (
	database = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USER")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func InitDB() (*DB, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
