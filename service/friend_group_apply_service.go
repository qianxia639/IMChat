package service

import (
	db "IMChat/db/sqlc"
	"IMChat/internal/errors"
	"IMChat/pb"
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FriendGroupApplyService struct {
	pb.UnimplementedFriendGroupApplyServiceServer
	*Server
}

func NewFriendGroupRequestService(server *Server) pb.FriendGroupApplyServiceServer {
	return &FriendGroupApplyService{
		Server: server,
	}
}

func (friendGroupApplyService *FriendGroupApplyService) CreateFriendGroupApply(ctx context.Context, req *pb.CreateFriendGroupApplyRequest) (*pb.Response, error) {

	// 身份校验并获取用户信息
	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.TargetId && req.ApplyType == USER {
		return nil, status.Error(codes.InvalidArgument, "无法添加自己")
	}

	// 好友或群组的逻辑判断
	friendGroupApply := CreateFriendGroupApply(req.ApplyType, friendGroupApplyService.store)
	err = friendGroupApply.CreateFriendGroupApply(ctx, user.ID, req.TargetId)
	if err != nil {
		return nil, err
	}

	arg := &db.CreateFriendGroupApplyParams{
		UserID:      user.ID,
		TargetID:    req.TargetId,
		Description: req.Description,
		ApplyType:   int16(req.GetApplyType()),
	}

	_, err = friendGroupApplyService.store.CreateFriendGroupApply(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create check: %v", err)
	}

	return &pb.Response{
		Message: "Create Successfully",
	}, nil
}

const (
	// 申请类型
	USER  = 1 // 好友
	GROUP = 2 // 群组

	PENDING  = 1 // 待确认
	ACCEPTEd = 2 // 已确认
	REJECTED = 3 // 已拒绝
)

func (friendGroupApplyService *FriendGroupApplyService) ReplyFriendGroupApply(ctx context.Context, req *pb.ReplyFriendGroupApplyRequest) (*pb.Response, error) {

	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.FriendId && req.ApplyType == USER {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己")
	}

	if req.Status == PENDING {
		return &pb.Response{Message: "Pending..."}, nil
	}

	// 判断是否是申请列表中的数据
	count, _ := friendGroupApplyService.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
		UserID:    req.FriendId,
		TargetID:  user.ID,
		ApplyType: int16(req.GetApplyType()),
	})
	if count < 1 {
		log.Error().Msgf("ExistsFriendGroupApply(_)")
		return nil, errors.ParamsErr
	}

	// 判断是否已经是好友或已在群组中
	switch req.ApplyType {
	case USER:
		friend, _ := friendGroupApplyService.store.GetFriend(ctx, &db.GetFriendParams{
			UserID:   user.ID,
			FriendID: req.FriendId,
		})
		if friend.UserID != 0 {
			return nil, errors.DuplicakeErr
		}
	case GROUP:
		log.Error().Err(errors.DuplicakeErr).Msg("响应群组申请")
		return nil, errors.DuplicakeErr
	}

	arg := &db.ReplyFriendGroupApplyTxParams{
		UserID:    user.ID,
		FriendID:  req.GetFriendId(),
		Status:    int32(req.GetStatus()),
		ApplyType: int32(req.GetApplyType()),
		Comment:   req.Comment,
	}

	_, err = friendGroupApplyService.store.ReplyFriendGroupApplyTx(ctx, arg)
	if err != nil {
		log.Error().Err(err).Msgf("ReplyFriendClusterApplyTx(_)")
		return nil, errors.ServerErr
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendGroupApplyService *FriendGroupApplyService) ListFriendGroupApply(req *emptypb.Empty, stream pb.FriendGroupApplyService_ListFriendGroupApplyServer) error {
	ctx := stream.Context()
	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return err
	}

	friendClusterApplys, _ := friendGroupApplyService.store.ListFriendGroupApply(ctx, user.ID)
	for _, friendClusterApply := range friendClusterApplys {
		res := &pb.ListFriendGroupApplyResponse{
			ApplyId:     friendClusterApply.UserID,
			Nickname:    friendClusterApply.Nickname,
			Description: friendClusterApply.Description,
			ApplyType:   int32(friendClusterApply.ApplyType),
			CreatedAt:   timestamppb.New(friendClusterApply.CreatedAt),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
