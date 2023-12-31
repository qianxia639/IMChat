package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	PermissionDenied = codes.PermissionDenied
	NotFound         = codes.NotFound
	AlreadyExists    = codes.AlreadyExists
)

func NewError(c codes.Code, msg string) error {
	return status.Error(c, msg)
}

func NewParamError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

func NewInternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

func NewUnauthorizaedError(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}
