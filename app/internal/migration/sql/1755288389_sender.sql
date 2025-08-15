-- +goose Up
CREATE TABLE IF NOT EXISTS sender (
phone_number text not null,
messenger text not null,
link jsonb not null,
    PRIMARY KEY (phone_number, messenger)
);
-- +goose Down
DROP TABLE IF EXISTS sender;
