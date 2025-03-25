-- +goose Up
DROP TYPE IF EXISTS BusStatus;
CREATE TYPE BusStatus AS ENUM ('new', 'in_progress', 'success', 'failed');

CREATE TABLE IF NOT EXISTS bus (
    uniqid uuid NOT NULL,
    queue VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    payload jsonb NOT NULL,
    status BusStatus NOT NULL,
    handled_at TIMESTAMP WITHOUT TIME ZONE,
    UNIQUE (uniqid, queue)
);
-- +goose Down
DROP TABLE IF EXISTS bus;
DROP TYPE IF EXISTS BusStatus;