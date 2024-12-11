-- +goose Up
CREATE TABLE messages
(
    id        SERIAL PRIMARY KEY,
    chat_id   BIGINT       NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    "from"    VARCHAR(255) NOT NULL,
    text      TEXT         NOT NULL,
    timestamp TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS messages;
