package main

import (
	"fmt"
	"log"

	"github.com/ethansaxenian/budgeting/db"
)

func main() {
	transactions, err := db.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transactions)
}
