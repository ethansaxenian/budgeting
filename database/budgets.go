package database

import "github.com/ethansaxenian/budgeting/types"

func (db *DB) GetBudgets(monthID int) ([]types.Budget, error) {
	rows, err := db.db.Query("SELECT id, month_id, category, amount, type FROM budgets WHERE month_id = $1", monthID)
	if err != nil {
		return nil, err
	}

	var budgets []types.Budget

	for rows.Next() {
		b := types.Budget{}
		if err = rows.Scan(
			&b.ID,
			&b.MonthID,
			&b.Category,
			&b.Amount,
			&b.Type,
		); err != nil {
			return nil, err
		}
		budgets = append(budgets, b)
	}

	return budgets, nil
}

func (db *DB) GetBudgetByID(id int) (types.Budget, error) {
	row := db.db.QueryRow("SELECT id, month_id, category, amount, type FROM budgets WHERE id = $1", id)

	b := types.Budget{}
	if err := row.Scan(
		&b.ID,
		&b.MonthID,
		&b.Category,
		&b.Amount,
		&b.Type,
	); err != nil {
		return types.Budget{}, err
	}

	return b, nil
}

func (db *DB) PatchBudget(id int, amount float64) error {
	_, err := db.db.Exec("UPDATE budgets SET amount = $1 WHERE id = $2", amount, id)
	if err != nil {
		return err
	}

	return nil
}
