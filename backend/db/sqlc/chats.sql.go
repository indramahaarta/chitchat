// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: chats.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const createChat = `-- name: CreateChat :one
INSERT INTO chats (sender_uid, receiver_uid, content)
VALUES ($1, $2, $3)
RETURNING id, content, created_at, sender_uid, receiver_uid
`

type CreateChatParams struct {
	SenderUid   uuid.UUID `json:"sender_uid"`
	ReceiverUid uuid.UUID `json:"receiver_uid"`
	Content     string    `json:"content"`
}

func (q *Queries) CreateChat(ctx context.Context, arg CreateChatParams) (Chats, error) {
	row := q.db.QueryRowContext(ctx, createChat, arg.SenderUid, arg.ReceiverUid, arg.Content)
	var i Chats
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedAt,
		&i.SenderUid,
		&i.ReceiverUid,
	)
	return i, err
}

const getChatsDetails = `-- name: GetChatsDetails :many
SELECT chats.id, chats.content, chats.created_at, chats.sender_uid, chats.receiver_uid,
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
ORDER BY created_at DESC
`

type GetChatsDetailsParams struct {
	SenderUid   uuid.UUID `json:"sender_uid"`
	ReceiverUid uuid.UUID `json:"receiver_uid"`
}

type GetChatsDetailsRow struct {
	ID          uuid.UUID       `json:"id"`
	Content     string          `json:"content"`
	CreatedAt   time.Time       `json:"created_at"`
	SenderUid   uuid.UUID       `json:"sender_uid"`
	ReceiverUid uuid.UUID       `json:"receiver_uid"`
	Sender      json.RawMessage `json:"sender"`
	Receiver    json.RawMessage `json:"receiver"`
}

func (q *Queries) GetChatsDetails(ctx context.Context, arg GetChatsDetailsParams) ([]GetChatsDetailsRow, error) {
	rows, err := q.db.QueryContext(ctx, getChatsDetails, arg.SenderUid, arg.ReceiverUid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetChatsDetailsRow{}
	for rows.Next() {
		var i GetChatsDetailsRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreatedAt,
			&i.SenderUid,
			&i.ReceiverUid,
			&i.Sender,
			&i.Receiver,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChatsHistories = `-- name: GetChatsHistories :many
SELECT c.sender_uid,
    sender.name AS sender_name,
    c.receiver_uid,
    receiver.name AS receiver_name,
    c.created_at AS latest_created_at,
    c.content AS latest_content
FROM (
        SELECT id, content, created_at, sender_uid, receiver_uid,
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
ORDER BY latest_created_at DESC
`

type GetChatsHistoriesRow struct {
	SenderUid       uuid.UUID      `json:"sender_uid"`
	SenderName      sql.NullString `json:"sender_name"`
	ReceiverUid     uuid.UUID      `json:"receiver_uid"`
	ReceiverName    sql.NullString `json:"receiver_name"`
	LatestCreatedAt time.Time      `json:"latest_created_at"`
	LatestContent   string         `json:"latest_content"`
}

func (q *Queries) GetChatsHistories(ctx context.Context, senderUid uuid.UUID) ([]GetChatsHistoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getChatsHistories, senderUid)
	if err != nil {
		return nil, err
	} 
	defer rows.Close()
	items := []GetChatsHistoriesRow{}
	for rows.Next() {
		var i GetChatsHistoriesRow
		if err := rows.Scan(
			&i.SenderUid,
			&i.SenderName,
			&i.ReceiverUid,
			&i.ReceiverName,
			&i.LatestCreatedAt,
			&i.LatestContent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
