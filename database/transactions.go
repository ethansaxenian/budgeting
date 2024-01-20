package database

import (
	"errors"

	"github.com/ethansaxenian/budgeting/types"
)

func (db *DB) GetTransactions() ([]types.Transaction, error) {
	rows, err := db.db.Query("SELECT id, date, amount, description, category, type FROM transactions")
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
	row := db.db.QueryRow("SELECT * FROM transactions WHERE id=$1", id)

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

func (db *DB) CreateTransaction(tr types.TransactionCreate) (int, error) {
	res, err := db.db.Exec("INSERT INTO transactions (description, amount, date, category, type) VALUES ($1, $2, $3, $4, $5, $6)",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		tr.Type,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *DB) UpdateTransaction(id int, tr types.TransactionUpdate) (int, error) {
	res, err := db.db.Exec("UPDATE transactions SET description=$1, amount=$2, date=$3, category=$4 WHERE id=$5",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		id,
	)
	if err != nil {
		return 0, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowCount), nil
}

func (db *DB) DeleteTransaction(id int) (int, error) {
	res, err := db.db.Exec("DELETE FROM transactions WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowCount), nil
}

func (db *DB) GetTransactionsByMonthID(monthID int) ([]types.Transaction, error) {
	rows, err := db.db.Query("SELECT * FROM transactions WHERE month_id=$1", monthID)
	if err != nil {
		return nil, err
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
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil

}
