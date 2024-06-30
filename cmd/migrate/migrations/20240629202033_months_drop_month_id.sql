-- +goose Up
-- +goose StatementBegin
ALTER TABLE months DROP COLUMN month_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE months ADD COLUMN month_id character varying(250) NOT NULL;
-- +goose StatementEnd
