-- name: CreateDefinition :one
INSERT INTO definitions (id, created_at, updated_at, word_id, definition, part_of_speech)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetDefinitionsByWordID :many
SELECT * FROM definitions
WHERE word_id = $1;

-- name: GetDefinitions :many
SELECT * FROM definitions;

-- name: GetDefinitionById :one
SELECT * FROM definitions
WHERE id = $1;

-- name: UpdateDefinition :one
UPDATE definitions
SET updated_at = $2, definition = $3, part_of_speech = $4, word_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteDefinition :exec
DELETE FROM definitions
WHERE id = $1;