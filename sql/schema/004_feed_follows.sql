-- +goose Up

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    userId UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feedId UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(userid, feedid)
);

-- +goose Down

DROP TABLE feed_follows;