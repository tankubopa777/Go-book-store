// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: modules/user/userPb/userPb.proto

package authPb

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

// UserGrpcServiceClient is the client API for UserGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcServiceClient interface {
	CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*UserProfile, error)
	FindOneUserProfileToRefresh(ctx context.Context, in *FindOneUserProfileToRefreshReq, opts ...grpc.CallOption) (*UserProfile, error)
	GetUserSavingAccount(ctx context.Context, in *GetUserSavingAccountReq, opts ...grpc.CallOption) (*GetUserSavingAccountRes, error)
}

type userGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcServiceClient(cc grpc.ClientConnInterface) UserGrpcServiceClient {
	return &userGrpcServiceClient{cc}
}

func (c *userGrpcServiceClient) CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/UserGrpcService/CredentialSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) FindOneUserProfileToRefresh(ctx context.Context, in *FindOneUserProfileToRefreshReq, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/UserGrpcService/FindOneUserProfileToRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) GetUserSavingAccount(ctx context.Context, in *GetUserSavingAccountReq, opts ...grpc.CallOption) (*GetUserSavingAccountRes, error) {
	out := new(GetUserSavingAccountRes)
	err := c.cc.Invoke(ctx, "/UserGrpcService/GetUserSavingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServiceServer is the server API for UserGrpcService service.
// All implementations must embed UnimplementedUserGrpcServiceServer
// for forward compatibility
type UserGrpcServiceServer interface {
	CredentialSearch(context.Context, *CredentialSearchReq) (*UserProfile, error)
	FindOneUserProfileToRefresh(context.Context, *FindOneUserProfileToRefreshReq) (*UserProfile, error)
	GetUserSavingAccount(context.Context, *GetUserSavingAccountReq) (*GetUserSavingAccountRes, error)
	mustEmbedUnimplementedUserGrpcServiceServer()
}

// UnimplementedUserGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserGrpcServiceServer struct {
}

func (UnimplementedUserGrpcServiceServer) CredentialSearch(context.Context, *CredentialSearchReq) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialSearch not implemented")
}
func (UnimplementedUserGrpcServiceServer) FindOneUserProfileToRefresh(context.Context, *FindOneUserProfileToRefreshReq) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUserProfileToRefresh not implemented")
}
func (UnimplementedUserGrpcServiceServer) GetUserSavingAccount(context.Context, *GetUserSavingAccountReq) (*GetUserSavingAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSavingAccount not implemented")
}
func (UnimplementedUserGrpcServiceServer) mustEmbedUnimplementedUserGrpcServiceServer() {}

// UnsafeUserGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServiceServer will
// result in compilation errors.
type UnsafeUserGrpcServiceServer interface {
	mustEmbedUnimplementedUserGrpcServiceServer()
}

func RegisterUserGrpcServiceServer(s grpc.ServiceRegistrar, srv UserGrpcServiceServer) {
	s.RegisterService(&UserGrpcService_ServiceDesc, srv)
}

func _UserGrpcService_CredentialSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).CredentialSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/CredentialSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).CredentialSearch(ctx, req.(*CredentialSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_FindOneUserProfileToRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserProfileToRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).FindOneUserProfileToRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/FindOneUserProfileToRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).FindOneUserProfileToRefresh(ctx, req.(*FindOneUserProfileToRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_GetUserSavingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSavingAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).GetUserSavingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/GetUserSavingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).GetUserSavingAccount(ctx, req.(*GetUserSavingAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpcService_ServiceDesc is the grpc.ServiceDesc for UserGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserGrpcService",
	HandlerType: (*UserGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialSearch",
			Handler:    _UserGrpcService_CredentialSearch_Handler,
		},
		{
			MethodName: "FindOneUserProfileToRefresh",
			Handler:    _UserGrpcService_FindOneUserProfileToRefresh_Handler,
		},
		{
			MethodName: "GetUserSavingAccount",
			Handler:    _UserGrpcService_GetUserSavingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/userPb/userPb.proto",
}
