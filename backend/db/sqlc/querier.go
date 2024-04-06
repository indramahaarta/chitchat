// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateChat(ctx context.Context, arg CreateChatParams) (Chats, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	GetAllUser(ctx context.Context, uid uuid.UUID) ([]Users, error)
	GetChatsDetails(ctx context.Context, arg GetChatsDetailsParams) ([]GetChatsDetailsRow, error)
	GetChatsHistories(ctx context.Context, senderUid uuid.UUID) ([]GetChatsHistoriesRow, error)
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserById(ctx context.Context, uid uuid.UUID) (Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
