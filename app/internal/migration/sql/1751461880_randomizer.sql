-- +goose Up
CREATE TABLE randomizer (
    key VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX randomizer_expired_at_key ON randomizer (expired_at, key);
CREATE INDEX randomizer_key ON randomizer (key);

ALTER TABLE users ADD COLUMN confirmed_email BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE users ADD COLUMN confirmed_phone BOOLEAN NOT NULL DEFAULT false;
-- +goose Down
DROP TABLE randomizer;

ALTER TABLE users DROP COLUMN confirmed_email;
ALTER TABLE users DROP COLUMN confirmed_phone;
