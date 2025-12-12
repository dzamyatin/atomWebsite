-- +goose Up
CREATE TABLE chat (
  messenger text not null,
  chat_id text not null,
  state text not null,
  PRIMARY KEY (messenger, chat_id)
);
-- +goose Down
DROP TABLE chat;
