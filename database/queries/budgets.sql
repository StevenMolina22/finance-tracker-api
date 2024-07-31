-- name: GetBudget :one
SELECT
    *
FROM
    budgets
WHERE
    id = ?
LIMIT
    1;

-- name: ListBudgets :many
SELECT
    *
FROM
    budgets
ORDER BY
    start_date;

-- name: CreateBudget :one
INSERT INTO
    budgets (category_id, amount, start_date, end_date)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: UpdateBudget :one
UPDATE budgets
SET
    category_id = ?,
    amount = ?,
    start_date = ?,
    end_date = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteBudget :one
DELETE FROM budgets
WHERE
    id = ? RETURNING *;