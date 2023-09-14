package service

import (
	"IMChat/core/token"
	db "IMChat/db/sqlc"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authorizationHeader = "authorization"

func (server *Server) authorization(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	if len(authHeader) < 1 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	payload, err := server.maker.VerifyToken(authHeader)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %v", err)
	}
	return payload, nil
}

func (server *Server) getUserInfo(ctx context.Context) (*db.User, error) {
	payload, err := server.authorization(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userInfoKey := fmt.Sprintf("userInfo:%s", payload.Username)
	var user db.User
	err = server.cache.Get(ctx, userInfoKey).Scan(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cache: %v", err)
	}

	return &user, nil
}
