package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	// 无需认证直接放行

	// 认证成功
	return handler(ctx, req)
}

func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) {

}
