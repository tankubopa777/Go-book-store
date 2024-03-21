package middlewareHandler

import (
	"tansan/config"
	"tansan/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {}

	middlewareHandler struct {
		cfg *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg : cfg,
		middlewareUsecase : middlewareUsecase}
}