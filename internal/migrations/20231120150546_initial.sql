-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL DEFAULT '' UNIQUE,
    password VARCHAR(60) NOT NULL DEFAULT ''
    );

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    order_number VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
    order_status VARCHAR(10) NOT NULL DEFAULT '',
    accrual DECIMAL NOT NULL DEFAULT 0,
    order_datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS orders_info (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    description VARCHAR(128) NOT NULL DEFAULT '',
    price DECIMAL NOT NULL DEFAULT 0
    );

CREATE TABLE IF NOT EXISTS balance (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    user_balance DECIMAL NOT NULL DEFAULT 0 CHECK (user_balance >= 0),
    withdrawn_balance DECIMAL NOT NULL DEFAULT 0
    );

CREATE TABLE IF NOT EXISTS history_balance (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    order_number_unregister VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
    withdrawn_sum DECIMAL NOT NULL DEFAULT 0 CHECK (withdrawn_sum >= 0),
    withdrawn_datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS order_accrual (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    order_id INT REFERENCES orders(id),
    order_number VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
    order_status_accrual VARCHAR(10) NOT NULL DEFAULT '',
    accrual DECIMAL NOT NULL DEFAULT 0
    );

-- +goose Down
SELECT 'down SQL query';
