package service

import (
	db "IMChat/db/sqlc"
	"IMChat/internal/errors"
	"IMChat/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FriendshipService struct {
	pb.UnimplementedFriendshipServiceServer
	*Server
}

func NewFriendService(server *Server) pb.FriendshipServiceServer {
	return &FriendshipService{
		Server: server,
	}
}

func (friendshipService *FriendshipService) ApplyFriendship(ctx context.Context, req *pb.ApplyFriendshipRequest) (*pb.Response, error) {
	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// 判断用户是否存在
	u, _ := friendshipService.store.GetUserById(ctx, req.TargetId)
	if u.ID == 0 {
		return nil, errors.UserNotFoundErr
	}

	// 不能添加自己
	if user.ID == req.GetTargetId() {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己为好友")
	}

	// 判断是否已经是好友

	// 判断是否重复申请

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendshipService *FriendshipService) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.Response, error) {

	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// 判断用户是否存在
	u, _ := friendshipService.store.GetUserById(ctx, req.ReceiverId)
	if u.ID == 0 {
		return nil, errors.UserNotFoundErr
	}

	// 不能添加自己
	if user.ID == req.GetReceiverId() {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己为好友")
	}

	// 不能重复添加
	friend, _ := friendshipService.store.GetFriend(ctx, &db.GetFriendParams{
		UserID:   user.ID,
		FriendID: req.ReceiverId,
	})
	if friend.UserID != 0 {
		return nil, errors.DuplicakeErr
	}

	_, err = friendshipService.store.AddFriendTx(ctx, &db.AddFriendTxParams{
		UserID:   user.ID,
		FriendID: req.ReceiverId,
		Comment:  req.Comment,
	})
	if err != nil {
		return nil, errors.ServerErr
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendshipService *FriendshipService) ListFriendshipPending(req *emptypb.Empty, stream pb.FriendshipService_ListFriendshipPendingServer) error {
	ctx := stream.Context()
	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return err
	}

	pendings, _ := friendshipService.store.ListFriendshipPending(ctx, user.ID)
	for _, pending := range pendings {
		res := &pb.ListFriendshipPendingResponse{
			UserId:   pending.UserID,
			Status:   int32(pending.Status),
			Nickname: pending.Nickname,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (friendshipService *FriendshipService) UpdateFriend(ctx context.Context, req *pb.UpdateFriendRequest) (*pb.Response, error) {
	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if req.Note == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	_, err = friendshipService.store.UpdateFriendComment(ctx, &db.UpdateFriendCommentParams{
		Comment:  req.Note,
		UserID:   user.ID,
		FriendID: req.FriendId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to update friend note: %v", err)
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendshipService *FriendshipService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.Response, error) {
	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	arg := &db.DeleteFriendTxParams{
		UserID:   user.ID,
		FriendID: req.GetFriendId(),
	}
	err = friendshipService.store.DeleteFriendTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete friend?: %v", err)
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendshipService *FriendshipService) ListFriend(req *emptypb.Empty, stream pb.FriendshipService_ListFriendServer) error {
	ctx := stream.Context()
	user, err := friendshipService.authorization(ctx)
	if err != nil {
		return err
	}

	friends, _ := friendshipService.store.ListFriends(ctx, user.ID)
	for _, friend := range friends {
		res := &pb.ListFriendResponse{
			FriendId: friend.FriendID,
			Note:     friend.Comment,
			Avatar:   friend.ProfilePictureUrl,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
