-- +goose Up
DROP TYPE IF EXISTS BusStatus;
CREATE TYPE BusStatus AS ENUM ('new', 'in_progress', 'success', 'failed');

CREATE TABLE IF NOT EXISTS bus (
    uniqid uuid NOT NULL,
    queue VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    command_name VARCHAR(255) NOT NULL,
    payload jsonb NOT NULL,
    status BusStatus NOT NULL,
    run_after TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    timeout_at TIMESTAMP WITHOUT TIME ZONE,
    attempt_left BIGINT,
    UNIQUE (uniqid, queue)
);
-- +goose Down
DROP TABLE IF EXISTS bus;
DROP TYPE IF EXISTS BusStatus;