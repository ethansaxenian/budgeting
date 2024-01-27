-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS months (
    id SERIAL PRIMARY KEY,
    month_id character varying(250) NOT NULL,
    starting_balance double precision NOT NULL,
    year integer,
    month integer
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date date NOT NULL,
    amount double precision NOT NULL,
    description character varying(250),
    category character varying(250) NOT NULL,
    type character varying(250) NOT NULL
);

CREATE TABLE IF NOT EXISTS budgets (
    id SERIAL PRIMARY KEY,
    month_id integer REFERENCES months(id),
    category character varying(250),
    amount double precision NOT NULL DEFAULT '0'::double precision,
    type character varying(250)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE months;
DROP TABLE transactions;
DROP TABLE budgets;
-- +goose StatementEnd
