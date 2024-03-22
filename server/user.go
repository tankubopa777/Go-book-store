package server

import (
	"tansan/modules/user/userHandler"
	"tansan/modules/user/userRepository"
	"tansan/modules/user/userUsecase"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo )
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)
	queueHandler := userHandler.NewUserQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)
}