// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: message.sql

package db

import (
	"context"
	"time"
)

const addMessage = `-- name: AddMessage :one
INSERT INTO messages (
    sender_id, receiver_id, message_type, content, send_type, sending_time
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, sender_id, receiver_id, message_type, content, send_type, sending_time, receiv_time
`

type AddMessageParams struct {
	SenderID    int32     `json:"sender_id"`
	ReceiverID  int32     `json:"receiver_id"`
	MessageType int16     `json:"message_type"`
	Content     string    `json:"content"`
	SendType    int16     `json:"send_type"`
	SendingTime time.Time `json:"sending_time"`
}

func (q *Queries) AddMessage(ctx context.Context, arg *AddMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, addMessage,
		arg.SenderID,
		arg.ReceiverID,
		arg.MessageType,
		arg.Content,
		arg.SendType,
		arg.SendingTime,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.MessageType,
		&i.Content,
		&i.SendType,
		&i.SendingTime,
		&i.ReceivTime,
	)
	return i, err
}
