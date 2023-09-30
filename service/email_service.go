package service

import (
	"IMChat/internal/errors"
	"IMChat/pb"
	"IMChat/utils"
	"context"
)

type EmailService struct {
	pb.UnimplementedEmailServiceServer
	Server
}

func NewEmailService(server Server) pb.EmailServiceServer {
	return &EmailService{
		Server: server,
	}
}

func (emailService *EmailService) SendEmailCode(ctx context.Context, req *pb.SendEmailCodeRequest) (*pb.SendEmailCodeResponse, error) {

	e := utils.Email{
		Username: emailService.conf.Email.Username,
		Password: emailService.conf.Email.Password,
		Host:     emailService.conf.Email.Host,
	}
	err := e.SendEmailCode(emailService.cache, req.Email)
	if err != nil {
		return nil, errors.SendEmailCodeErr
	}

	return &pb.SendEmailCodeResponse{Message: "send email code successfully"}, nil
}
