package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func converUser(user db.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Gender:    pb.Gender(pb.Gender(user.Gender).Number()),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func converMesagge(message db.Message) *pb.Message {
	return &pb.Message{
		SendId:      message.SendID,
		ReceiveId:   message.ReceiveID,
		Content:     message.Content,
		SendTime:    timestamppb.New(message.SendTime),
		ReceiveTime: timestamppb.New(message.ReceiveTime),
	}
}
