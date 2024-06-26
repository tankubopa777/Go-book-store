package server

import (
	"log"
	"tansan/modules/auth/authHandler"
	authPb "tansan/modules/auth/authPb"
	"tansan/modules/auth/authRepository"
	"tansan/modules/auth/authUsecase"
	"tansan/pkg/grpccon"
)


func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("gRPC server running at %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/test/:user_id", s.healthCheckService, s.middleware.JwtAuthorization, s.middleware.UserIdParamValidation)
	auth.POST("/auth/login", httpHandler.Login)
	auth.POST("/auth/refresh-token", httpHandler.RefreshToken)
	auth.POST("/auth/logout", httpHandler.Logout)
}