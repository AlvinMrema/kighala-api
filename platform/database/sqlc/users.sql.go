// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, password, role)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, username, password, role
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
	Role      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Username,
		arg.Password,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, username, password, role FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, created_at, updated_at, username, password, role FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Password,
		&i.Role,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, created_at, updated_at, username, password, role FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
			&i.Password,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}