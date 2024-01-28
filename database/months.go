package database

import (
	"strings"

	"github.com/ethansaxenian/budgeting/types"
)

func (db *DB) GetMonths() ([]types.Month, error) {
	rows, err := db.db.Query("SELECT id, month, year FROM months")
	if err != nil {
		return nil, err
	}

	var months []types.Month

	for rows.Next() {
		m := types.Month{}
		if err = rows.Scan(
			&m.ID,
			&m.Month,
			&m.Year,
		); err != nil {
			return nil, err
		}
		months = append(months, m)
	}

	return months, nil
}

func (db *DB) GetMonthByID(id int) (types.Month, error) {
	row := db.db.QueryRow("SELECT id, month, year  FROM months WHERE id=$1", id)

	m := types.Month{}
	if err := row.Scan(
		&m.ID,
		&m.Month,
		&m.Year,
	); err != nil {
		return types.Month{}, err
	}

	return m, nil
}

func (db *DB) GetMonthByMonthAndYear(monthStr string) (types.Month, error) {
	parts := strings.Split(monthStr, "-")
	row := db.db.QueryRow("SELECT id, month, year FROM months WHERE month=$1 AND year=$2", parts[1], parts[0])

	m := types.Month{}
	if err := row.Scan(
		&m.ID,
		&m.Month,
		&m.Year,
	); err != nil {
		return types.Month{}, err
	}

	return m, nil
}

func (db *DB) CreateMonth(newMonth types.MonthCreate) (int, error) {
	res, err := db.db.Exec("INSERT INTO months (month, year) VALUES ($1, $2)", newMonth.Month, newMonth.Year)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
