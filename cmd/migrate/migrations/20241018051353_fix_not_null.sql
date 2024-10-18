-- +goose Up
-- +goose StatementBegin
ALTER TABLE months
  ALTER COLUMN year SET NOT NULL,
  ALTER COLUMN month SET NOT NULL;

ALTER TABLE budgets
  ALTER COLUMN month_id SET NOT NULL,
  ALTER COLUMN category SET NOT NULL,
  ALTER COLUMN transaction_type SET NOT NULL;

ALTER TABLE transactions
  ALTER COLUMN description SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
  ALTER COLUMN description DROP NOT NULL;

ALTER TABLE budgets
  ALTER COLUMN month_id DROP NOT NULL,
  ALTER COLUMN category DROP NOT NULL,
  ALTER COLUMN transaction_type DROP NOT NULL;

ALTER TABLE months
  ALTER COLUMN year DROP NOT NULL,
  ALTER COLUMN month DROP NOT NULL;
-- +goose StatementEnd
