-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT
    *
FROM
    users
ORDER BY
    clerk_id;

-- name: CreateUser :one
INSERT INTO
    users (clerk_id)
VALUES
    (?) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    clerk_id = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE
    id = ? RETURNING *;