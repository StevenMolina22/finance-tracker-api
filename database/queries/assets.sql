-- name: GetAsset :one
SELECT
    *
FROM
    assets
WHERE
    id = ?
LIMIT
    1;

-- name: ListAssets :many
SELECT
    *
FROM
    assets
ORDER BY
    name;

-- name: CreateAsset :one
INSERT INTO
    assets (user_id, name, amount)
VALUES
    (?, ?, ?) RETURNING *;

-- name: UpdateAsset :one
UPDATE assets
SET
    user_id = ?,
    name = ?,
    amount = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteAsset :one
DELETE FROM assets
WHERE
    id = ? RETURNING *;