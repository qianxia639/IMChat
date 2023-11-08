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
)

type GroupService struct {
	pb.UnimplementedGroupServiceServer
	*Server
}

func NewGroupService(server *Server) pb.GroupServiceServer {
	return &GroupService{
		Server: server,
	}
}

func (groupService *GroupService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.Response, error) {
	user, err := groupService.authorization(ctx)
	if err != nil {
		return nil, err
	}

	// 校验身份
	if user.ID != req.CreatorId {
		return nil, errors.PermissionDeniedErr
	}

	arg := &db.CreateGroupTxParams{
		Nickname: user.Nickname,
		CreateGroupParams: db.CreateGroupParams{
			CreatorID:   req.CreatorId,
			GroupName:   req.GroupName,
			Icon:        req.Icon,
			Description: req.Description,
		},
	}
	err = groupService.store.CreateGroupTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.Internal, "群组已存在")
		}
		log.Error().Err(err).Msgf("CreateGroupTx(_)")
		return nil, errors.ServerErr
	}

	return &pb.Response{Message: "Successfully..."}, nil
}

func (groupService *GroupService) ListGroup(empty *emptypb.Empty, stream pb.GroupService_ListGroupServer) error {
	ctx := stream.Context()
	user, err := groupService.authorization(ctx)
	if err != nil {
		return err
	}

	groups, _ := groupService.store.ListGroup(ctx, user.ID)
	for _, group := range groups {
		res := &pb.ListGroupResponse{
			GroupId:   group.GroupID,
			UserId:    group.UserID,
			GroupName: group.GroupName,
			Username:  group.Username,
			Nickname:  group.Nickname,
			Icon:      group.Icon,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}
