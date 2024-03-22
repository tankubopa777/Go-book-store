package userBooksUsecase

import (
	"tansan/modules/userBooks/userBooksRepository"
)

type (
	UserBooksUsecaseService interface {

	}
	userBooksUsecase struct {
		userBooksRepository userBooksRepository.UserBooksRepositoryService
	}
)

func NewUserBooksUsecase(userBooksRepository userBooksRepository.UserBooksRepositoryService) UserBooksUsecaseService {
	return &userBooksUsecase{
		userBooksRepository: userBooksRepository,
	}
}