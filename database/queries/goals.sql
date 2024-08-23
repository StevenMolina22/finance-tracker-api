-- name: ListGoals :many
SELECT
    *
FROM
    goals
ORDER BY
    start_date;

-- name: GetGoal :one
SELECT
    *
FROM
    goals
WHERE
    id = ?
LIMIT
    1;

-- name: CreateGoal :one
INSERT INTO
    goals (
        user_id,
        type,
        amount,
        start_date,
        end_date,
        status
    )
VALUES
    (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateGoal :one
UPDATE goals
SET
    user_id = ?,
    type = ?,
    amount = ?,
    start_date = ?,
    end_date = ?,
    status = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteGoal :one
DELETE FROM goals
WHERE
    id = ? RETURNING *;

-- name: GetGoalsByUserId :many
SELECT
    *
FROM
    goals
WHERE
    user_id = ?
ORDER BY
    start_date;
