package db

import "github.com/ethansaxenian/budgeting/types"

func GetMonths() ([]types.Month, error) {
	rows, err := DB.Query("SELECT * FROM months")
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

func GetMonthByID(id int) (types.Month, error) {
	row := DB.QueryRow("SELECT * FROM months WHERE id=$1", id)

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
