-- +goose Up

CREATE TABLE words (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    word TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE words;