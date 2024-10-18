-- +goose Up
-- +goose StatementBegin
CREATE VIEW budget_items AS
SELECT
    b.id AS budget_id,
    b.month_id AS month_id,
    b.category,
    b.transaction_type,
    b.amount::numeric AS planned,
    COALESCE(SUM(t.amount), 0)::numeric AS actual
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
GROUP BY
    b.id, b.category, b.transaction_type, b.amount;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS budget_item
-- +goose StatementEnd
