package service

import (
	db "IMChat/db/sqlc"
	errDefine "IMChat/internal/errors"
	"IMChat/pb"
	"errors"
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

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errDefine.ServerErr
		}

		// 判断发送者是否是当前的登录用户

		// 判断发送类型是群发还是私聊

		// 持久化消息
		arg := &db.AddMessageTxParams{
			AddMessageParams: db.AddMessageParams{
				SenderID:    req.Message.SenderId,
				ReceiverID:  req.Message.ReceiverId,
				MessageType: int16(req.Message.MessageType),
				Content:     req.Message.Content,
				SendType:    int16(req.Message.SendType),
				SendingTime: time.Now(),
			},
			AfterFunc: func(message db.Message) error {
				msgPb := converMesagge(message)
				resp := &pb.ChatMessagg{Message: msgPb}
				return stream.Send(resp)
			},
		}

		_, err = chatService.store.AddMessageTx(stream.Context(), arg)
		if err != nil {
			log.Error().Err(err).Msgf("AddMessageTx(_)")
			return errDefine.ServerErr
		}
	}
}

func (chatService *ChatService) SenderMessage(stream pb.ChatService_SenderMessageServer) error {
	ctx := stream.Context()
	user, err := chatService.authorization(ctx)
	if err != nil {
		return err
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errDefine.ServerErr
		}

		// 判断发送者是否与当前登录者一致
		if user.ID != in.SenderId {
			log.Error().Int32("userId", user.ID).Int32("senderId", in.SenderId).Msg("非当前登录用户")
			return errDefine.PermissionDeniedErr
		}

		// 判断双方是否为好友关系，只允许好友之间互发消息
		_, err = chatService.store.GetFriend(ctx, &db.GetFriendParams{
			UserID:   in.SenderId,
			FriendID: in.ReceiverId,
		})
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "not friend relation")
			}
			return errDefine.ServerErr
		}

		// 持久化消息
		arg := &db.AddMessageTxParams{
			AddMessageParams: db.AddMessageParams{
				SenderID:    user.ID,
				ReceiverID:  in.ReceiverId,
				MessageType: int16(in.MessageType),
				Content:     in.Content,
				SendType:    int16(in.SendType),
				SendingTime: time.Now(),
			},
			AfterFunc: func(message db.Message) error {
				msgPb := converMesagge(message)
				resp := &pb.SenderMessageResponse{Message: msgPb}
				return stream.Send(resp)
			},
		}
		_, err = chatService.store.AddMessageTx(ctx, arg)
		if err != nil {
			// return status.Errorf(codes.Internal, "failed to add message: %v\n", err)
			log.Error().Err(err).Msgf("%v.AddMessageTx", chatService.store)
			return errDefine.ServerErr
		}
		// msgPb := converMesagge(msg)
		// resp := &pb.SenderMessageResponse{Message: msgPb}
		// if err := stream.Send(resp); err != nil {
		// 	return status.Errorf(codes.Unknown, "failed send message: %v\n", err)
		// }
	}
	// return nil
}
