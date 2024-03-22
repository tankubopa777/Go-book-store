package server

import (
	"tansan/modules/book/bookHandler"
	"tansan/modules/book/bookRepository"
	"tansan/modules/book/bookUsecase"
)

func (s *server) bookService() {
	repo := bookRepository.NewBookRepository(s.db)
	usecase := bookUsecase.NewBookUsecase(repo)
	httpHandler := bookHandler.NewBookHttpHandler(s.cfg, usecase)
	grpcHandler := bookHandler.NewBookGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	book := s.app.Group("/book_v1")

	// Health Check
	book.GET("", s.healthCheckService)
}