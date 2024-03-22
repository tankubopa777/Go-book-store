package paymentHandler

import (
	"tansan/config"
	"tansan/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface {}

	paymentQueueHandler struct {
		cfg *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}