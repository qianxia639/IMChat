package service

import (
	"IMChat/core/token"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func (server *Server) GrpcUnaryAuthorization(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
		// return nil, fmt.Errorf("missing metadata")
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

	// TODO 暂时先这样，目前未使用
	key := fmt.Sprintf(":%s", payload.Username)
	server.cache.Set(ctx, key, payload, 0)

	return handler(ctx, req)
}
