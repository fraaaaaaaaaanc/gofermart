-- Thank you for giving goose a try!
-- 
-- This file was automatically created running goose init. If you're familiar with goose
-- feel free to remove/rename this file, write some SQL and goose up. Briefly,
-- 
-- Documentation can be found here: https://pressly.github.io/goose
--
-- A single goose .sql file holds both Up and Down migrations.
-- 
-- All goose .sql files are expected to have a -- +goose Up annotation.
-- The -- +goose Down annotation is optional, but recommended, and must come after the Up annotation.
-- 
-- The -- +goose NO TRANSACTION annotation may be added to the top of the file to run statements 
-- outside a transaction. Both Up and Down migrations within this file will be run without a transaction.
-- 
-- More complex statements that have semicolons within them must be annotated with 
-- the -- +goose StatementBegin and -- +goose StatementEnd annotations to be properly recognized.
-- 
-- Use GitHub issues for reporting bugs and requesting features, enjoy!

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
