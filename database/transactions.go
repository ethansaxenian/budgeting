package database

import (
	"fmt"

	"github.com/ethansaxenian/budgeting/types"
)

func (db *DB) GetTransactionsByMonthIDAndType(monthID int, transactionType types.TransactionType) ([]types.Transaction, error) {
	month, err := db.GetMonthByID(monthID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving month with id %d", monthID)
	}

	startDate, endDate := month.StartEndDates()

	rows, err := db.DB.Query(`
			SELECT id, date, amount, description, category, transaction_type
			FROM transactions
			WHERE date BETWEEN $1 AND $2 AND transaction_type = $3
		`,
		startDate,
		endDate,
		transactionType,
	)

	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions for %s %d: %v", month.Month.String(), month.Year, err)
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

func (db *DB) GetTransactionsByMonthIDAndCategoryAndType(monthID int, category types.Category, transactionType types.TransactionType) ([]types.Transaction, error) {
	month, err := db.GetMonthByID(monthID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving month with id %d", monthID)
	}

	startDate, endDate := month.StartEndDates()

	rows, err := db.DB.Query(`
			SELECT id, date, amount, description, category, transaction_type
			FROM transactions
			WHERE date BETWEEN $1 AND $2 AND transaction_type = $3 AND category = $4
		`,
		startDate,
		endDate,
		transactionType,
		category,
	)

	if err != nil {
		return nil, fmt.Errorf("error retrieving %s transactions for %s %d: %v", category, month.Month.String(), month.Year, err)
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
	_, err := db.DB.Exec("INSERT INTO transactions (description, amount, date, category, transaction_type) VALUES ($1, $2, $3, $4, $5)",
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
