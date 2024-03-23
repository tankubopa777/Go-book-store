package userbooksHandler

import (
	"context"
	userbooksPb "tansan/modules/userbooks/userbooksPb"
	"tansan/modules/userbooks/userbooksUsecase"
)

type (
	userbooksGrpcHandler struct {
		userbooksUsecase userbooksUsecase.UserbooksUsecaseService
		userbooksPb.UnimplementedUserbooksGrpcServiceServer
	}
)

func NewUserbooksGrpcHandler(userbooksUsecase userbooksUsecase.UserbooksUsecaseService) *userbooksGrpcHandler {
	return &userbooksGrpcHandler{
		userbooksUsecase: userbooksUsecase,
	}
}

func (g *userbooksGrpcHandler) IsAvailableToSell(ctx context.Context, req *userbooksPb.IsAvailableToSellReq) (*userbooksPb.IsAvailableToSellRes, error) {
	return nil, nil
}