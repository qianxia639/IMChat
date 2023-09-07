package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FriendApplyService struct {
	pb.UnimplementedFriendApplyServiceServer
	Server
}

func NewFriendApplyService(server Server) pb.FriendApplyServiceServer {
	return &FriendApplyService{
		Server: server,
	}
}

func (friendApplyService *FriendApplyService) CreateFriendApply(ctx context.Context, req *pb.CreateFriendApplyRequest) (*pb.CreateFriendApplyResponse, error) {
	payload, err := friendApplyService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	var user db.User
	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	err = friendApplyService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if user.ID == req.GetReplyId() {
		return nil, status.Error(codes.InvalidArgument, "无法添加自己")
	}

	// 判断是否已申请
	// TODO 此处由数据库去做，后期可能会改

	// 判断是否已经是好友
	// TODO: 暂未完成
	friend, _ := friendApplyService.store.GetFriend(ctx, &db.GetFriendParams{
		UserID:   user.ID,
		FriendID: req.GetReplyId(),
	})

	if friend.FriendID == req.ReplyId {
		return nil, status.Errorf(codes.Internal, "不能重复添加")
	}

	arg := &db.CreateFriendApplyParams{
		ApplyID:   user.ID,
		ReplyID:   req.GetReplyId(),
		ApplyDesc: req.GetApplyDesc(),
		Note:      req.GetNote(),
	}

	_, err = friendApplyService.store.CreateFriendApply(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create check: %v", err)
	}

	return &pb.CreateFriendApplyResponse{
		Message: "Create Successfully",
	}, nil
}

func (friendApplyService *FriendApplyService) ListFriendApply(req *pb.EmptyRequest, stream pb.FriendApplyService_ListFriendApplyServer) error {
	ctx := stream.Context()
	payload, err := friendApplyService.authorization(ctx)
	if err != nil {
		return unauthenticatedError(err)
	}

	var user db.User
	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	err = friendApplyService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return status.Errorf(codes.Internal, "get cache error: %v", err)
	}

	friendApplys, _ := friendApplyService.store.ListFriendApply(ctx, user.ID)
	for _, friendApply := range friendApplys {
		res := &pb.ListFriendApplyResponse{
			ReplyId:   friendApply.ApplyID,
			Nickname:  friendApply.Nickname,
			ApplyDesc: friendApply.ApplyDesc,
			CreatedAt: timestamppb.New(friendApply.CreatedAt),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
