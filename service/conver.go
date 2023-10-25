package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"IMChat/utils"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func converUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Nickname:          user.Nickname,
		Email:             utils.DesnsitizeEmail(user.Email),
		ProfilePictureUrl: user.ProfilePictureUrl,
		Gender:            pb.Gender(user.Gender),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func converMesagge(message db.Message) *pb.ChatMeg {
	return &pb.ChatMeg{
		SenderId:    message.SenderID,
		ReceiverId:  message.ReceiverID,
		Content:     message.Content,
		MessageType: int32(message.MessageType),
		SendType:    pb.SendType(message.SendType),
		// SendingTime: timestamppb.New(message.SendingTime),
		// ReceivTime:  timestamppb.New(message.ReceivTime),
	}
}
