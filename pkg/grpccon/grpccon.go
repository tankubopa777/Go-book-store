package grpccon

import (
	"errors"
	"log"
	"net"
	"tansan/config"
	authPb "tansan/modules/auth/authPb"
	bookPb "tansan/modules/book/bookPb"
	userPb "tansan/modules/user/userPb"
	userbooksPb "tansan/modules/userbooks/userbooksPb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		User() userPb.UserGrpcServiceClient
		Book() bookPb.BookGrpcServiceClient
		Userbooks() userbooksPb.UserbooksGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
}
)

// Module grpccon.go is a module that contains the grpc client and server connection functions.
func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient{
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) User() userPb.UserGrpcServiceClient{
	return userPb.NewUserGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Book() bookPb.BookGrpcServiceClient{
	return bookPb.NewBookGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Userbooks() userbooksPb.UserbooksGrpcServiceClient{
	return userbooksPb.NewUserbooksGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: failed to listen: %v", err)
	}

	return grpcServer, lis
}