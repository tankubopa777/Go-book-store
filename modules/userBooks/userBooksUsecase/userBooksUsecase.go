package userbooksUsecase

import (
	"tansan/modules/userbooks/userbooksRepository"
)

type (
	UserbooksUsecaseService interface {}

	userbooksUsecase struct {
		userbooksRepository userbooksRepository.UserbooksRepositoryService
	}
)

func NewUserbooksUsecase(userbooksRepository userbooksRepository.UserbooksRepositoryService) UserbooksUsecaseService {
	return &userbooksUsecase{
		userbooksRepository: userbooksRepository,
	}
}

