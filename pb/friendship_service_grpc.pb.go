// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: friendship_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FriendshipServiceClient is the client API for FriendshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendshipServiceClient interface {
	// 更新好友信息
	UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*Response, error)
	// 删除好友
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*Response, error)
	// 获取好友列表
	ListFriend(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FriendshipService_ListFriendClient, error)
}

type friendshipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendshipServiceClient(cc grpc.ClientConnInterface) FriendshipServiceClient {
	return &friendshipServiceClient{cc}
}

func (c *friendshipServiceClient) UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/qianxia.IMChat.FriendshipService/UpdateFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/qianxia.IMChat.FriendshipService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendshipServiceClient) ListFriend(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FriendshipService_ListFriendClient, error) {
	stream, err := c.cc.NewStream(ctx, &FriendshipService_ServiceDesc.Streams[0], "/qianxia.IMChat.FriendshipService/ListFriend", opts...)
	if err != nil {
		return nil, err
	}
	x := &friendshipServiceListFriendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FriendshipService_ListFriendClient interface {
	Recv() (*ListFriendResponse, error)
	grpc.ClientStream
}

type friendshipServiceListFriendClient struct {
	grpc.ClientStream
}

func (x *friendshipServiceListFriendClient) Recv() (*ListFriendResponse, error) {
	m := new(ListFriendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FriendshipServiceServer is the server API for FriendshipService service.
// All implementations must embed UnimplementedFriendshipServiceServer
// for forward compatibility
type FriendshipServiceServer interface {
	// 更新好友信息
	UpdateFriend(context.Context, *UpdateFriendRequest) (*Response, error)
	// 删除好友
	DeleteFriend(context.Context, *DeleteFriendRequest) (*Response, error)
	// 获取好友列表
	ListFriend(*emptypb.Empty, FriendshipService_ListFriendServer) error
	mustEmbedUnimplementedFriendshipServiceServer()
}

// UnimplementedFriendshipServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendshipServiceServer struct {
}

func (UnimplementedFriendshipServiceServer) UpdateFriend(context.Context, *UpdateFriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedFriendshipServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendshipServiceServer) ListFriend(*emptypb.Empty, FriendshipService_ListFriendServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFriend not implemented")
}
func (UnimplementedFriendshipServiceServer) mustEmbedUnimplementedFriendshipServiceServer() {}

// UnsafeFriendshipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendshipServiceServer will
// result in compilation errors.
type UnsafeFriendshipServiceServer interface {
	mustEmbedUnimplementedFriendshipServiceServer()
}

func RegisterFriendshipServiceServer(s grpc.ServiceRegistrar, srv FriendshipServiceServer) {
	s.RegisterService(&FriendshipService_ServiceDesc, srv)
}

func _FriendshipService_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qianxia.IMChat.FriendshipService/UpdateFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).UpdateFriend(ctx, req.(*UpdateFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendshipServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qianxia.IMChat.FriendshipService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendshipServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendshipService_ListFriend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FriendshipServiceServer).ListFriend(m, &friendshipServiceListFriendServer{stream})
}

type FriendshipService_ListFriendServer interface {
	Send(*ListFriendResponse) error
	grpc.ServerStream
}

type friendshipServiceListFriendServer struct {
	grpc.ServerStream
}

func (x *friendshipServiceListFriendServer) Send(m *ListFriendResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FriendshipService_ServiceDesc is the grpc.ServiceDesc for FriendshipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendshipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qianxia.IMChat.FriendshipService",
	HandlerType: (*FriendshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateFriend",
			Handler:    _FriendshipService_UpdateFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _FriendshipService_DeleteFriend_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFriend",
			Handler:       _FriendshipService_ListFriend_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "friendship_service.proto",
}