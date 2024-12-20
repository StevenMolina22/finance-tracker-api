// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: assets.sql

package sqlc

import (
	"context"
)

const createAsset = `-- name: CreateAsset :one
INSERT INTO
    assets (user_id, name, amount)
VALUES
    (?, ?, ?) RETURNING id, user_id, name, amount
`

type CreateAssetParams struct {
	UserID int64   `json:"user_id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

func (q *Queries) CreateAsset(ctx context.Context, arg CreateAssetParams) (Asset, error) {
	row := q.db.QueryRowContext(ctx, createAsset, arg.UserID, arg.Name, arg.Amount)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Amount,
	)
	return i, err
}

const deleteAsset = `-- name: DeleteAsset :one
DELETE FROM assets
WHERE
    id = ? RETURNING id, user_id, name, amount
`

func (q *Queries) DeleteAsset(ctx context.Context, id int64) (Asset, error) {
	row := q.db.QueryRowContext(ctx, deleteAsset, id)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Amount,
	)
	return i, err
}

const getAsset = `-- name: GetAsset :one
SELECT
    id, user_id, name, amount
FROM
    assets
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetAsset(ctx context.Context, id int64) (Asset, error) {
	row := q.db.QueryRowContext(ctx, getAsset, id)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Amount,
	)
	return i, err
}

const listAssets = `-- name: ListAssets :many
SELECT
    id, user_id, name, amount
FROM
    assets
ORDER BY
    name
`

func (q *Queries) ListAssets(ctx context.Context) ([]Asset, error) {
	rows, err := q.db.QueryContext(ctx, listAssets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Asset
	for rows.Next() {
		var i Asset
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Amount,
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

const updateAsset = `-- name: UpdateAsset :one
UPDATE assets
SET
    user_id = ?,
    name = ?,
    amount = ?
WHERE
    id = ? RETURNING id, user_id, name, amount
`

type UpdateAssetParams struct {
	UserID int64   `json:"user_id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	ID     int64   `json:"id"`
}

func (q *Queries) UpdateAsset(ctx context.Context, arg UpdateAssetParams) (Asset, error) {
	row := q.db.QueryRowContext(ctx, updateAsset,
		arg.UserID,
		arg.Name,
		arg.Amount,
		arg.ID,
	)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Amount,
	)
	return i, err
}
