-- +goose Up
CREATE TABLE IF NOT EXISTS counter (
  key text not null PRIMARY KEY,
  value bigint,
  created_at timestamp without time zone
);
-- +goose Down
DROP TABLE IF EXISTS counter;
DROP INDEX IF EXISTS idx_counter_key;

