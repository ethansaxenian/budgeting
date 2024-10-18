// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: budgets.sql

package database

import (
	"context"

	"github.com/ethansaxenian/budgeting/types"
)

const createNewBudgetsForMonth = `-- name: CreateNewBudgetsForMonth :many
INSERT INTO budgets (month_id, category, amount, transaction_type)
VALUES
($1, 'food', 0, 'expense'),
($1, 'food', 0, 'income'),
($1, 'gifts', 0, 'expense'),
($1, 'gifts', 0, 'income'),
($1, 'home', 0, 'expense'),
($1, 'home', 0, 'income'),
($1, 'medical', 0, 'expense'),
($1, 'medical', 0, 'income'),
($1, 'transportation', 0, 'expense'),
($1, 'transportation', 0, 'income'),
($1, 'personal', 0, 'expense'),
($1, 'personal', 0, 'income'),
($1, 'savings', 0, 'expense'),
($1, 'savings', 0, 'income'),
($1, 'utilities', 0, 'expense'),
($1, 'utilities', 0, 'income'),
($1, 'travel', 0, 'expense'),
($1, 'travel', 0, 'income'),
($1, 'other', 0, 'expense'),
($1, 'other', 0, 'income'),
($1, 'paycheck', 0, 'expense'),
($1, 'paycheck', 0, 'income'),
($1, 'bonus', 0, 'expense'),
($1, 'bonus', 0, 'income'),
($1, 'interest', 0, 'expense'),
($1, 'interest', 0, 'income'),
($1, 'cashback', 0, 'income'),
($1, 'cashback', 0, 'expense')
RETURNING id, month_id, category, amount, transaction_type, created_at, updated_at
`

func (q *Queries) CreateNewBudgetsForMonth(ctx context.Context, monthID int) ([]Budget, error) {
	rows, err := q.db.QueryContext(ctx, createNewBudgetsForMonth, monthID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Budget
	for rows.Next() {
		var i Budget
		if err := rows.Scan(
			&i.ID,
			&i.MonthID,
			&i.Category,
			&i.Amount,
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

const getBudgetByID = `-- name: GetBudgetByID :one
SELECT id, month_id, category, amount, transaction_type, created_at, updated_at FROM budgets WHERE id = $1
`

func (q *Queries) GetBudgetByID(ctx context.Context, id int) (Budget, error) {
	row := q.db.QueryRowContext(ctx, getBudgetByID, id)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.MonthID,
		&i.Category,
		&i.Amount,
		&i.TransactionType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBudgetItemsForMonthIDByTransactionType = `-- name: GetBudgetItemsForMonthIDByTransactionType :many
SELECT
    b.id AS budget_id,
    b.category,
    b.transaction_type AS type,
    b.amount AS planned,
    COALESCE(SUM(t.amount), 0)::float8 AS actual
FROM
    budgets b
JOIN
    months m ON b.month_id = m.id
LEFT JOIN
    transactions t
    ON b.category = t.category
    AND b.transaction_type = t.transaction_type
    AND t.date BETWEEN
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')))
        AND
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')) + INTERVAL '1 month' - INTERVAL '1 day')
WHERE
    b.month_id = $1 AND b.transaction_type = $2
GROUP BY
    b.id, b.category, b.transaction_type, b.amount
ORDER BY
    b.category
`

type GetBudgetItemsForMonthIDByTransactionTypeParams struct {
	MonthID         int
	TransactionType types.TransactionType
}

type GetBudgetItemsForMonthIDByTransactionTypeRow struct {
	BudgetID int
	Category types.Category
	Type     types.TransactionType
	Planned  float64
	Actual   float64
}

func (q *Queries) GetBudgetItemsForMonthIDByTransactionType(ctx context.Context, arg GetBudgetItemsForMonthIDByTransactionTypeParams) ([]GetBudgetItemsForMonthIDByTransactionTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getBudgetItemsForMonthIDByTransactionType, arg.MonthID, arg.TransactionType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBudgetItemsForMonthIDByTransactionTypeRow
	for rows.Next() {
		var i GetBudgetItemsForMonthIDByTransactionTypeRow
		if err := rows.Scan(
			&i.BudgetID,
			&i.Category,
			&i.Type,
			&i.Planned,
			&i.Actual,
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

const getBudgetsByMonthIDAndType = `-- name: GetBudgetsByMonthIDAndType :many
SELECT id, month_id, category, amount, transaction_type
FROM budgets
WHERE month_id = $1 AND transaction_type = $2
`

type GetBudgetsByMonthIDAndTypeParams struct {
	MonthID         int
	TransactionType types.TransactionType
}

type GetBudgetsByMonthIDAndTypeRow struct {
	ID              int
	MonthID         int
	Category        types.Category
	Amount          float64
	TransactionType types.TransactionType
}

func (q *Queries) GetBudgetsByMonthIDAndType(ctx context.Context, arg GetBudgetsByMonthIDAndTypeParams) ([]GetBudgetsByMonthIDAndTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getBudgetsByMonthIDAndType, arg.MonthID, arg.TransactionType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBudgetsByMonthIDAndTypeRow
	for rows.Next() {
		var i GetBudgetsByMonthIDAndTypeRow
		if err := rows.Scan(
			&i.ID,
			&i.MonthID,
			&i.Category,
			&i.Amount,
			&i.TransactionType,
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

const patchBudget = `-- name: PatchBudget :one
UPDATE budgets SET amount = $1 WHERE id = $2
RETURNING id, month_id, category, amount, transaction_type, created_at, updated_at
`

type PatchBudgetParams struct {
	Amount float64
	ID     int
}

func (q *Queries) PatchBudget(ctx context.Context, arg PatchBudgetParams) (Budget, error) {
	row := q.db.QueryRowContext(ctx, patchBudget, arg.Amount, arg.ID)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.MonthID,
		&i.Category,
		&i.Amount,
		&i.TransactionType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
