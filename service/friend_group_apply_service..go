package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (friendGroupApplyService *FriendGroupApplyService) CreateFriendGrpupApply(ctx context.Context, req *pb.CreateFriendGroupApplyRequest) (*pb.CreateFriendGroupApplyResponse, error) {

	// 身份校验并获取用户信息
	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.ReceiverId && req.ApplyType == 0 {
		return nil, status.Error(codes.InvalidArgument, "无法添加自己")
	}

	switch req.ApplyType {
	case 0: // 好友
		// 判断是否已申请
		count, _ := friendGroupApplyService.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
			SenderID:   user.ID,
			ReceiverID: req.GetReceiverId(),
			ApplyType:  int16(req.ApplyType),
		})
		if count > 0 {
			return nil, status.Errorf(codes.InvalidArgument, "请勿重复申请")
		}

		// 判断是否已经是好友
		// TODO: 暂未完成
		friend, _ := friendGroupApplyService.store.GetFriend(ctx, &db.GetFriendParams{
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

	if pb.Status_PENDING == req.Status {
		return &pb.ReplyFriendGroupApplyResponse{Message: "Waiting..."}, nil
	}

	user, err := friendGroupApplyService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.GetFriendId() && req.ApplyType == pb.ApplyType_FRIEND {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己为好友")
	}

	// 判断该申请是否存在与申请表中

	// 判断是否已经是好友或意在群组中

	// 添加记录

	// TODO 需要更改
	// friend, err := friendService.store.GetFriendApply(ctx, &db.GetFriendApplyParams{
	// 	ApplyID: req.GetFriendId(),
	// 	ReplyID: user.ID,
	// })
	// if friend.ReplyID != user.ID && err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to get friend apply: %v", err)
	// }

	// 判断是否是申请列表中的数据
	count, _ := friendGroupApplyService.store.ExistsFriendGroupApply(ctx, &db.ExistsFriendGroupApplyParams{
		SenderID:   req.FriendId,
		ReceiverID: user.ID,
		ApplyType:  int16(req.GetApplyType()),
	})
	if count < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "非法的数据")
	}

	arg := &db.ReplyFriendClusterApplyTxParams{
		UserID:   user.ID,
		FriendID: req.GetFriendId(),
		Status:   int32(req.GetStatus()),
		Flag:     int32(req.GetApplyType()),
		Note:     req.Note,
	}

	_, err = friendGroupApplyService.store.ReplyFriendClusterApplyTx(ctx, arg)
	if err != nil {
		pgErr := db.ErrorCode(err)
		if pgErr == db.ForeignKeyViolation || pgErr == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "faile add friend error: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "filaed to add friend: %v", err)
	}

	return &pb.ReplyFriendGroupApplyResponse{Message: "Successfully..."}, nil
}

func (friendGroupApplyService *FriendGroupApplyService) ListFriendGroupApply(req *pb.EmptyRequest, stream pb.FriendGroupApplyService_ListFriendGroupApplyServer) error {
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
