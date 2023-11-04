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

type FriendService struct {
	pb.UnimplementedFriendshipServiceServer
	*Server
}

func NewFriendService(server *Server) pb.FriendshipServiceServer {
	return &FriendService{
		Server: server,
	}
}

func (friendService *FriendService) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.Response, error) {

	user, err := friendService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// 判断用户是否存在
	u, _ := friendService.store.GetUserById(ctx, req.ReceiverId)
	if u.ID == 0 {
		return nil, errors.UserNotFoundErr
	}

	// 不能添加自己
	if user.ID == req.GetReceiverId() {
		return nil, status.Error(codes.InvalidArgument, "不能添加自己为好友")
	}

	// 不能重复添加
	friend, _ := friendService.store.GetFriend(ctx, &db.GetFriendParams{
		UserID:   user.ID,
		FriendID: req.ReceiverId,
	})
	if friend.UserID != 0 {
		return nil, errors.DuplicakeErr
	}

	_, err = friendService.store.AddFriendTx(ctx, &db.AddFriendTxParams{
		UserID:   user.ID,
		FriendID: req.ReceiverId,
		Comment:  req.Note,
	})
	if err != nil {
		return nil, errors.ServerErr
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendService *FriendService) UpdateFriend(ctx context.Context, req *pb.UpdateFriendRequest) (*pb.Response, error) {
	user, err := friendService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	if req.Note == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	_, err = friendService.store.UpdateFriendComment(ctx, &db.UpdateFriendCommentParams{
		Comment:  req.Note,
		UserID:   user.ID,
		FriendID: req.FriendId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to update friend note: %v", err)
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendService *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.Response, error) {
	user, err := friendService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	arg := &db.DeleteFriendTxParams{
		UserID:   user.ID,
		FriendID: req.GetFriendId(),
	}
	err = friendService.store.DeleteFriendTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete friend?: %v", err)
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (friendService *FriendService) ListFriends(req *emptypb.Empty, stream pb.FriendshipService_ListFriendsServer) error {
	ctx := stream.Context()
	user, err := friendService.authorization(ctx)
	if err != nil {
		return err
	}

	friends, _ := friendService.store.ListFriends(ctx, user.ID)
	for _, friend := range friends {
		res := &pb.ListFriendsResponse{
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
