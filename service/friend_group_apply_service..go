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

func (friendGroupApplyService *FriendGroupApplyService) CreateFriendGroupApply(ctx context.Context, req *pb.CreateFriendGroupApplyRequest) (*pb.CreateFriendGroupApplyResponse, error) {

	// 身份校验并获取用户信息
	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.ReceiverId && req.ApplyType == pb.ApplyType_FRIEND {
		return nil, status.Error(codes.InvalidArgument, "无法添加自己")
	}

	friendGroupApply := CreateFriendGroupApply(req.ApplyType, friendGroupApplyService.store)
	err = friendGroupApply.CreateFriendGroupApply(ctx, user.ID, req.ReceiverId)
	if err != nil {
		return nil, err
	}

	// switch req.ApplyType {
	// case pb.ApplyType_FRIEND: // 好友
	// 	if err := friendGroupApplyService.createFriendApply(ctx, user.ID, req.ReceiverId); err != nil {
	// 		return nil, err
	// 	}

	// case pb.ApplyType_CLUSTER: // 群组
	// 	// 判断是否已申请

	// 	// 判断是否已经在群组中
	// 	return nil, status.Error(codes.Internal, "未实现")
	// }

	arg := &db.CreateFriendGroupApplyParams{
		SenderID:   user.ID,
		ReceiverID: req.GetReceiverId(),
		ApplyDesc:  req.GetApplyDesc(),
		ApplyType:  int16(req.GetApplyType()),
	}

	_, err = friendGroupApplyService.store.CreateFriendGroupApply(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create check: %v", err)
	}

	return &pb.CreateFriendGroupApplyResponse{
		Message: "Create Successfully",
	}, nil
}

func (friendGroupApplyService *FriendGroupApplyService) ReplyFriendGroupApply(ctx context.Context, req *pb.ReplyFriendGroupApplyRequest) (*pb.ReplyFriendGroupApplyResponse, error) {

	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.FriendId && req.ApplyType == pb.ApplyType_FRIEND {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己")
	}

	if pb.Status_PENDING == req.Status {
		return &pb.ReplyFriendGroupApplyResponse{Message: "Pending..."}, nil
	}

	// 判断是否是申请列表中的数据
	count, _ := friendGroupApplyService.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
		SenderID:   req.FriendId,
		ReceiverID: user.ID,
		ApplyType:  int16(req.GetApplyType()),
	})
	if count < 1 {
		log.Error().Msgf("ExistsFriendGroupApply(_)")
		return nil, errors.ParamsErr
	}

	// 判断是否已经是好友或已在群组中
	switch req.ApplyType {
	case pb.ApplyType_FRIEND:
		friend, _ := friendGroupApplyService.store.GetFriend(ctx, &db.GetFriendParams{
			UserID:   user.ID,
			FriendID: req.FriendId,
		})
		if friend.ID != 0 {
			return nil, errors.DuplicakeErr
		}
	case pb.ApplyType_CLUSTER:
		log.Error().Err(errors.DuplicakeErr).Msg("响应群组申请")
		return nil, errors.DuplicakeErr
	}

	arg := &db.ReplyFriendGroupApplyTxParams{
		UserID:    user.ID,
		FriendID:  req.GetFriendId(),
		Status:    int32(req.GetStatus()),
		ApplyType: int32(req.GetApplyType()),
		Note:      req.Note,
	}

	_, err = friendGroupApplyService.store.ReplyFriendGroupApplyTx(ctx, arg)
	if err != nil {
		// pgErr := db.ErrorCode(err)
		// if pgErr == db.ForeignKeyViolation || pgErr == db.UniqueViolation {
		// 	return nil, status.Errorf(codes.AlreadyExists, "faile add friend error: %v", err)
		// }
		log.Error().Err(err).Msgf("ReplyFriendClusterApplyTx(_)")
		return nil, errors.ServerErr
	}

	return &pb.ReplyFriendGroupApplyResponse{Message: "Successfully..."}, nil
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
			ApplyId:   friendClusterApply.SenderID,
			Nickname:  friendClusterApply.Nickname,
			ApplyDesc: friendClusterApply.ApplyDesc,
			ApplyType: int32(friendClusterApply.ApplyType),
			ApplyTime: timestamppb.New(friendClusterApply.ApplyTime),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
