package authUsecase

import (
	"context"
	"errors"
	"log"
	"strings"
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
		RefreshToken(pctx context.Context, cfg *config.Config, req *auth.RefreshTokenReq) (*auth.ProfileIntercepter, error)
		Logout(pctx context.Context, credentialId string) (int64, error)
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

	credential, err := u.authRepository.FindOneUserCredential(pctx, credentialId.Hex())
	if err != nil {
		return nil, err
	
	}


	loc, _ := time.LoadLocation("Asia/Bangkok")

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
			UserId: profile.Id,
			RoleCode: int(profile.RoleCode),
			AccessToken: credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt: credential.CreatedAt.In(loc),
			UpdatedAt: credential.UpdatedAt.In(loc),
		},
	}, nil
}

func (u *authUsecase) RefreshToken(pctx context.Context, cfg *config.Config, req *auth.RefreshTokenReq) (*auth.ProfileIntercepter, error) {
	claims, err := jwtauth.ParseToken(cfg.Jwt.RefreshSecretKey, req.RefreshToken)
	if err != nil {
		log.Printf("Error: RefreshToken: %v", err.Error())
		return nil, errors.New(err.Error())
	}

	profile, err := u.authRepository.FindOneUserProfileToRefresh(pctx, cfg.Grpc.UserUrl, &userPb.FindOneUserProfileToRefreshReq{
		UserId: strings.TrimPrefix(claims.UserId, "user:"),
	})
	if err != nil {
		return nil, err
	}

	accessToken := jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		UserId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtauth.ReloadToken(cfg.Jwt.RefreshSecretKey, claims.ExpiresAt.Unix(), &jwtauth.Claims{
		UserId: profile.Id,
		RoleCode: int(profile.RoleCode),
	})

	if err := u.authRepository.UpdateOneUserCredential(pctx, req.CredentialId, &auth.UpdateRefreshTokenReq{
		UserId: profile.Id,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		UpdatedAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	credential, err := u.authRepository.FindOneUserCredential(pctx, req.CredentialId)
	if err != nil {
		return nil, err
	}

	return &auth.ProfileIntercepter{
		UserProfile: &user.UserProfile{
			Id:        "user:" + profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringTimeToTime(profile.CreatedAt).String(),
			UpdatedAt: utils.ConvertStringTimeToTime(profile.UpdatedAt).String(),
		},
		Credential: &auth.CredentialRes{
			Id: credential.Id.Hex(),
			UserId: credential.UserId,
			RoleCode: credential.RoleCode,
			AccessToken: credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt: credential.CreatedAt,
			UpdatedAt: credential.UpdatedAt,
		},
	}, nil
}

func (u *authUsecase) Logout(pctx context.Context, credentialId string) (int64, error){
	return u.authRepository.DeleteOneUserCredential(pctx, credentialId)
}