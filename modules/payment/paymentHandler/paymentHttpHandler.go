package paymentHandler

import (
	"tansan/config"
	"tansan/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface {}

	paymentHttpHandler struct {
		cfg *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}