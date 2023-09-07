package service

import (
	"IMChat/core/cache"
	"IMChat/core/config"
	"IMChat/core/token"
	db "IMChat/db/sqlc"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	conf  config.Config
	store db.Store
	maker token.Maker
	cache *redis.Client
}

func NewServer(conf config.Config, store db.Store) *Server {
	server := &Server{
		store: store,
		conf:  conf,
		maker: token.NewPasetoMaker(conf.Token),
		cache: cache.InitRedis(&conf.Redis),
	}

	return server
}
