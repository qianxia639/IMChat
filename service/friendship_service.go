package service

import (
	db "IMChat/db/sqlc"
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
