package authUsecase

import (
	"tansan/modules/auth/authRepository"
)


type(
	AuthUsecaseService interface {}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}