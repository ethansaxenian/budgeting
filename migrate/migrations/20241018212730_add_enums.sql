-- +goose Up
-- +goose StatementBegin
CREATE TYPE transaction_type AS ENUM (
    'expense',
    'income'
);

CREATE TYPE category AS ENUM (
    'food',
    'gifts',
    'medical',
    'home',
    'transportation',
    'personal',
    'savings',
    'utilities',
    'travel',
    'other',
    'paycheck',
    'bonus',
    'interest',
    'cashback'
);

ALTER TABLE transactions
  ALTER COLUMN transaction_type TYPE transaction_type USING transaction_type::transaction_type,
  ALTER COLUMN category TYPE category USING category::category;

ALTER TABLE budgets
  ALTER COLUMN transaction_type TYPE transaction_type USING transaction_type::transaction_type,
  ALTER COLUMN category TYPE category USING category::category;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE budgets
  ALTER COLUMN category TYPE text USING category::text,
  ALTER COLUMN transaction_type TYPE text USING transaction_type::text;

ALTER TABLE transactions
  ALTER COLUMN category TYPE text USING category::text,
  ALTER COLUMN transaction_type TYPE text USING transaction_type::text;

DROP TYPE IF EXISTS category;
DROP TYPE IF EXISTS transaction_type;
-- +goose StatementEnd
