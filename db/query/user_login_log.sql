-- name: GetLastUserLoginLog :one
SELECT * FROM user_login_logs
WHERE user_id = $1
ORDER BY id DESC
LIMIT 1;

-- name: AddUserLoginLog :one
INSERT INTO user_login_logs (
    user_id, login_ip, login_ip_region, is_login_exceptional, platform, user_agent
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateLastUserLoginLog :exec
UPDATE user_login_logs
SET
    is_login_exceptional = true
WHERE
	id = (
            SELECT id FROM user_login_logs 
            WHERE user_login_logs.user_id = $1 AND NOT is_login_exceptional 
            ORDER BY id DESC 
            LIMIT 1
        );
