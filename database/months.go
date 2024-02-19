package database

import (
	"database/sql"
	"time"

	"github.com/ethansaxenian/budgeting/types"
	"github.com/ethansaxenian/budgeting/util"
)

func (db *DB) GetMonths() ([]types.Month, error) {
	rows, err := db.DB.Query("SELECT id, month, year FROM months")
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
	row := db.DB.QueryRow("SELECT id, month, year  FROM months WHERE id=$1", id)

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

func (db *DB) GetMonthByMonthAndYear(m time.Month, y int) (types.Month, error) {
	row := db.DB.QueryRow("SELECT id, month, year FROM months WHERE month=$1 AND year=$2", m, y)

	newMonth := types.Month{}
	if err := row.Scan(
		&newMonth.ID,
		&newMonth.Month,
		&newMonth.Year,
	); err != nil {
		return types.Month{}, err
	}

	return newMonth, nil
}

func (db *DB) CreateMonth(newMonth types.MonthCreate) error {
	_, err := db.DB.Exec("INSERT INTO months (month, year) VALUES ($1, $2)", newMonth.Month, newMonth.Year)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) createCurrMonth(m time.Month, y int) (types.Month, error) {
	if err := db.CreateMonth(types.MonthCreate{Month: m, Year: y}); err != nil {
		return types.Month{}, err
	}

	currMonth, err := db.GetMonthByMonthAndYear(m, y)
	if err != nil {
		return types.Month{}, err
	}

	if err := db.CreateNewBudgetsForMonth(currMonth.ID); err != nil {
		return types.Month{}, err
	}

	return currMonth, nil
}

func (db *DB) GetOrCreateCurrentMonth() (types.Month, error) {
	currYear, currMonth, _ := util.CurrentDate()

	row := db.DB.QueryRow("SELECT id, month, year FROM months WHERE month=$1 AND year=$2", currMonth, currYear)

	m := types.Month{}
	err := row.Scan(
		&m.ID,
		&m.Month,
		&m.Year,
	)

	switch err {
	case nil:
		return m, nil
	case sql.ErrNoRows:
		return db.createCurrMonth(currMonth, currYear)
	default:
		return types.Month{}, err
	}

}
