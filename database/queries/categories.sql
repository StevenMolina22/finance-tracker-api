-- name: GetCategory :one
SELECT
    *
FROM
    categories
WHERE
    id = ?
LIMIT
    1;

-- name: ListCategories :many
SELECT
    *
FROM
    categories
ORDER BY
    name;

-- name: CreateCategory :one
INSERT INTO
    categories (name, description)
VALUES
    (?, ?) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = ?,
    description = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteCategory :one
DELETE FROM categories
WHERE
    id = ? RETURNING *;