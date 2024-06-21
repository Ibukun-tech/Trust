-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- To implement transaction database and trying to see the relationship between users and how they can communicate with each other

CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(255), last_name VARCHAR(255), email VARCHAR(255) UNIQUE, password VARCHAR(255), hashed_password VARCHAR(255), age INT,  active BOOLEAN,  created_at TIMESTAMP, updated_at TIMESTAMP)
CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY,account_number VARCHAR(50) , user_id INTEGER FOREIGN KEY UNIQUE REFERENCES users(id) ON DELETE CASCADE,active BOOLEAN, created_at TIMESTAMP, updated_at TIMESTAMP  )
-- CREATE TABLE IF NOT EXISTS transactions (id SERIAL PRIMARY KEY, describe VARCHAR(255),  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users
DROP TABLE IF EXISTS accounts
-- DROP TABLE IF EXISTS transactions
-- +goose StatementEnd
