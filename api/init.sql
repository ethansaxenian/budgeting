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
    type varchar(250) NOT NULL,
    month_id INT NOT NULL REFERENCES months(id),
    food FLOAT NOT NULL,
    gifts FLOAT NOT NULL,
    medical FLOAT NOT NULL,
    home FLOAT NOT NULL,
    transportation FLOAT NOT NULL,
    personal FLOAT NOT NULL,
    savings FLOAT NOT NULL,
    utilities FLOAT NOT NULL,
    travel FLOAT NOT NULL,
    other FLOAT NOT NULL,
    paycheck FLOAT NOT NULL,
    bonus FLOAT NOT NULL,
    interest FLOAT NOT NULL
);
