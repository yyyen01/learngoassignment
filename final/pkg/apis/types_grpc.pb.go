// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: types.proto

package apis

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Register(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	GetOnlineUsers(ctx context.Context, in *PageNumber, opts ...grpc.CallOption) (*OnlineUsers, error)
	Chat(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*ChatRecords, error)
	ChatHistory(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*ChatRecords, error)
	Subscribe(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*Status, error)
	EndChat(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*Status, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Register(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/apis.ChatService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/apis.ChatService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetOnlineUsers(ctx context.Context, in *PageNumber, opts ...grpc.CallOption) (*OnlineUsers, error) {
	out := new(OnlineUsers)
	err := c.cc.Invoke(ctx, "/apis.ChatService/GetOnlineUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Chat(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*ChatRecords, error) {
	out := new(ChatRecords)
	err := c.cc.Invoke(ctx, "/apis.ChatService/Chat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ChatHistory(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*ChatRecords, error) {
	out := new(ChatRecords)
	err := c.cc.Invoke(ctx, "/apis.ChatService/ChatHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subscribe(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/apis.ChatService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EndChat(ctx context.Context, in *AccountNumber, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/apis.ChatService/EndChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations should embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Register(context.Context, *UserInfo) (*UserInfo, error)
	Login(context.Context, *UserInfo) (*UserInfo, error)
	GetOnlineUsers(context.Context, *PageNumber) (*OnlineUsers, error)
	Chat(context.Context, *AccountNumber) (*ChatRecords, error)
	ChatHistory(context.Context, *AccountNumber) (*ChatRecords, error)
	Subscribe(context.Context, *AccountNumber) (*Status, error)
	EndChat(context.Context, *AccountNumber) (*Status, error)
}

// UnimplementedChatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Register(context.Context, *UserInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChatServiceServer) Login(context.Context, *UserInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatServiceServer) GetOnlineUsers(context.Context, *PageNumber) (*OnlineUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineUsers not implemented")
}
func (UnimplementedChatServiceServer) Chat(context.Context, *AccountNumber) (*ChatRecords, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServiceServer) ChatHistory(context.Context, *AccountNumber) (*ChatRecords, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatHistory not implemented")
}
func (UnimplementedChatServiceServer) Subscribe(context.Context, *AccountNumber) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedChatServiceServer) EndChat(context.Context, *AccountNumber) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndChat not implemented")
}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Register(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetOnlineUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetOnlineUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/GetOnlineUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetOnlineUsers(ctx, req.(*PageNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/Chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Chat(ctx, req.(*AccountNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ChatHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ChatHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/ChatHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ChatHistory(ctx, req.(*AccountNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Subscribe(ctx, req.(*AccountNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EndChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EndChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ChatService/EndChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EndChat(ctx, req.(*AccountNumber))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ChatService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ChatService_Login_Handler,
		},
		{
			MethodName: "GetOnlineUsers",
			Handler:    _ChatService_GetOnlineUsers_Handler,
		},
		{
			MethodName: "Chat",
			Handler:    _ChatService_Chat_Handler,
		},
		{
			MethodName: "ChatHistory",
			Handler:    _ChatService_ChatHistory_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _ChatService_Subscribe_Handler,
		},
		{
			MethodName: "EndChat",
			Handler:    _ChatService_EndChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}