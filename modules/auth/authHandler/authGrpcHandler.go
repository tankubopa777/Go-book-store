package authHandler

import "tansan/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase	authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{authUsecase}
}