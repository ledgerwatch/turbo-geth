// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sentry

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

// SentryClient is the client API for Sentry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentryClient interface {
	PenalizePeer(ctx context.Context, in *PenalizePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PeerMinBlock(ctx context.Context, in *PeerMinBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessageByMinBlock(ctx context.Context, in *SendMessageByMinBlockRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageById(ctx context.Context, in *SendMessageByIdRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageToRandomPeers(ctx context.Context, in *SendMessageToRandomPeersRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageToAll(ctx context.Context, in *OutboundMessageData, opts ...grpc.CallOption) (*SentPeers, error)
	SetStatus(ctx context.Context, in *StatusData, opts ...grpc.CallOption) (*SetStatusReply, error)
	ReceiveMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveMessagesClient, error)
	ReceiveUploadMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveUploadMessagesClient, error)
	ReceiveTxMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveTxMessagesClient, error)
}

type sentryClient struct {
	cc grpc.ClientConnInterface
}

func NewSentryClient(cc grpc.ClientConnInterface) SentryClient {
	return &sentryClient{cc}
}

func (c *sentryClient) PenalizePeer(ctx context.Context, in *PenalizePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/PenalizePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) PeerMinBlock(ctx context.Context, in *PeerMinBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/PeerMinBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageByMinBlock(ctx context.Context, in *SendMessageByMinBlockRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageByMinBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageById(ctx context.Context, in *SendMessageByIdRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageToRandomPeers(ctx context.Context, in *SendMessageToRandomPeersRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageToRandomPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageToAll(ctx context.Context, in *OutboundMessageData, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageToAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SetStatus(ctx context.Context, in *StatusData, opts ...grpc.CallOption) (*SetStatusReply, error) {
	out := new(SetStatusReply)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) ReceiveMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sentry_ServiceDesc.Streams[0], "/sentry.Sentry/ReceiveMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryReceiveMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sentry_ReceiveMessagesClient interface {
	Recv() (*InboundMessage, error)
	grpc.ClientStream
}

type sentryReceiveMessagesClient struct {
	grpc.ClientStream
}

func (x *sentryReceiveMessagesClient) Recv() (*InboundMessage, error) {
	m := new(InboundMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sentryClient) ReceiveUploadMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveUploadMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sentry_ServiceDesc.Streams[1], "/sentry.Sentry/ReceiveUploadMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryReceiveUploadMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sentry_ReceiveUploadMessagesClient interface {
	Recv() (*InboundMessage, error)
	grpc.ClientStream
}

type sentryReceiveUploadMessagesClient struct {
	grpc.ClientStream
}

func (x *sentryReceiveUploadMessagesClient) Recv() (*InboundMessage, error) {
	m := new(InboundMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sentryClient) ReceiveTxMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Sentry_ReceiveTxMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sentry_ServiceDesc.Streams[2], "/sentry.Sentry/ReceiveTxMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryReceiveTxMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sentry_ReceiveTxMessagesClient interface {
	Recv() (*InboundMessage, error)
	grpc.ClientStream
}

type sentryReceiveTxMessagesClient struct {
	grpc.ClientStream
}

func (x *sentryReceiveTxMessagesClient) Recv() (*InboundMessage, error) {
	m := new(InboundMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SentryServer is the server API for Sentry service.
// All implementations must embed UnimplementedSentryServer
// for forward compatibility
type SentryServer interface {
	PenalizePeer(context.Context, *PenalizePeerRequest) (*emptypb.Empty, error)
	PeerMinBlock(context.Context, *PeerMinBlockRequest) (*emptypb.Empty, error)
	SendMessageByMinBlock(context.Context, *SendMessageByMinBlockRequest) (*SentPeers, error)
	SendMessageById(context.Context, *SendMessageByIdRequest) (*SentPeers, error)
	SendMessageToRandomPeers(context.Context, *SendMessageToRandomPeersRequest) (*SentPeers, error)
	SendMessageToAll(context.Context, *OutboundMessageData) (*SentPeers, error)
	SetStatus(context.Context, *StatusData) (*SetStatusReply, error)
	ReceiveMessages(*emptypb.Empty, Sentry_ReceiveMessagesServer) error
	ReceiveUploadMessages(*emptypb.Empty, Sentry_ReceiveUploadMessagesServer) error
	ReceiveTxMessages(*emptypb.Empty, Sentry_ReceiveTxMessagesServer) error
	mustEmbedUnimplementedSentryServer()
}

// UnimplementedSentryServer must be embedded to have forward compatible implementations.
type UnimplementedSentryServer struct {
}

func (UnimplementedSentryServer) PenalizePeer(context.Context, *PenalizePeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PenalizePeer not implemented")
}
func (UnimplementedSentryServer) PeerMinBlock(context.Context, *PeerMinBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerMinBlock not implemented")
}
func (UnimplementedSentryServer) SendMessageByMinBlock(context.Context, *SendMessageByMinBlockRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageByMinBlock not implemented")
}
func (UnimplementedSentryServer) SendMessageById(context.Context, *SendMessageByIdRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageById not implemented")
}
func (UnimplementedSentryServer) SendMessageToRandomPeers(context.Context, *SendMessageToRandomPeersRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToRandomPeers not implemented")
}
func (UnimplementedSentryServer) SendMessageToAll(context.Context, *OutboundMessageData) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToAll not implemented")
}
func (UnimplementedSentryServer) SetStatus(context.Context, *StatusData) (*SetStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedSentryServer) ReceiveMessages(*emptypb.Empty, Sentry_ReceiveMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessages not implemented")
}
func (UnimplementedSentryServer) ReceiveUploadMessages(*emptypb.Empty, Sentry_ReceiveUploadMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveUploadMessages not implemented")
}
func (UnimplementedSentryServer) ReceiveTxMessages(*emptypb.Empty, Sentry_ReceiveTxMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveTxMessages not implemented")
}
func (UnimplementedSentryServer) mustEmbedUnimplementedSentryServer() {}

// UnsafeSentryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentryServer will
// result in compilation errors.
type UnsafeSentryServer interface {
	mustEmbedUnimplementedSentryServer()
}

func RegisterSentryServer(s grpc.ServiceRegistrar, srv SentryServer) {
	s.RegisterService(&Sentry_ServiceDesc, srv)
}

func _Sentry_PenalizePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PenalizePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).PenalizePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/PenalizePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).PenalizePeer(ctx, req.(*PenalizePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_PeerMinBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMinBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).PeerMinBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/PeerMinBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).PeerMinBlock(ctx, req.(*PeerMinBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageByMinBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageByMinBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageByMinBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageByMinBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageByMinBlock(ctx, req.(*SendMessageByMinBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageById(ctx, req.(*SendMessageByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageToRandomPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToRandomPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageToRandomPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageToRandomPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageToRandomPeers(ctx, req.(*SendMessageToRandomPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageToAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutboundMessageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageToAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageToAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageToAll(ctx, req.(*OutboundMessageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SetStatus(ctx, req.(*StatusData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_ReceiveMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServer).ReceiveMessages(m, &sentryReceiveMessagesServer{stream})
}

type Sentry_ReceiveMessagesServer interface {
	Send(*InboundMessage) error
	grpc.ServerStream
}

type sentryReceiveMessagesServer struct {
	grpc.ServerStream
}

func (x *sentryReceiveMessagesServer) Send(m *InboundMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Sentry_ReceiveUploadMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServer).ReceiveUploadMessages(m, &sentryReceiveUploadMessagesServer{stream})
}

type Sentry_ReceiveUploadMessagesServer interface {
	Send(*InboundMessage) error
	grpc.ServerStream
}

type sentryReceiveUploadMessagesServer struct {
	grpc.ServerStream
}

func (x *sentryReceiveUploadMessagesServer) Send(m *InboundMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Sentry_ReceiveTxMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServer).ReceiveTxMessages(m, &sentryReceiveTxMessagesServer{stream})
}

type Sentry_ReceiveTxMessagesServer interface {
	Send(*InboundMessage) error
	grpc.ServerStream
}

type sentryReceiveTxMessagesServer struct {
	grpc.ServerStream
}

func (x *sentryReceiveTxMessagesServer) Send(m *InboundMessage) error {
	return x.ServerStream.SendMsg(m)
}

// Sentry_ServiceDesc is the grpc.ServiceDesc for Sentry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sentry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentry.Sentry",
	HandlerType: (*SentryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PenalizePeer",
			Handler:    _Sentry_PenalizePeer_Handler,
		},
		{
			MethodName: "PeerMinBlock",
			Handler:    _Sentry_PeerMinBlock_Handler,
		},
		{
			MethodName: "SendMessageByMinBlock",
			Handler:    _Sentry_SendMessageByMinBlock_Handler,
		},
		{
			MethodName: "SendMessageById",
			Handler:    _Sentry_SendMessageById_Handler,
		},
		{
			MethodName: "SendMessageToRandomPeers",
			Handler:    _Sentry_SendMessageToRandomPeers_Handler,
		},
		{
			MethodName: "SendMessageToAll",
			Handler:    _Sentry_SendMessageToAll_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _Sentry_SetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessages",
			Handler:       _Sentry_ReceiveMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiveUploadMessages",
			Handler:       _Sentry_ReceiveUploadMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiveTxMessages",
			Handler:       _Sentry_ReceiveTxMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "p2psentry/sentry.proto",
}
