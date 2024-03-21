package userHandler

import (
	"tansan/modules/user/userUsecase"
)

type (
	userGrpcHandlerService struct {
		userUsecase	userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase	userUsecase.UserUsecaseService) userGrpcHandlerService {
	return userGrpcHandlerService{userUsecase}
}