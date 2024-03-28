package middlewareUsecase

import (
	"errors"
	"log"
	"tansan/config"
	"tansan/modules/middleware/middlewareRepository"
	"tansan/pkg/jwtauth"
	"tansan/pkg/rbac"

	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error)
		RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error)
		UserIdParamValidation(c echo.Context) (echo.Context, error)
	}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}

func (u *middlewareUsecase) JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error){
	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return nil, err
	}

	if err := u.middlewareRepository.AccessTokenSearch(ctx, cfg.Grpc.AuthUrl, accessToken); err != nil {
		return nil, err
	}

	c.Set("user_id", claims.UserId)
	c.Set("role_code", claims.RoleCode)

	return c, nil
}

func (u *middlewareUsecase) RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error){
	ctx := c.Request().Context()

	userRoleCode := c.Get("role_code").(int)

	rolesCount, err := u.middlewareRepository.RolesCount(ctx, cfg.Grpc.AuthUrl)
	if err != nil {
		return nil, err
	}

	// LoguserRoleCode
	// log.Printf("userRoleCode: ----->%v", userRoleCode)
	log.Printf("rolesCount: ----->%v", int(rolesCount))


	userRoleBinary := rbac.IntToBinary(userRoleCode, 1)
	// log.Printf("userRoleBinary: %v", userRoleBinary)

	for i := 0; i < len(userRoleBinary); i++ {
    if (userRoleBinary[i] & expected[i]) != 0 {
        return c, nil
    }
}

	return nil, errors.New("errors: permission denied")
}

func (u *middlewareUsecase) UserIdParamValidation(c echo.Context) (echo.Context, error) {
    userIdReq := c.Param("user_id") // สมมติว่าค่านี้มาจาก URL parameter
    userIdTokenValue, ok := c.Get("user_id").(string) // ปรับให้เป็นการตรวจสอบแบบปลอดภัย

    log.Printf("userIdReq: %v", userIdReq)
    log.Printf("userIdToken: %v", userIdTokenValue)

    if userIdReq == "" {
        log.Printf("Error: userId from request is empty")
        return nil, errors.New("errors: userId from request is empty")
    }

    // ถ้า c.Get("user_id") ไม่ใช่ string หรือไม่มีค่า ตัวแปร ok จะเป็น false
    if !ok || userIdTokenValue == "" {
        log.Printf("Error: userId from token is empty or not a string")
        return nil, errors.New("errors: userId from token is empty or not a string")
    }

    if userIdTokenValue != userIdReq {
        log.Printf("Error: userId is not match, user_id_req: %v, user_id_token: %v", userIdReq, userIdTokenValue)
        return nil, errors.New("errors: userId does not match")
    }

    return c, nil
}
