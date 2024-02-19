-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions
  ALTER COLUMN description TYPE TEXT,
  ALTER COLUMN category TYPE TEXT,
  ALTER COLUMN type TYPE TEXT;

ALTER TABLE budgets
  ALTER COLUMN category TYPE TEXT,
  ALTER COLUMN type TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
  ALTER COLUMN description TYPE character varying(250),
  ALTER COLUMN category TYPE character varying(250),
  ALTER COLUMN type TYPE character varying(250);

ALTER TABLE budgets
  ALTER COLUMN category TYPE character varying(250),
  ALTER COLUMN type TYPE character varying(250);
-- +goose StatementEnd
