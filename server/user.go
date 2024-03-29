package server

import (
	"log"
	"tansan/modules/user/userHandler"
	userPb "tansan/modules/user/userPb"
	"tansan/modules/user/userRepository"
	"tansan/modules/user/userUsecase"
	"tansan/pkg/grpccon"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo )
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)
	queueHandler := userHandler.NewUserQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)

		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("gRPC server running at %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()


	_ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)

	user.POST("/user/register", httpHandler.CreateUser)
	user.POST("/user/add-money", httpHandler.AddUserMoney, s.middleware.JwtAuthorization)
	user.GET("/user/:user_id", httpHandler.FindOneUserProfile)
	user.GET("/user/saving-account/my-account", httpHandler.GetUserSavingAccount, s.middleware.JwtAuthorization)
}