package main

import (
	"log"
	"net/http"

	"github.com/ethansaxenian/budgeting/db"
	"github.com/ethansaxenian/budgeting/router"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := router.InitRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
