CREATE TABLE IF NOT EXISTS months (
    id SERIAL PRIMARY KEY,
    month_id varchar(250) NOT NULL,
    starting_balance FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    amount FLOAT NOT NULL,
    description varchar(250),
    category varchar(250) NOT NULL,
    type varchar(250) NOT NULL,
    month_id INT NOT NULL REFERENCES months (id)
);

CREATE TABLE IF NOT EXISTS plans (
    id SERIAL PRIMARY KEY,
    category varchar(250) NOT NULL,
    amount FLOAT NOT NULL,
    type varchar(250) NOT NULL,
    month_id INT NOT NULL REFERENCES months (id)
);
