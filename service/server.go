package service

import (
	"IMChat/core/cache"
	"IMChat/core/config"
	"IMChat/core/token"
	db "IMChat/db/sqlc"
	"IMChat/internal/email"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	conf      config.Config
	store     db.Store
	maker     token.Maker
	cache     *redis.Client
	emailPool sync.Pool
}

func NewServer(conf config.Config, store db.Store) *Server {
	server := &Server{
		store: store,
		conf:  conf,
		maker: token.NewPasetoMaker(conf.Token),
		cache: cache.InitRedis(&conf.Redis),
	}

	server.emailPool.New = func() interface{} {
		return &email.Email{
			Username: server.conf.Email.Username,
			Password: server.conf.Email.Password,
			Host:     server.conf.Email.Host,
		}
	}

	return server
}
