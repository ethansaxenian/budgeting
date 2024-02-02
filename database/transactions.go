package database

import (
	"errors"

	"github.com/ethansaxenian/budgeting/types"
)

func (db *DB) GetTransactions() ([]types.Transaction, error) {
	rows, err := db.DB.Query("SELECT id, date, amount, description, category, type FROM transactions")
	if err != nil {
		return nil, errors.New("error retrieving transactions")
	}

	var transactions []types.Transaction

	for rows.Next() {
		tr := types.Transaction{}
		if err = rows.Scan(
			&tr.ID,
			&tr.Date,
			&tr.Amount,
			&tr.Description,
			&tr.Category,
			&tr.Type,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}

func (db *DB) GetTransactionByID(id int) (types.Transaction, error) {
	row := db.DB.QueryRow("SELECT * FROM transactions WHERE id=$1", id)

	tr := types.Transaction{}
	if err := row.Scan(
		&tr.ID,
		&tr.Date,
		&tr.Amount,
		&tr.Description,
		&tr.Category,
		&tr.Type,
	); err != nil {
		return types.Transaction{}, err
	}

	return tr, nil
}

func (db *DB) CreateTransaction(tr types.TransactionCreate) error {
	_, err := db.DB.Exec("INSERT INTO transactions (description, amount, date, category, type) VALUES ($1, $2, $3, $4, $5)",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		tr.Type,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTransaction(id int, tr types.TransactionUpdate) error {
	_, err := db.DB.Exec("UPDATE transactions SET description=$1, amount=$2, date=$3, category=$4 WHERE id=$5",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteTransaction(id int) error {
	_, err := db.DB.Exec("DELETE FROM transactions WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
