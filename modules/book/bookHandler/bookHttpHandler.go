package bookHandler

import (
	"tansan/config"
	"tansan/modules/book/bookUsecase"
)

type (
	BookHttpHandlerService interface {}

	bookHttpHandler struct {
		cfg *config.Config
		bookUsecase bookUsecase.BookUsecaseService
	}
)

func NewBookHttpHandler(cfg *config.Config, bookUsecase bookUsecase.BookUsecaseService) BookHttpHandlerService {
	return &bookHttpHandler{
		cfg: cfg,
		bookUsecase: bookUsecase}
}