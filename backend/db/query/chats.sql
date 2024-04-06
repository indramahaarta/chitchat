-- name: GetChatsHistories :many
SELECT c.sender_uid,
    sender.name AS sender_name,
    c.receiver_uid,
    receiver.name AS receiver_name,
    c.created_at AS latest_created_at,
    c.content AS latest_content
FROM (
        SELECT *,
            ROW_NUMBER() OVER (
                PARTITION BY LEAST(sender_uid, receiver_uid),
                GREATEST(sender_uid, receiver_uid)
                ORDER BY created_at DESC
            ) AS row_num
        FROM chats
        WHERE sender_uid = $1
            OR receiver_uid = $1
    ) c
    LEFT JOIN users sender ON c.sender_uid = sender.uid
    LEFT JOIN users receiver ON c.receiver_uid = receiver.uid
WHERE c.row_num = 1
ORDER BY latest_created_at DESC;
-- name: GetChatsDetails :many
SELECT chats.*,
    (
        SELECT json_build_object(
                'id',
                sender.uid,
                'name',
                sender.name,
                'avatar',
                sender.avatar,
                'email',
                sender.email
            )
        FROM users AS sender
        WHERE sender.uid = chats.sender_uid
    ) AS sender,
    (
        SELECT json_build_object(
                'id',
                receiver.uid,
                'name',
                receiver.name,
                'avatar',
                receiver.avatar,
                'email',
                receiver.email
            )
        FROM users AS receiver
        WHERE receiver.uid = chats.receiver_uid
    ) AS receiver
FROM chats
WHERE (
        sender_uid = $1
        AND receiver_uid = $2
    )
    OR (
        sender_uid = $2
        AND receiver_uid = $1
    )
ORDER BY created_at DESC;
-- name: CreateChat :one
INSERT INTO chats (sender_uid, receiver_uid, content)
VALUES ($1, $2, $3)
RETURNING *;