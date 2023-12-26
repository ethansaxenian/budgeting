package db

import (
	"github.com/ethansaxenian/budgeting/types"
)

func GetTransactions() ([]types.Transaction, error) {
	rows, err := DB.Query("SELECT * FROM transactions")
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

func CreateTransaction(tr types.TransactionCreate) (int, error) {
	res, err := DB.Exec("INSERT INTO transactions (description, amount, date, category, type, month_id) VALUES ($1, $2, $3, $4, $5, $6)",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		tr.Type,
		tr.MonthID,
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

func UpdateTransaction(id int, tr types.TransactionUpdate) (int, error) {
	res, err := DB.Exec("UPDATE transactions SET description=$1, amount=$2, date=$3, category=$4, type=$5, month_id=$6 WHERE id=$7",
		tr.Description,
		tr.Amount,
		tr.Date,
		tr.Category,
		tr.Type,
		tr.MonthID,
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

func DeleteTransaction(id int) (int, error) {
	res, err := DB.Exec("DELETE FROM transactions WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowCount), nil
}

func GetTransactionsByMonthID(monthID int) ([]types.Transaction, error) {
	rows, err := DB.Query("SELECT * FROM transactions WHERE month_id=$1", monthID)
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
			&tr.MonthID,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil

}
