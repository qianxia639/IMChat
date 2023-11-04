-- name: AddFriend :one
INSERT INTO friendships (user_id, friend_id, comment)
VALUEs ($1, $2, $3) 
RETURNING *;

-- name: GetFriend :one
SELECT * FROM friendships
WHERE 
    user_id = $1 AND friend_id = $2;

-- name: UpdateFriendComment :one
UPDATE friendships
SET
    comment = @comment
WHERE
    user_id = @user_id AND friend_id = @friend_id
RETURNING *;

-- name: DeleteFriend :exec
DELETE FROM friendships
WHERE
    user_id = $1 AND friend_id = $2;

-- name: ListFriends :many
SELECT f.friend_id, f.comment, u.profile_picture_url FROM friendships AS f
JOIN users AS u ON f.user_id = u.id
WHERE u.id = $1;

-- name: ListFriendshipPending :many
SELECT f.status, f.user_id, u.nickname
FROM friendships f
JOIN users u ON f.user_id = u.id
WHERE
    friend_id = $1 AND status = 1;