package bookHandler

import (
	"tansan/modules/book/bookUsecase"
)

type (
	bookGrpcHandler struct {
		bookUsecase bookUsecase.BookUsecaseService
	}
)

func NewBookGrpcHandler(bookUsecase bookUsecase.BookUsecaseService) *bookGrpcHandler {
	return &bookGrpcHandler{
		bookUsecase: bookUsecase}
}