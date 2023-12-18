package db

import (
	"log"

	"github.com/ethansaxenian/budgeting/types"
)

func GetTransactions() ([]types.Transaction, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		log.Fatal(err)
	}

	var transactions []types.Transaction

	for rows.Next() {
		tr := types.Transaction{}
		if err = rows.Scan(
			&tr.ID,
			&tr.Description,
			&tr.Amount,
			&tr.Date,
			&tr.Category,
			&tr.Type,
			&tr.MonthID,
		); err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}
