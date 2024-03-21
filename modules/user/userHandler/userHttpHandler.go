package userHandler

import (
	"tansan/config"
	"tansan/modules/user/userUsecase"
)

type (
	UserHttpHandlerService interface {}

	userHttpHandler struct {
		cfg *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{userUsecase: userUsecase}
}