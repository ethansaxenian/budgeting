package database

import "github.com/ethansaxenian/budgeting/types"

func (db *DB) GetBudgetsByMonthIDAndType(monthID int, transactionType types.TransactionType) ([]types.Budget, error) {
	rows, err := db.DB.Query("SELECT id, month_id, category, amount, transaction_type FROM budgets WHERE month_id = $1 AND transaction_type = $2", monthID, transactionType)
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
	row := db.DB.QueryRow("SELECT id, month_id, category, amount, transaction_type FROM budgets WHERE id = $1", id)

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
	_, err := db.DB.Exec("UPDATE budgets SET amount = $1 WHERE id = $2", amount, id)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateNewBudgetsForMonth(monthID int) error {
	_, err := db.DB.Exec(`
		INSERT INTO
			budgets (month_id, category, amount, transaction_type)
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
		`, monthID)

	if err != nil {
		return err
	}

	return nil
}
