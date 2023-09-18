package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"errors"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageService struct {
	pb.UnimplementedMessageServiceServer
	Server
}

func NewMessageService(server Server) pb.MessageServiceServer {
	return &MessageService{Server: server}
}

func (messageService *MessageService) SenderMessage(stream pb.MessageService_SenderMessageServer) error {
	ctx := stream.Context()
	user, err := messageService.getUserInfo(ctx)
	if err != nil {
		return err
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v\n", err)
		}

		// 判断发送者是否与当前登录者一致
		if user.ID != in.SenderId {
			return status.Errorf(codes.InvalidArgument, "sneder id error")
		}

		// 判断双方是否为好友关系，只允许好友之间互发消息
		_, err = messageService.store.GetFriend(ctx, &db.GetFriendParams{
			UserID:   in.SenderId,
			FriendID: in.ReceiveId,
		})
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "not friend relation")
			}
			return status.Errorf(codes.Internal, "failed to find friend")
		}

		// 持久化消息
		arg := &db.AddMessageTxParams{
			AddMessageParams: db.AddMessageParams{
				SenderID:   user.ID,
				ReceiverID: in.ReceiveId,
				Content:    in.Content,
				SenderTime: time.Now(),
			},
			AfterFunc: func(message db.Message) error {
				msgPb := converMesagge(message)
				resp := &pb.SenderMessageResponse{Message: msgPb}
				return stream.Send(resp)
			},
		}
		_, err = messageService.store.AddMessageTx(ctx, arg)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to add message: %v\n", err)
		}
		// msgPb := converMesagge(msg)
		// resp := &pb.SenderMessageResponse{Message: msgPb}
		// if err := stream.Send(resp); err != nil {
		// 	return status.Errorf(codes.Unknown, "failed send message: %v\n", err)
		// }
	}
	return nil
}
