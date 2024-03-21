package userHandler

import (
	"tansan/config"
	"tansan/modules/user/userUsecase"
)

type (
	UserQueueHandlerService interface {}

	userQueueHandler struct {
		cfg *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserQueueHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserQueueHandlerService {
	return &userQueueHandler{userUsecase: userUsecase}
}