package userbooksHandler

import (
	"context"
	"net/http"
	"tansan/config"
	"tansan/modules/userbooks"
	"tansan/modules/userbooks/userbooksUsecase"
	"tansan/pkg/request"
	"tansan/pkg/response"

	"github.com/labstack/echo/v4"
)

type (
	UserbooksHttpHandlerService interface {
		FindUserBooks(c echo.Context) error
	}

	userbooksHttpHandler struct {
		cfg *config.Config
		userbooksUsecase userbooksUsecase.UserbooksUsecaseService
	}
)

func NewUserbooksHttpHandler(cfg *config.Config, userbooksUsecase userbooksUsecase.UserbooksUsecaseService) UserbooksHttpHandlerService {
	return &userbooksHttpHandler{
		cfg : cfg,
		userbooksUsecase: userbooksUsecase,
	}
}

func (h *userbooksHttpHandler) FindUserBooks(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(userbooks.UserbooksSearchReq)
	userId := c.Param("user_id")

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}


	res, err := h.userbooksUsecase.FindUserBooks(ctx, h.cfg, userId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
