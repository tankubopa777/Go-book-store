// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: modules/book/bookPb/bookPb.proto

package bookPb

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

// BookGrpcServiceClient is the client API for BookGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookGrpcServiceClient interface {
	FindBooksInIds(ctx context.Context, in *FindBooksInIdsReq, opts ...grpc.CallOption) (*FindBooksInIdsRes, error)
}

type bookGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookGrpcServiceClient(cc grpc.ClientConnInterface) BookGrpcServiceClient {
	return &bookGrpcServiceClient{cc}
}

func (c *bookGrpcServiceClient) FindBooksInIds(ctx context.Context, in *FindBooksInIdsReq, opts ...grpc.CallOption) (*FindBooksInIdsRes, error) {
	out := new(FindBooksInIdsRes)
	err := c.cc.Invoke(ctx, "/bookGrpcService/FindBooksInIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookGrpcServiceServer is the server API for BookGrpcService service.
// All implementations must embed UnimplementedBookGrpcServiceServer
// for forward compatibility
type BookGrpcServiceServer interface {
	FindBooksInIds(context.Context, *FindBooksInIdsReq) (*FindBooksInIdsRes, error)
	mustEmbedUnimplementedBookGrpcServiceServer()
}

// UnimplementedBookGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookGrpcServiceServer struct {
}

func (UnimplementedBookGrpcServiceServer) FindBooksInIds(context.Context, *FindBooksInIdsReq) (*FindBooksInIdsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBooksInIds not implemented")
}
func (UnimplementedBookGrpcServiceServer) mustEmbedUnimplementedBookGrpcServiceServer() {}

// UnsafeBookGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookGrpcServiceServer will
// result in compilation errors.
type UnsafeBookGrpcServiceServer interface {
	mustEmbedUnimplementedBookGrpcServiceServer()
}

func RegisterBookGrpcServiceServer(s grpc.ServiceRegistrar, srv BookGrpcServiceServer) {
	s.RegisterService(&BookGrpcService_ServiceDesc, srv)
}

func _BookGrpcService_FindBooksInIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBooksInIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookGrpcServiceServer).FindBooksInIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookGrpcService/FindBooksInIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookGrpcServiceServer).FindBooksInIds(ctx, req.(*FindBooksInIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BookGrpcService_ServiceDesc is the grpc.ServiceDesc for BookGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookGrpcService",
	HandlerType: (*BookGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBooksInIds",
			Handler:    _BookGrpcService_FindBooksInIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/book/bookPb/bookPb.proto",
}
