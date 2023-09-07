package main

import (
	"IMChat/core/config"
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"IMChat/service"
	"IMChat/service/interceptor"
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if conf.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	dbConn, err := sql.Open(conf.Postgres.Driver, conf.Postgres.Source)
	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to db")
	}

	runDBMigration(conf.Postgres.MigrateUrl, conf.Postgres.Source)

	store := db.NewStore(dbConn)

	go runGatewayServer(conf, store)
	runGrpcServer(conf, store)
}

func runDBMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}
	log.Info().Msg("db migrated successfully")
}

func runGrpcServer(conf config.Config, store db.Store) {

	server := service.NewServer(conf, store)
	userService := service.NewUserService(*server)
	friendApplyService := service.NewFriendApplyService(*server)
	friendService := service.NewFriendService(*server)

	grpcLogger := grpc.UnaryInterceptor(interceptor.GrpcUnaryLogger)
	grpcServer := grpc.NewServer(
		grpcLogger,
		grpc.StreamInterceptor(interceptor.GrpcStreamLogger),
	)

	pb.RegisterUserServiceServer(grpcServer, userService)
	pb.RegisterFriendApplyServiceServer(grpcServer, friendApplyService)
	pb.RegisterFriendServiceServer(grpcServer, friendService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", conf.Server.GrpcServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(conf config.Config, store db.Store) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := service.NewServer(conf, store)

	opts := getGatewayServerMuxOptions()

	grpcMux := runtime.NewServeMux(opts...)
	// dialOpts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }
	// err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, grpcMux, conf.Server.GrpcServerAddress, dialOpts)
	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, service.NewUserService(*server))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	pb.RegisterFriendApplyServiceHandlerServer(ctx, grpcMux, service.NewFriendApplyService(*server))
	pb.RegisterFriendServiceHandlerServer(ctx, grpcMux, service.NewFriendService(*server))

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux) // 覆盖所有路由

	listener, err := net.Listen("tcp", conf.Server.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := interceptor.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}

func getGatewayServerMuxOptions() []runtime.ServeMuxOption {

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	opts := []runtime.ServeMuxOption{
		jsonOption,
	}

	return opts
}
