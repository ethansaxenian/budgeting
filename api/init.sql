CREATE TABLE IF NOT EXISTS months (
    id varchar(250) NOT NULL,
    starting_balance FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT NOT NULL,
    date DATE NOT NULL,
    amount FLOAT NOT NULL,
    description varchar(250) NOT NULL,
    category varchar(250) NOT NULL,
    type varchar(250) NOT NULL,
    month_id varchar(250) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS plans (
    id INT NOT NULL,
    month INT NOT NULL,
    year INT NOT NULL,
    category varchar(250) NOT NULL,
    amount FLOAT NOT NULL,
    type varchar(250) NOT NULL,
    month_id varchar(250) NOT NULL,
    PRIMARY KEY (id)
);
