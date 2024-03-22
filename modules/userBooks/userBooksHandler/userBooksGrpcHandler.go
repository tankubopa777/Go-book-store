package userBooksHandler

import "tansan/modules/userBooks/userBooksUsecase"

type (
	userBooksGrpcHandler struct {
		userBooksUsecase userBooksUsecase.UserBooksUsecaseService
	}
)

func NewUserBooksGrpcHandler(userBooksUsecase userBooksUsecase.UserBooksUsecaseService) *userBooksGrpcHandler {
	return &userBooksGrpcHandler{
		userBooksUsecase: userBooksUsecase,
	}
}