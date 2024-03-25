package userHandler

import (
	"context"
	"net/http"
	"tansan/config"
	"tansan/modules/user"
	"tansan/modules/user/userUsecase"
	"tansan/pkg/request"
	"tansan/pkg/response"

	"github.com/labstack/echo/v4"
)

type (
	UserHttpHandlerService interface {
		CreateUser(c echo.Context) error
	}

	userHttpHandler struct {
		cfg *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{userUsecase: userUsecase}
}

func (h *userHttpHandler) CreateUser(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.CreateUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)

}