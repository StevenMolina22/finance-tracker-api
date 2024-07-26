-- name: GetTransaction :one
SELECT
    *
FROM
    transactions
WHERE
    id = ?
LIMIT
    1;

-- name: ListTransactions :many
SELECT
    *
FROM
    transactions
ORDER BY
    date;

-- name: CreateTransaction :one
INSERT INTO
    transactions (user_id, amount, date, description, type)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions
SET
    user_id = ?,
    amount = ?,
    date = ?,
    description = ?,
    type = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteTransaction :one
DELETE FROM transactions
WHERE
    id = ? RETURNING *;