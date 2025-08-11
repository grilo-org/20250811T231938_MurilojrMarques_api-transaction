-- +goose Up
CREATE TABLE IF NOT EXISTS transaction(
    id SERIAL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    date DATE NOT NULL,
    value NUMERIC(10, 2) CHECK (VALUE > 0)
);


-- +goose Down
DROP TABLE IF EXISTS transaction;

