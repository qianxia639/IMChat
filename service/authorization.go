package service

import (
	db "IMChat/db/sqlc"
	"IMChat/internal/errors"
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

const authorizationHeader = "authorization"

func (server *Server) authorization(ctx context.Context) (*db.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error().Msgf("missing meetadata")
		return nil, errors.UnauthorizaedErr
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		log.Error().Msgf("missing authorization header")
		return nil, errors.UnauthorizaedErr
	}

	authHeader := values[0]
	if len(authHeader) < 1 {
		log.Error().Msgf("invalid authorization header format")
		return nil, errors.UnauthorizaedErr
	}

	payload, err := server.maker.VerifyToken(authHeader)
	if err != nil {
		log.Error().Err(err).Msgf("invalid access token")
		return nil, errors.UnauthorizaedErr
	}

	var user db.User
	err = server.cache.Get(ctx, getUserInfoKey(payload.Username)).Scan(&user)
	if err != nil {
		log.Error().Err(err)
		return nil, errors.ServerErr
	}

	return &user, nil
}

// func (server *Server) getUserInfo(ctx context.Context) (*db.User, error) {
// 	payload, err := server.authorization(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user db.User
// 	err = server.cache.Get(ctx, getUserInfoKey(payload.Username)).Scan(&user)
// 	if err != nil {
// 		log.Error().Err(err)
// 		return nil, errors.ServerErr
// 	}

// 	return &user, nil
// }
