package userHandler

import (
	"context"
	userPb "tansan/modules/user/userPb"
	"tansan/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		userUsecase	userUsecase.UserUsecaseService
		userPb.UnimplementedUserGrpcServiceServer
	}
)

func NewUserGrpcHandler(userUsecase	userUsecase.UserUsecaseService) *userGrpcHandler {
	return &userGrpcHandler{
		userUsecase: userUsecase,
	}
}

func (g *userGrpcHandler) CredentialSearch(ctx context.Context, req *userPb.CredentialSearchReq) (*userPb.UserProfile, error) {
	return g.userUsecase.FindOneUserCredential(ctx, req.Password, req.Email)
}

func (g *userGrpcHandler) FindOneUserProfileToRefresh(ctx context.Context, req *userPb.FindOneUserProfileToRefreshReq) (*userPb.UserProfile, error) {
	return g.userUsecase.FindOneUserProfileToRefresh(ctx, req.UserId)
}

func (g *userGrpcHandler) GetUserSavingAccount(ctx context.Context, req *userPb.GetUserSavingAccountReq) (*userPb.GetUserSavingAccountRes, error) {
	return nil, nil
}