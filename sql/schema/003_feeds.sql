-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;