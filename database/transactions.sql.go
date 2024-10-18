// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transactions.sql

package database

import (
	"context"
	"time"

	"github.com/ethansaxenian/budgeting/types"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (description, amount, date, category, transaction_type)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, date, amount, description, category, transaction_type, created_at, updated_at
`

type CreateTransactionParams struct {
	Description     string
	Amount          float64
	Date            time.Time
	Category        types.Category
	TransactionType types.TransactionType
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.Description,
		arg.Amount,
		arg.Date,
		arg.Category,
		arg.TransactionType,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Amount,
		&i.Description,
		&i.Category,
		&i.TransactionType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const getTransactionsByMonthIDAndType = `-- name: GetTransactionsByMonthIDAndType :many
SELECT
    t.id,
    t.date,
    t.amount,
    t.description,
    t.category,
    t.transaction_type,
    t.created_at,
    t.updated_at
FROM
    transactions t
JOIN
    months m ON t.date BETWEEN
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')))
        AND
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')) + INTERVAL '1 month' - INTERVAL '1 day')
WHERE
    m.id = $1 AND t.transaction_type = $2
`

type GetTransactionsByMonthIDAndTypeParams struct {
	ID              int
	TransactionType types.TransactionType
}

func (q *Queries) GetTransactionsByMonthIDAndType(ctx context.Context, arg GetTransactionsByMonthIDAndTypeParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactionsByMonthIDAndType, arg.ID, arg.TransactionType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Amount,
			&i.Description,
			&i.Category,
			&i.TransactionType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions
SET description = $1, amount = $2, date = $3, category = $4
WHERE id = $5
RETURNING id, date, amount, description, category, transaction_type, created_at, updated_at
`

type UpdateTransactionParams struct {
	Description string
	Amount      float64
	Date        time.Time
	Category    types.Category
	ID          int
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, updateTransaction,
		arg.Description,
		arg.Amount,
		arg.Date,
		arg.Category,
		arg.ID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Date,
		&i.Amount,
		&i.Description,
		&i.Category,
		&i.TransactionType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
