-- name: GetTransactionByRow :one
SELECT * FROM transactions WHERE id = $1;

-- name: GetTransactionsByMonthIDAndType :many
SELECT
    t.id,
    t.date,
    t.amount,
    t.description,
    t.category,
    t.transaction_type,
    t.created_at,
    t.updated_at
FROM
    transactions t
JOIN
    months m ON t.date BETWEEN
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')))
        AND
        (DATE_TRUNC('month', TO_DATE(m.year || '-' || m.month || '-01', 'YYYY-MM-DD')) + INTERVAL '1 month' - INTERVAL '1 day')
WHERE
    m.id = $1 AND t.transaction_type = $2;

-- name: CreateTransaction :one
INSERT INTO transactions (description, amount, date, category, transaction_type)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions
SET description = $1, amount = $2, date = $3, category = $4
WHERE id = $5
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1;

