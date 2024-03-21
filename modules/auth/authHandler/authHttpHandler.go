package authHandler

import (
	"tansan/config"
	"tansan/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface {}

	authHttpHandler struct {
		cfg *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}