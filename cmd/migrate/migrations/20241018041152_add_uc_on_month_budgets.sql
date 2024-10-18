-- +goose Up
-- +goose StatementBegin
ALTER TABLE budgets
ADD CONSTRAINT budgets_month_id_category_transaction_type_uc
UNIQUE (month_id, category, transaction_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE budgets
DROP CONSTRAINT budgets_month_id_category_transaction_type_uc;
-- +goose StatementEnd
