// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: modules/userbooks/userbooksPb/userbooksPb.proto

package userbooksPb

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

// UserbooksGrpcServiceClient is the client API for UserbooksGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserbooksGrpcServiceClient interface {
	IsAvailabelToSell(ctx context.Context, in *IsAvailableToSellReq, opts ...grpc.CallOption) (*IsAvailableToSellRes, error)
}

type userbooksGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserbooksGrpcServiceClient(cc grpc.ClientConnInterface) UserbooksGrpcServiceClient {
	return &userbooksGrpcServiceClient{cc}
}

func (c *userbooksGrpcServiceClient) IsAvailabelToSell(ctx context.Context, in *IsAvailableToSellReq, opts ...grpc.CallOption) (*IsAvailableToSellRes, error) {
	out := new(IsAvailableToSellRes)
	err := c.cc.Invoke(ctx, "/UserbooksGrpcService/IsAvailabelToSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserbooksGrpcServiceServer is the server API for UserbooksGrpcService service.
// All implementations must embed UnimplementedUserbooksGrpcServiceServer
// for forward compatibility
type UserbooksGrpcServiceServer interface {
	IsAvailabelToSell(context.Context, *IsAvailableToSellReq) (*IsAvailableToSellRes, error)
	mustEmbedUnimplementedUserbooksGrpcServiceServer()
}

// UnimplementedUserbooksGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserbooksGrpcServiceServer struct {
}

func (UnimplementedUserbooksGrpcServiceServer) IsAvailabelToSell(context.Context, *IsAvailableToSellReq) (*IsAvailableToSellRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailabelToSell not implemented")
}
func (UnimplementedUserbooksGrpcServiceServer) mustEmbedUnimplementedUserbooksGrpcServiceServer() {}

// UnsafeUserbooksGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserbooksGrpcServiceServer will
// result in compilation errors.
type UnsafeUserbooksGrpcServiceServer interface {
	mustEmbedUnimplementedUserbooksGrpcServiceServer()
}

func RegisterUserbooksGrpcServiceServer(s grpc.ServiceRegistrar, srv UserbooksGrpcServiceServer) {
	s.RegisterService(&UserbooksGrpcService_ServiceDesc, srv)
}

func _UserbooksGrpcService_IsAvailabelToSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvailableToSellReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserbooksGrpcServiceServer).IsAvailabelToSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserbooksGrpcService/IsAvailabelToSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserbooksGrpcServiceServer).IsAvailabelToSell(ctx, req.(*IsAvailableToSellReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserbooksGrpcService_ServiceDesc is the grpc.ServiceDesc for UserbooksGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserbooksGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserbooksGrpcService",
	HandlerType: (*UserbooksGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailabelToSell",
			Handler:    _UserbooksGrpcService_IsAvailabelToSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/userbooks/userbooksPb/userbooksPb.proto",
}