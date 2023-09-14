package service

import (
	db "IMChat/db/sqlc"
	"IMChat/pb"
	"io"

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

func (messageService *MessageService) SendMessage(stream pb.MessageService_SendMessageServer) error {
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

		// 持久化消息
		arg := &db.AddMessageParams{
			SenderID:   user.ID,
			ReceiverID: in.ReceiveId,
			Content:    in.Content,
		}
		msg, err := messageService.store.AddMessage(ctx, arg)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to add message: %v\n", err)
		}
		msgPb := converMesagge(msg)
		resp := &pb.SendMessageResponse{Message: msgPb}
		if err := stream.Send(resp); err != nil {
			return status.Errorf(codes.Unknown, "failed send message: %v\n", err)
		}
	}
	return nil
}
