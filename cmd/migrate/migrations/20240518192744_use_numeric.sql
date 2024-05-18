-- +goose Up
-- +goose StatementBegin
ALTER TABLE budgets ALTER COLUMN amount TYPE numeric USING amount::numeric;
ALTER TABLE budgets ALTER COLUMN amount SET DEFAULT 0;
ALTER TABLE transactions ALTER COLUMN amount TYPE numeric USING amount::numeric;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE budgets ALTER COLUMN amount TYPE double precision USING amount::double precision;
ALTER TABLE budgets ALTER COLUMN amount SET DEFAULT '0'::double precision;
ALTER TABLE transactions ALTER COLUMN amount TYPE double precision USING amount::double precision;
-- +goose StatementEnd
