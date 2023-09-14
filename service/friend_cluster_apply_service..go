package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (friendClusterApplyService *FriendClusterApplyService) ListFriendApply(req *pb.EmptyRequest, stream pb.FriendClusterApplyService_ListFriendClusterApplyServer) error {
	ctx := stream.Context()
	payload, err := friendClusterApplyService.authorization(ctx)
	if err != nil {
		return unauthenticatedError(err)
	}

	var user db.User
	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	err = friendClusterApplyService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return status.Errorf(codes.Internal, "get cache error: %v", err)
	}

	// TODO 需要更改
	// friendApplys, _ := friendClusterApplyService.store.ListFriendApply(ctx, user.ID)
	// for _, friendApply := range friendApplys {
	// 	res := &pb.ListFriendClusterApplyResponse{
	// 		ReplyId:   friendApply.ApplyID,
	// 		Nickname:  friendApply.Nickname,
	// 		ApplyDesc: friendApply.ApplyDesc,
	// 		CreatedAt: timestamppb.New(friendApply.CreatedAt),
	// 	}
	// 	if err := stream.Send(res); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
