package server

import (
	"log"
	"tansan/modules/userbooks/userbooksHandler"
	userbooksPb "tansan/modules/userbooks/userbooksPb"
	"tansan/modules/userbooks/userbooksRepository"
	"tansan/modules/userbooks/userbooksUsecase"
	"tansan/pkg/grpccon"
)

func (s *server) userbooksService() {
	repo := userbooksRepository.NewUserbooksRepository(s.db)
	usecase := userbooksUsecase.NewUserbooksUsecase(repo)
	httpHandler := userbooksHandler.NewUserbooksHttpHandler(s.cfg, usecase)
	grpcHandler := userbooksHandler.NewUserbooksGrpcHandler(usecase)
	queueHandler := userbooksHandler.NewUserbooksQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserbooksUrl)

		userbooksPb.RegisterUserbooksGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("gRPC server running at %s", s.cfg.Grpc.UserbooksUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	userbooks := s.app.Group("/userbooks_v1")

	// Health Check
	userbooks.GET("", s.healthCheckService)
}