package main

import (
	"log"

	"github.com/ethansaxenian/budgeting/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server, close, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
