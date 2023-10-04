package service

import (
	"IMChat/internal/errors"
	"IMChat/pb"
	"context"
)

type EmailService struct {
	pb.UnimplementedEmailServiceServer
	*Server
}

func NewEmailService(server *Server) pb.EmailServiceServer {
	return &EmailService{
		Server: server,
	}
}

func (emailService *EmailService) SendEmailCode(ctx context.Context, req *pb.SendEmailCodeRequest) (*pb.SendEmailCodeResponse, error) {

	err := emailService.mail.SendEmailCode(emailService.cache, req.Email)
	if err != nil {
		return nil, errors.SendEmailCodeErr
	}

	return &pb.SendEmailCodeResponse{Message: "send email code successfully"}, nil
}
