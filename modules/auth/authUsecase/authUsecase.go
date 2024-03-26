package authUsecase

import (
	"context"
	"tansan/config"
	"tansan/modules/auth"
	"tansan/modules/auth/authRepository"
	"tansan/modules/user"
	userPb "tansan/modules/user/userPb"
	"tansan/pkg/jwtauth"
	"tansan/pkg/utils"
	"time"
)


type(
	AuthUsecaseService interface {
		Login(pctx context.Context, cfg *config.Config, req *auth.UserLoginReq) (*auth.ProfileIntercepter, error)
	}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}

func (u *authUsecase) Login(pctx context.Context, cfg *config.Config, req *auth.UserLoginReq) (*auth.ProfileIntercepter, error){
	profile, err := u.authRepository.CredentialSearch(pctx, cfg.Grpc.UserUrl, &userPb.CredentialSearchReq{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	profile.Id = "user:" + profile.Id

	accessToken := jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		UserId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtauth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtauth.Claims{
		UserId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	credentialId, err := u.authRepository.InsertOneUserCredential(pctx, &auth.Credential{
		UserId: profile.Id,
		RoleCode: int(profile.RoleCode),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	})
	loc, _ := time.LoadLocation("Asia/Jakarta")

	return &auth.ProfileIntercepter{
		UserProfile: &user.UserProfile{
			Id:        profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringTimeToTime(profile.CreatedAt).In(loc).String(),
			UpdatedAt: utils.ConvertStringTimeToTime(profile.UpdatedAt).In(loc).String(),
		},
		Credential: &auth.CredentialRes{
			Id: credentialId.Hex(),
		},
	}, nil
}

