-- name: CreateUser :one
INSERT INTO users (
    username, password, nickname, email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    nickname = COALESCE(sqlc.narg(nickname), nickname),
    gender = COALESCE(sqlc.narg(gender), gender),
    status = COALESCE(sqlc.narg(status), status),
    last_login_at = COALESCE(sqlc.narg(last_login_at), last_login_at),
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

-- name: ExistEmail :one
SELECT COUNT(*) FROM users
WHERE email = $1;

-- name: ExistNickname :one
SELECT COUNT(*) FROM users
WHERE nickname = $1;