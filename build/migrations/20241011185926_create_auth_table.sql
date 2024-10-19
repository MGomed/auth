-- +goose Up
CREATE TYPE role AS ENUM ('UNKNOWN', 'USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS auth (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    role role NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS auth;
DROP TYPE role;
