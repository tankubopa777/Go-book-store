package server

import (
	"tansan/modules/userbooks/userbooksHandler"
	"tansan/modules/userbooks/userbooksRepository"
	"tansan/modules/userbooks/userbooksUsecase"
)

func (s *server) userbooksService() {
	repo := userbooksRepository.NewUserbooksRepository(s.db)
	usecase := userbooksUsecase.NewUserbooksUsecase(repo)
	httpHandler := userbooksHandler.NewUserbooksHttpHandler(s.cfg, usecase)
	grpcHandler := userbooksHandler.NewUserbooksGrpcHandler(usecase)
	queueHandler := userbooksHandler.NewUserbooksQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	userbooks := s.app.Group("/userbooks_v1")

	// Health Check
	userbooks.GET("", s.healthCheckService)
}