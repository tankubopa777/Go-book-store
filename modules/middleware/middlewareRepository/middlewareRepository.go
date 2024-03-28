package middlewareRepository

import (
	"context"
	"errors"
	"log"
	authPb "tansan/modules/auth/authPb"
	"tansan/pkg/grpccon"
	"time"
)

type (
	MiddlewareRepositoryService interface {
		AccessTokenSearch(pctx context.Context, grpcUrl, accessToken string) error
		RolesCount(pctx context.Context, grpcUrl string,) (int64, error)
	}

	middlewareRepository struct {}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &middlewareRepository{}
}


func (r *middlewareRepository) AccessTokenSearch(pctx context.Context, grpcUrl, accessToken string) error{
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().AccessTokenSearch(ctx, &authPb.AccessTokenSearchReq{AccessToken: accessToken})
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %v", err.Error())
		return errors.New("errors: email or password is incorrect")
	}

	if result == nil {
		log.Printf("Error: access token is invalid")
		return errors.New("errors: access token is invalid")
	}

	if !result.IsValid {
		log.Printf("Error: access token is invalid")
		return errors.New("errors: access token is invalid")
	}

	return nil
}

func (r *middlewareRepository) RolesCount(pctx context.Context, grpcUrl string,) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return -1, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Auth().RolesCount(ctx, &authPb.RolesCountReq{})
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %v", err.Error())
		return -1, errors.New("errors: email or password is incorrect")
	}

	if result == nil {
		log.Printf("Error: role count failed")
		return -1, errors.New("errors: role count failed")
	}

	return result.Count, nil
}