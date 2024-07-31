// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: budgets.sql

package sqlc

import (
	"context"
	"time"
)

const createBudget = `-- name: CreateBudget :one
INSERT INTO
    budgets (category_id, amount, start_date, end_date)
VALUES
    (?, ?, ?, ?) RETURNING id, category_id, amount, start_date, end_date
`

type CreateBudgetParams struct {
	CategoryID int64     `json:"category_id"`
	Amount     float64   `json:"amount"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

func (q *Queries) CreateBudget(ctx context.Context, arg CreateBudgetParams) (Budget, error) {
	row := q.db.QueryRowContext(ctx, createBudget,
		arg.CategoryID,
		arg.Amount,
		arg.StartDate,
		arg.EndDate,
	)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Amount,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const deleteBudget = `-- name: DeleteBudget :one
DELETE FROM budgets
WHERE
    id = ? RETURNING id, category_id, amount, start_date, end_date
`

func (q *Queries) DeleteBudget(ctx context.Context, id int64) (Budget, error) {
	row := q.db.QueryRowContext(ctx, deleteBudget, id)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Amount,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getBudget = `-- name: GetBudget :one
SELECT
    id, category_id, amount, start_date, end_date
FROM
    budgets
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetBudget(ctx context.Context, id int64) (Budget, error) {
	row := q.db.QueryRowContext(ctx, getBudget, id)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Amount,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const listBudgets = `-- name: ListBudgets :many
SELECT
    id, category_id, amount, start_date, end_date
FROM
    budgets
ORDER BY
    start_date
`

func (q *Queries) ListBudgets(ctx context.Context) ([]Budget, error) {
	rows, err := q.db.QueryContext(ctx, listBudgets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Budget
	for rows.Next() {
		var i Budget
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Amount,
			&i.StartDate,
			&i.EndDate,
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

const updateBudget = `-- name: UpdateBudget :one
UPDATE budgets
SET
    category_id = ?,
    amount = ?,
    start_date = ?,
    end_date = ?
WHERE
    id = ? RETURNING id, category_id, amount, start_date, end_date
`

type UpdateBudgetParams struct {
	CategoryID int64     `json:"category_id"`
	Amount     float64   `json:"amount"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	ID         int64     `json:"id"`
}

func (q *Queries) UpdateBudget(ctx context.Context, arg UpdateBudgetParams) (Budget, error) {
	row := q.db.QueryRowContext(ctx, updateBudget,
		arg.CategoryID,
		arg.Amount,
		arg.StartDate,
		arg.EndDate,
		arg.ID,
	)
	var i Budget
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Amount,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}
