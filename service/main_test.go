package service

import (
	"IMChat/core/config"
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"IMChat/utils"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	conf := config.Config{
		Token: config.Token{
			TokenSymmetricKey:   utils.RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	server := NewServer(conf, store)

	return server
}

func newTestUserService(t *testing.T, store db.Store) pb.UserServiceServer {
	server := newTestServer(t, store)
	return NewUserService(server)
}
