-- +goose Up
CREATE TABLE chat_users
(
    chat_id  BIGINT       NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL,
    PRIMARY KEY (chat_id, username)
);

-- +goose Down
DROP TABLE IF EXISTS chat_users;
