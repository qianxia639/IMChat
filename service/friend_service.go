package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FriendService struct {
	pb.UnimplementedFriendServiceServer
	Server
}

func NewFriendService(server Server) pb.FriendServiceServer {
	return &FriendService{
		Server: server,
	}
}

func (friendService *FriendService) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.AddFriendResponse, error) {

	user, err := friendService.getUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	if user.ID == req.GetFriendId() {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己为好友")
	}

	// TODO 需要更改
	// friend, err := friendService.store.GetFriendApply(ctx, &db.GetFriendApplyParams{
	// 	ApplyID: req.GetFriendId(),
	// 	ReplyID: user.ID,
	// })
	// if friend.ReplyID != user.ID && err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to get friend apply: %v", err)
	// }

	// 判断是否是申请列表中的数据
	count, _ := friendService.store.ExistsFriendClusterApply(ctx, &db.ExistsFriendClusterApplyParams{
		ApplyID:    req.FriendId,
		ReceiverID: user.ID,
		Flag:       0,
	})
	if count < 1 {
		return nil, status.Errorf(codes.Internal, "数据不匹配")
	}

	arg := &db.AddFriendTxParams{
		UserID:   user.ID,
		FriendID: req.GetFriendId(),
		Status:   req.GetStatus(),
		Note:     req.Note,
	}

	_, err = friendService.store.AddFriendTx(ctx, arg)
	if err != nil {
		pgErr := db.ErrorCode(err)
		if pgErr == db.ForeignKeyViolation || pgErr == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "faile add friend error: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "filaed to add friend: %v", err)
	}

	return &pb.AddFriendResponse{Message: "Successfully..."}, nil
}

func (friendService *FriendService) UpdateFriend(ctx context.Context, req *pb.UpdateFriendRequest) (*pb.UpdateFriendResponse, error) {
	payload, err := friendService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if req.Note == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	err = friendService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to get cache: %v", err)
	}

	_, err = friendService.store.UpdateFriendNote(ctx, &db.UpdateFriendNoteParams{
		Note:     req.Note,
		UserID:   user.ID,
		FriendID: req.FriendId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to update friend note: %v", err)
	}

	return &pb.UpdateFriendResponse{Message: "Successfully..."}, nil
}

func (friendService *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	payload, err := friendService.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	err = friendService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to get cache: %v", err)
	}

	arg := &db.DeleteFriendTxParams{
		UserID:   user.ID,
		FriendID: req.GetFriendId(),
	}
	err = friendService.store.DeleteFriendTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete friend?: %v", err)
	}

	return &pb.DeleteFriendResponse{Message: "Successfully..."}, nil
}

func (friendService *FriendService) ListFriends(req *pb.EmptyRequest, stream pb.FriendService_ListFriendsServer) error {
	ctx := stream.Context()
	payload, err := friendService.authorization(ctx)
	if err != nil {
		return unauthenticatedError(err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	err = friendService.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return status.Errorf(codes.Internal, "faile to get cache: %v", err)
	}

	friends, _ := friendService.store.ListFriends(ctx, user.ID)
	for _, friend := range friends {
		res := &pb.ListFriendsResponse{
			FriendId: friend.FriendID,
			Note:     friend.Note,
			Avatar:   friend.Avatar,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
