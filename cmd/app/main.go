package main

import (
	"log"

	"github.com/ethansaxenian/budgeting/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
