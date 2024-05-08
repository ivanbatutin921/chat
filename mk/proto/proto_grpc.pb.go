// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LiveChatClient is the client API for LiveChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveChatClient interface {
	ChatStream(ctx context.Context, opts ...grpc.CallOption) (LiveChat_ChatStreamClient, error)
}

type liveChatClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveChatClient(cc grpc.ClientConnInterface) LiveChatClient {
	return &liveChatClient{cc}
}

func (c *liveChatClient) ChatStream(ctx context.Context, opts ...grpc.CallOption) (LiveChat_ChatStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &LiveChat_ServiceDesc.Streams[0], "/proto.LiveChat/ChatStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &liveChatChatStreamClient{stream}
	return x, nil
}

type LiveChat_ChatStreamClient interface {
	Send(*LiveChatData) error
	Recv() (*LiveChatData, error)
	grpc.ClientStream
}

type liveChatChatStreamClient struct {
	grpc.ClientStream
}

func (x *liveChatChatStreamClient) Send(m *LiveChatData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *liveChatChatStreamClient) Recv() (*LiveChatData, error) {
	m := new(LiveChatData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LiveChatServer is the server API for LiveChat service.
// All implementations must embed UnimplementedLiveChatServer
// for forward compatibility
type LiveChatServer interface {
	ChatStream(LiveChat_ChatStreamServer) error
	mustEmbedUnimplementedLiveChatServer()
}

// UnimplementedLiveChatServer must be embedded to have forward compatible implementations.
type UnimplementedLiveChatServer struct {
}

func (UnimplementedLiveChatServer) ChatStream(LiveChat_ChatStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatStream not implemented")
}
func (UnimplementedLiveChatServer) mustEmbedUnimplementedLiveChatServer() {}

// UnsafeLiveChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveChatServer will
// result in compilation errors.
type UnsafeLiveChatServer interface {
	mustEmbedUnimplementedLiveChatServer()
}

func RegisterLiveChatServer(s grpc.ServiceRegistrar, srv LiveChatServer) {
	s.RegisterService(&LiveChat_ServiceDesc, srv)
}

func _LiveChat_ChatStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LiveChatServer).ChatStream(&liveChatChatStreamServer{stream})
}

type LiveChat_ChatStreamServer interface {
	Send(*LiveChatData) error
	Recv() (*LiveChatData, error)
	grpc.ServerStream
}

type liveChatChatStreamServer struct {
	grpc.ServerStream
}

func (x *liveChatChatStreamServer) Send(m *LiveChatData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *liveChatChatStreamServer) Recv() (*LiveChatData, error) {
	m := new(LiveChatData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LiveChat_ServiceDesc is the grpc.ServiceDesc for LiveChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LiveChat",
	HandlerType: (*LiveChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatStream",
			Handler:       _LiveChat_ChatStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto.proto",
}
