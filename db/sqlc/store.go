package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	ReplyFriendGroupApplyTx(ctx context.Context, arg *ReplyFriendGroupApplyTxParams) (Friendship, error)
	DeleteFriendTx(ctx context.Context, arg *DeleteFriendTxParams) error
	AddMessageTx(ctx context.Context, arg *AddMessageTxParams) (AddMessageTxResult, error)
	DeleteUserTx(ctx context.Context, userId int32) error
	GroupTx
}

type GroupTx interface {
	CreateGroupTx(ctx context.Context, arg *CreateGroupTxParams) error
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
