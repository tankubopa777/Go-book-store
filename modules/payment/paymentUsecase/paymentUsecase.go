package paymentUsecase

import (
	"tansan/modules/payment/paymentRepository"
)

type (
	PaymentUsecaseService interface {}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentRepository(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}