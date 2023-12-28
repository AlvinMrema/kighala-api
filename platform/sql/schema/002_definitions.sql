-- +goose Up

CREATE TABLE definitions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    definition TEXT UNIQUE NOT NULL,
    part_of_speech TEXT NOT NULL,
    word_id UUID NOT NULL REFERENCES words(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE definitions;