package userbooksHandler

import (
	"tansan/config"
	"tansan/modules/userbooks/userbooksUsecase"
)

type (
	UserbooksQueueHandlerService interface {

	}
	userbooksQueueHandler struct {
		cfg *config.Config
		userbooksUsecase userbooksUsecase.UserbooksUsecaseService
	}
)

func NewUserbooksQueueHandler(cfg *config.Config, userbooksUsecase userbooksUsecase.UserbooksUsecaseService) UserbooksQueueHandlerService {
	return &userbooksQueueHandler{
		cfg : cfg,
		userbooksUsecase: userbooksUsecase,
	}
}