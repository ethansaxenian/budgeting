-- +goose Up
-- +goose StatementBegin
ALTER TABLE months DROP COLUMN starting_balance;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE months ADD COLUMN starting_balance double precision NOT NULL;
-- +goose StatementEnd
