package userBooksHandler

import (
	"tansan/config"
	"tansan/modules/userBooks/userBooksUsecase"
)

type (
	UserBooksQueueHandlerService interface {

	}
	userBooksQueueHandler struct {
		cfg *config.Config
		userBooksUsecase userBooksUsecase.UserBooksUsecaseService
	}
)

func NewUserBooksQueueHandler(cfg *config.Config, userBooksUsecase userBooksUsecase.UserBooksUsecaseService) UserBooksQueueHandlerService {
	return &userBooksQueueHandler{
		cfg : cfg,
		userBooksUsecase: userBooksUsecase,
	}
}