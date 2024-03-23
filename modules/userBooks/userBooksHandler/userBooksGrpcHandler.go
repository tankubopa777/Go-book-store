package userbooksHandler

import "tansan/modules/userbooks/userbooksUsecase"

type (
	userbooksGrpcHandler struct {
		userbooksUsecase userbooksUsecase.UserbooksUsecaseService
	}
)

func NewUserbooksGrpcHandler(userbooksUsecase userbooksUsecase.UserbooksUsecaseService) *userbooksGrpcHandler {
	return &userbooksGrpcHandler{
		userbooksUsecase: userbooksUsecase,
	}
}