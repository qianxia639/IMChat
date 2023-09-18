package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"IMChat/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func converUser(user db.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     utils.DesnsitizeEmail(user.Email),
		Avatar:    user.Avatar,
		Gender:    pb.Gender(pb.Gender(user.Gender).Number()),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func converMesagge(message db.Message) *pb.Message {
	return &pb.Message{
		SenderId:     message.SenderID,
		ReceiverId:   message.ReceiverID,
		Content:      message.Content,
		SenderTime:   timestamppb.New(message.SenderTime),
		ReceiverTime: timestamppb.New(message.ReceiverTime),
	}
}
