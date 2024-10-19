-- name: GetAllMonths :many
SELECT * FROM months ORDER BY year DESC, month DESC;

-- name: GetMonthByID :one
SELECT * FROM months WHERE id = $1;

-- name: GetMonthByMonthAndYear :one
SELECT * FROM months WHERE month=$1 AND year=$2;

-- name: CreateMonth :one
INSERT INTO months (month, year)
VALUES ($1, $2)
RETURNING *;

