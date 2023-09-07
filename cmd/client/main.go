package main

import (
	"IMChat/client"
	db "IMChat/db/sqlc"
	"IMChat/utils"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc dial error: ", err)
	}
	defer conn.Close()

	user := newUser()

	userClient := client.NewUserClient(conn)

	userClient.CreateUser(user)

	accessToken, err := userClient.LogigUser(user.Username, user.Password)
	if err != nil {
		log.Fatalf("failed login: %v\n", err)
	}

	ctx := attachToken(context.Background(), accessToken)

	u := userClient.GetUser(ctx)
	log.Printf("user info: %+v\n", u)
}

func attachToken(ctx context.Context, accessToken string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", accessToken)
}

func newUser() *db.User {
	return &db.User{
		Username: utils.RandomString(6),
		Email:    utils.RandomString(6) + "@email.com",
		Password: utils.RandomString(6),
		Gender:   int16(utils.RandomInt(0, 2)),
	}
}
