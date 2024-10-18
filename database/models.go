// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/ethansaxenian/budgeting/types"
)

type Budget struct {
	ID              int
	MonthID         int
	Category        types.Category
	Amount          float64
	TransactionType types.TransactionType
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type Month struct {
	ID        int
	Year      int
	Month     time.Month
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Transaction struct {
	ID              int
	Date            time.Time
	Amount          float64
	Description     sql.NullString
	Category        types.Category
	TransactionType types.TransactionType
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}
