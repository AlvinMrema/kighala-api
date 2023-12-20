-- name: CreateWord :one
INSERT INTO words (id, created_at, updated_at, word)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetWords :many
SELECT * FROM words;

-- name: GetWordById :one
SELECT * FROM words
WHERE id = $1;

-- name: UpdateWord :one
UPDATE words
SET updated_at = $2, word = $3
WHERE id = $1
RETURNING *;

-- name: DeleteWord :exec
DELETE FROM words
WHERE id = $1;