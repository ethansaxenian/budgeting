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
