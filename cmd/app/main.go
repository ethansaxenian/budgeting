package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ethansaxenian/budgeting/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	databaseURL := os.Getenv("DB_URL")
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal("No APP_PORT provided, exiting...")
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
