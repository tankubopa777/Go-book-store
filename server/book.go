package server

import (
	"log"
	"tansan/modules/book/bookHandler"
	bookPb "tansan/modules/book/bookPb"
	"tansan/modules/book/bookRepository"
	"tansan/modules/book/bookUsecase"
	"tansan/pkg/grpccon"
)

func (s *server) bookService() {
	repo := bookRepository.NewBookRepository(s.db)
	usecase := bookUsecase.NewBookUsecase(repo)
	httpHandler := bookHandler.NewBookHttpHandler(s.cfg, usecase)
	grpcHandler := bookHandler.NewBookGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.BookUrl)

		bookPb.RegisterBookGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("gRPC server running at %s", s.cfg.Grpc.BookUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler

	book := s.app.Group("/book_v1")

	// Health Check
	book.GET("", s.healthCheckService)
	book.POST("/book", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.CreateBook, []int{1,0})))
	book.GET("/book/:book_id", httpHandler.FindOneBook)
	book.GET("/book", httpHandler.FindManyBooks)
}