-- +goose Up
-- +goose StatementBegin
ALTER TABLE months
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN updated_at TIMESTAMP NULL;

ALTER TABLE budgets
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN updated_at TIMESTAMP NULL;

ALTER TABLE transactions
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN updated_at TIMESTAMP NULL;

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_months_updated_at BEFORE UPDATE ON months FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_budgets_updated_at BEFORE UPDATE ON budgets FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_transactions_updated_at BEFORE UPDATE ON transactions FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE months DROP COLUMN created_at, DROP COLUMN updated_at;
ALTER TABLE budgets DROP COLUMN created_at, DROP COLUMN updated_at;
ALTER TABLE transactions DROP COLUMN created_at, DROP COLUMN updated_at;

DROP TRIGGER IF EXISTS update_months_updated_at ON months;
DROP TRIGGER IF EXISTS update_budgets_updated_at ON budgets;
DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;

DROP FUNCTION IF EXISTS update_updated_at_column;
-- +goose StatementEnd
