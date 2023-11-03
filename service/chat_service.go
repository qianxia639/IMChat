package service

import (
	db "IMChat/db/sqlc"
	errDefine "IMChat/internal/errors"
	"IMChat/pb"
	"io"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	*Server
}

func NewChatService(server *Server) pb.ChatServiceServer {
	return &ChatService{Server: server}
}

func (chatService *ChatService) ChatMessage(stream pb.ChatService_ChatMessageServer) error {
	ctx := stream.Context()
	user, err := chatService.authorization(ctx)
	if err != nil {
		return err
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errDefine.ServerErr
		}

		// 判断发送者是否与当前登录者一致
		if user.ID != req.SenderId {
			return errDefine.UserNotSameErr
		}

		// 判断发送类型是群发还是私聊
		switch req.SendType {
		case pb.SendType_USER:
			// 判断接收者是否存在
			u, _ := chatService.store.GetUserById(ctx, req.ReceiverId)
			if u.ID < 1 {
				return errDefine.UserNotFoundErr
			}
			log.Info().Msgf("私聊成功")
			//

		case pb.SendType_GROUP:
			return status.Error(codes.Internal, "群聊暂未实现")
		}

		// 持久化消息
		arg := &db.AddMessageTxParams{
			AddMessageParams: db.AddMessageParams{
				SenderID:    req.SenderId,
				ReceiverID:  req.ReceiverId,
				MessageType: int16(req.MessageType),
				Content:     req.Content,
				SendType:    int16(req.SendType) + 1,
				SendingTime: time.Now(),
			},
			AfterFunc: func(message db.Message) error {
				msgPb := converMesagge(message)
				return stream.Send(msgPb)
			},
		}

		_, err = chatService.store.AddMessageTx(stream.Context(), arg)
		if err != nil {
			log.Error().Err(err).Msgf("AddMessageTx(_)")
			return errDefine.ServerErr
		}
	}
}
