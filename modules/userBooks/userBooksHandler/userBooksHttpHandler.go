package userBooksHandler

import (
	"tansan/config"
	"tansan/modules/userBooks/userBooksUsecase"
)

type (
	UserBooksHttpHandlerService interface {

	}
	userBooksHttpHandler struct {
		cfg *config.Config
		userBooksUsecase userBooksUsecase.UserBooksUsecaseService
	}
)

func NewUserBooksHttpHandler(cfg *config.Config, userBooksUsecase userBooksUsecase.UserBooksUsecaseService) UserBooksHttpHandlerService {
	return &userBooksHttpHandler{
		cfg : cfg,
		userBooksUsecase: userBooksUsecase,
	}
}