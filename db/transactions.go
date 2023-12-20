package db

import (
	"log"

	"github.com/ethansaxenian/budgeting/types"
)

func GetTransactions() ([]types.Transaction, error) {
	rows, err := DB.Query("SELECT * FROM transactions")
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
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}

func GetTransactionByID(id int) (types.Transaction, error) {
	row := DB.QueryRow("SELECT * FROM transactions WHERE id=$1", id)

	tr := types.Transaction{}
	if err := row.Scan(
		&tr.ID,
		&tr.Description,
		&tr.Amount,
		&tr.Date,
		&tr.Category,
		&tr.Type,
		&tr.MonthID,
	); err != nil {
		return types.Transaction{}, err
	}

	return tr, nil
}
