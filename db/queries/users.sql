-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ($1, $2) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET email = $1,
    password = $2
WHERE id = $3;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;