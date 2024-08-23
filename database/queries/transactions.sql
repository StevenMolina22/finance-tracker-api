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
    date DESC;

-- name: CreateTransaction :one
INSERT INTO
    transactions (
        user_id,
        category_id,
        name,
        amount,
        description,
        type,
        date
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions
SET
    user_id = ?,
    category_id = ?,
    name = ?,
    amount = ?,
    description = ?,
    type = ?,
    date = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteTransaction :one
DELETE FROM transactions
WHERE
    id = ? RETURNING *;