package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"localhost",
		os.Getenv("DB_PORT"),
		"postgres",
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
