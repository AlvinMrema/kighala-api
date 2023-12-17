-- name: CreateWord :one
INSERT INTO words (id, created_at, updated_at, word)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetWords :many
SELECT * FROM words;