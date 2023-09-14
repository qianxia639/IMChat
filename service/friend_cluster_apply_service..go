package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FriendClusterApplyService struct {
	pb.UnimplementedFriendClusterApplyServiceServer
	Server
}

func NewFriendClusterApplyService(server Server) pb.FriendClusterApplyServiceServer {
	return &FriendClusterApplyService{
		Server: server,
	}
}

func (friendClusterApplyService *FriendClusterApplyService) CreateFriendClusterApply(ctx context.Context, req *pb.CreateFriendClusterApplyRequest) (*pb.CreateFriendClusterApplyResponse, error) {

	// 身份校验并获取用户信息
	user, err := friendClusterApplyService.getUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.ReceiverId && req.Flag == 0 {
		return nil, status.Error(codes.InvalidArgument, "无法添加自己")
	}

	switch req.GetFlag() {
	case 0: // 好友
		// 判断是否已申请
		count, _ := friendClusterApplyService.store.ExistsFriendClusterApply(ctx, &db.ExistsFriendClusterApplyParams{
			ApplyID:    user.ID,
			ReceiverID: req.GetReceiverId(),
			Flag:       int16(req.Flag),
		})
		if count > 0 {
			return nil, status.Errorf(codes.InvalidArgument, "请勿重复申请")
		}

		// 判断是否已经是好友
		// TODO: 暂未完成
		friend, _ := friendClusterApplyService.store.GetFriend(ctx, &db.GetFriendParams{
			UserID:   user.ID,
			FriendID: req.GetReceiverId(),
		})

		if friend.FriendID == req.ReceiverId {
			return nil, status.Errorf(codes.Internal, "不能重复添加")
		}
	case 1: // 群组
		// 判断是否已申请

		// 判断是否已经在群组中
		// TODO: 暂未完成

	}

	arg := &db.CreateFriendClsuterApplyParams{
		ApplyID:    user.ID,
		ReceiverID: req.GetReceiverId(),
		ApplyDesc:  req.GetApplyDesc(),
		Flag:       int16(req.GetFlag()),
	}

	_, err = friendClusterApplyService.store.CreateFriendClsuterApply(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create check: %v", err)
	}

	return &pb.CreateFriendClusterApplyResponse{
		Message: "Create Successfully",
	}, nil
}

func (friendClusterApplyService *FriendClusterApplyService) ListFriendClusterApply(req *pb.EmptyRequest, stream pb.FriendClusterApplyService_ListFriendClusterApplyServer) error {
	ctx := stream.Context()
	user, err := friendClusterApplyService.getUserInfo(ctx)
	if err != nil {
		return err
	}

	friendClusterApplys, _ := friendClusterApplyService.store.ListFriendClusterApply(ctx, user.ID)
	for _, friendClusterApply := range friendClusterApplys {
		res := &pb.ListFriendClusterApplyResponse{
			ApplyId:   friendClusterApply.ApplyID,
			Nickname:  friendClusterApply.Nickname,
			ApplyDesc: friendClusterApply.ApplyDesc,
			Flag:      int32(friendClusterApply.Flag),
			ApplyTime: timestamppb.New(friendClusterApply.ApplyTime),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
