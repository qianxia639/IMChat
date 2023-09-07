-- name: CreateUser :one
INSERT INTO users (
    username, password, nickname, email, gender
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    nickname = COALESCE(sqlc.narg(nickname), nickname),
    avatar = COALESCE(sqlc.narg(avatar), avatar),
    gender = COALESCE(sqlc.narg(gender), gender),
    updated_at = sqlc.arg(updated_at)
WHERE
    username = sqlc.arg(username)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET
    password = sqlc.arg(password),
    password_changed_at = sqlc.arg(password_changed_at),
    updated_at = sqlc.arg(updated_at)
WHERE
    username = sqlc.arg(username);