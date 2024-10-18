-- name: GetTransactionByRow :one
SELECT * FROM transactions WHERE id = $1;

-- name: GetTransactionsByTypeInDateRange :many
SELECT *
FROM transactions
WHERE transaction_type = $1 AND date BETWEEN sqlc.arg(start_date)::date AND sqlc.arg(end_date)::date;


-- name: GetTransactionsByTypeAndCategoryInDateRange :many
SELECT *
FROM transactions
WHERE transaction_type = $1 AND category = $2 AND date BETWEEN sqlc.arg(start_date)::date AND sqlc.arg(end_date)::date;

-- name: CreateTransaction :one
INSERT INTO transactions (description, amount, date, category, transaction_type)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateTransaction :exec
UPDATE transactions
SET description = $1, amount = $2, date = $3, category = $4
WHERE id = $5;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = $1;
