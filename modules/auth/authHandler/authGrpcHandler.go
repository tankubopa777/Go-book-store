package authHandler

import (
	"context"
	authPb "tansan/modules/auth/authPb"
	"tansan/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase	authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.CredentialSearchReq) (*authPb.CredentialSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RoleCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}