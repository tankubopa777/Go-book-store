package server

import (
	"tansan/modules/userBooks/userBooksHandler"
	"tansan/modules/userBooks/userBooksRepository"
	"tansan/modules/userBooks/userBooksUsecase"
)

func (s *server) userBooksService() {
	repo := userBooksRepository.NewUserBooksRepository(s.db)
	usecase := userBooksUsecase.NewUserBooksUsecase(repo)
	httpHandler := userBooksHandler.NewUserBooksHttpHandler(s.cfg, usecase)
	grpcHandler := userBooksHandler.NewUserBooksGrpcHandler(usecase)
	queueHandler := userBooksHandler.NewUserBooksQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	userBooks := s.app.Group("/userBooks_v1")

	// Health Check
	userBooks.GET("", s.healthCheckService)
}