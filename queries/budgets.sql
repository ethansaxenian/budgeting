-- name: GetBudgetByID :one
SELECT * FROM budgets WHERE id = $1;

-- name: GetBudgetsByMonthIDAndType :many
SELECT id, month_id, category, amount, transaction_type
FROM budgets
WHERE month_id = $1 AND transaction_type = $2;

-- name: PatchBudget :one
UPDATE budgets SET amount = $1 WHERE id = $2
RETURNING *;

-- name: CreateNewBudgetsForMonth :many
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
RETURNING *;

-- name: GetBudgetItemsForMonthIDByTransactionType :many
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
    b.category;
