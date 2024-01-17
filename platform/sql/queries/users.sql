-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, password, role)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;