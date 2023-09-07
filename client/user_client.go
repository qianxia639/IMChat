package client

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"context"
	"log"
	"time"
	"unsafe"

	"google.golang.org/grpc"
)

type UserClient struct {
	service pb.UserServiceClient
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	service := pb.NewUserServiceClient(cc)
	return &UserClient{service: service}
}

func (userClient *UserClient) LogigUser(username, password string) (string, error) {
	req := &pb.LoginUserRequest{
		Username: username,
		Password: password,
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	res, err := userClient.service.LoginUser(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

func (userClient *UserClient) CreateUser(user *db.User) {
	req := &pb.CreateUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Gender:   (*pb.Gender)(unsafe.Pointer(&user.Gender)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userClient.service.CreateUser(ctx, req)
	if err != nil {
		log.Printf("cannot create user: %v\n", err)
		return
	}

	log.Println(res.Message)
}

func (userClient *UserClient) GetUser(c context.Context) *pb.User {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	res, err := userClient.service.GetUser(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("get user err: %v\n", err)
	}

	return res.User
}
