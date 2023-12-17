-- name: CreateDefinition :one
INSERT INTO definitions (id, created_at, updated_at, word_id, definition, part_of_speech)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetDefinitionsByWordID :many
SELECT * FROM definitions
WHERE word_id = $1;