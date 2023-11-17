package service

import (
	db "IMChat/db/sqlc"
	"IMChat/internal/errors"
	"context"
)

type FriendGroupApply interface {
	CreateFriendGroupApply(ctx context.Context, userId, targetId int32) error
	// ReplyFriendGroupApply(ctx context.Context,userId, targetId int32) error	// 这个主要是用来判断是否已经是好友或已在群组中
}

type Friend struct {
	store db.Store
}

func (f *Friend) CreateFriendGroupApply(ctx context.Context, userId, targetId int32) error {
	// 判断用户是否存在
	user, _ := f.store.GetUserById(ctx, targetId)
	if user.ID == 0 {
		return errors.UserNotFoundErr
	}

	// 判断是否有未同意的申请记录
	count, _ := f.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
		UserID:    userId,
		TargetID:  targetId,
		ApplyType: USER,
	})
	if count > 0 {
		return errors.DuplicakeErr
	}

	// 不能重复添加
	friend, _ := f.store.GetFriend(ctx, &db.GetFriendParams{
		UserID:   userId,
		FriendID: targetId,
	})

	if friend.UserID != 0 {
		return errors.DuplicakeErr
	}
	return nil
}

type Group struct {
	store db.Store
}

func (g *Group) CreateFriendGroupApply(ctx context.Context, userId, targetId int32) error {
	// 判断群组是否存在
	user, _ := g.store.GetGroupById(ctx, targetId)
	if user.ID == 0 {
		return errors.UserNotFoundErr
	}

	// 判断是否有未同意的申请记录
	count, _ := g.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
		UserID:    userId,
		TargetID:  targetId,
		ApplyType: USER,
	})
	if count > 0 {
		return errors.DuplicakeErr
	}

	group, _ := g.store.GetGroup(ctx, &db.GetGroupParams{
		UserID:  userId,
		GroupID: targetId,
	})
	if group.ID != 0 {
		return errors.DuplicakeErr
	}

	return nil
}

func CreateFriendGroupApply(applyType int32, store db.Store) FriendGroupApply {
	switch applyType {
	case USER:
		return &Friend{store: store}
	case GROUP:
		return &Group{store: store}
	}

	return nil
}
