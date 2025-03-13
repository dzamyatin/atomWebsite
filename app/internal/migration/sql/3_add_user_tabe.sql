-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    uuid VARCHAR(255) NOT NULL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    phone VARCHAR(255)
);
CREATE UNIQUE INDEX uidx_email ON users (email);
CREATE UNIQUE INDEX uidx_phone ON users (phone);
-- +goose Down
DROP TABLE IF EXISTS users;