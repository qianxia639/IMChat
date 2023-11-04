-- name: CreateFriendGroupApply :one
INSERT INTO friend_group_applys(
    user_id, target_id, description, apply_type
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ExistsFriendGroupApply :one
SELECT COUNT(*) FROM friend_group_applys
WHERE
    user_id = $1 AND target_id = $2 AND status = 1 AND apply_type = $3;

-- name: ListFriendGroupApply :many
SELECT fca.user_id, u.nickname, fca.description,  fca.apply_type, fca.created_at
FROM friend_group_applys fca
JOIN users u ON fca.user_id = u.id
WHERE fca.target_id = $1 AND fca.status = 1;

-- name: UpdateFriendGroupApply :exec
UPDATE friend_group_applys
SET
    status = @status,
    updated_at = now()
WHERE 
    (user_id = @user_id AND target_id = @target_id AND status = 1 AND apply_type = @apply_type)
    OR
    (user_id = @target_id AND target_id = @user_id AND status = 1 AND apply_type = @apply_type);

-- name: DeleteFriendGroupApply :exec
DELETE FROM friend_group_applys
WHERE user_id = $1 OR target_id = $1;