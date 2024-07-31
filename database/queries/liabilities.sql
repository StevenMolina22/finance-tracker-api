-- name: GetLiability :one
SELECT
    *
FROM
    liabilities
WHERE
    id = ?
LIMIT
    1;

-- name: ListLiabilities :many
SELECT
    *
FROM
    liabilities
ORDER BY
    name;

-- name: CreateLiability :one
INSERT INTO
    liabilities (user_id, name, amount)
VALUES
    (?, ?, ?) RETURNING *;

-- name: UpdateLiability :one
UPDATE liabilities
SET
    user_id = ?,
    name = ?,
    amount = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteLiability :one
DELETE FROM liabilities
WHERE
    id = ? RETURNING *;