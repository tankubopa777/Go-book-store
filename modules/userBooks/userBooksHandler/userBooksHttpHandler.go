package userbooksHandler

import (
	"tansan/config"
	"tansan/modules/userbooks/userbooksUsecase"
)

type (
	UserbooksHttpHandlerService interface {}
	userbooksHttpHandler struct {
		cfg *config.Config
		userbooksUsecase userbooksUsecase.UserbooksUsecaseService
	}
)

func NewUserbooksHttpHandler(cfg *config.Config, userbooksUsecase userbooksUsecase.UserbooksUsecaseService) UserbooksHttpHandlerService {
	return &userbooksHttpHandler{
		cfg : cfg,
		userbooksUsecase: userbooksUsecase,
	}
}
