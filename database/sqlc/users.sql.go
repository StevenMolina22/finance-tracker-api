// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (clerk_id)
VALUES
    (?) RETURNING id, clerk_id, created_at, updated_at
`

func (q *Queries) CreateUser(ctx context.Context, clerkID string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, clerkID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE
    id = ? RETURNING id, clerk_id, created_at, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT
    id, clerk_id, created_at, updated_at
FROM
    users
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT
    id, clerk_id, created_at, updated_at
FROM
    users
ORDER BY
    clerk_id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    clerk_id = ?
WHERE
    id = ? RETURNING id, clerk_id, created_at, updated_at
`

type UpdateUserParams struct {
	ClerkID string `json:"clerk_id"`
	ID      int64  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ClerkID, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
