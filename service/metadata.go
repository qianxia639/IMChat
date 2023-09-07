package service

import (
	"context"
	"net"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// 获取客户端user-agent
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// 获取HTTP user-agent
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// // 获取HTTP IP
		if clientIp := md.Get(xForwardedForHeader); len(clientIp) > 0 {
			mtdt.ClientIp = clientIp[0]
		}
	}

	// // 获取客户端Ip
	if p, ok := peer.FromContext(ctx); ok {
		host, _, _ := net.SplitHostPort(p.Addr.String())
		mtdt.ClientIp = host
	}

	return mtdt
}
