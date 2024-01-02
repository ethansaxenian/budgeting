package database

import "github.com/ethansaxenian/budgeting/types"

func (db *DB) GetMonths() ([]types.Month, error) {
	rows, err := db.db.Query("SELECT * FROM months")
	if err != nil {
		return nil, err
	}

	var months []types.Month

	for rows.Next() {
		m := types.Month{}
		if err = rows.Scan(
			&m.ID,
			&m.MonthID,
			&m.StartingBalance,
		); err != nil {
			return nil, err
		}
		months = append(months, m)
	}

	return months, nil
}

func (db *DB) GetMonthByID(id int) (types.Month, error) {
	row := db.db.QueryRow("SELECT * FROM months WHERE id=$1", id)

	m := types.Month{}
	if err := row.Scan(
		&m.ID,
		&m.MonthID,
		&m.StartingBalance,
	); err != nil {
		return types.Month{}, err
	}

	return m, nil
}

func (db *DB) CreateMonth(month types.MonthCreate) (int, error) {
	res, err := db.db.Exec("INSERT INTO months (month_id, starting_balance) VALUES ($1, $2)", month.MonthID, month.StartingBalance)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (db *DB) UpdateMonth(id int, month types.MonthUpdate) (int, error) {
	res, err := db.db.Exec("UPDATE months SET month_id=$1, starting_balance=$2 WHERE id=$3", month.MonthID, month.StartingBalance, id)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows), nil
}
