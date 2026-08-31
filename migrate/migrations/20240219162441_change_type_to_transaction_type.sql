-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions RENAME COLUMN type TO transaction_type;
ALTER TABLE budgets RENAME COLUMN type TO transaction_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions RENAME COLUMN transaction_type TO type;
ALTER TABLE budgets RENAME COLUMN transaction_type TO type;
-- +goose StatementEnd
