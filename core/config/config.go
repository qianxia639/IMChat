package config

import "time"

type Config struct {
	Environment string   `mapstructure:"environment"`
	Postgres    Postgres `mapstructure:"postgres"`
	Server      Server   `mapstructure:"server"`
	Token       Token    `mapstructure:"token"`
	Redis       Redis    `mapstructure:"redis"`
}

type Postgres struct {
	Source     string `mapstructure:"source"`
	MigrateUrl string `mapstructure:"migrate_url"`
}

type Server struct {
	GrpcServerAddress string `mapstructure:"grpc_server_address"`
	HTTPServerAddress string `mapstructure:"http_server_address"`
}

type Token struct {
	TokenSymmetricKey   string        `mapstructure:"token_symmetric_key"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
