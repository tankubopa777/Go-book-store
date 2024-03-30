package bookHandler

import (
	"context"
	"net/http"
	"tansan/config"
	"tansan/modules/book"
	"tansan/modules/book/bookUsecase"
	"tansan/pkg/request"
	"tansan/pkg/response"

	"github.com/labstack/echo/v4"
)

type (
	BookHttpHandlerService interface {
		CreateBook(c echo.Context) error
	}

	bookHttpHandler struct {
		cfg *config.Config
		bookUsecase bookUsecase.BookUsecaseService
	}
)

func NewBookHttpHandler(cfg *config.Config, bookUsecase bookUsecase.BookUsecaseService) BookHttpHandlerService {
	return &bookHttpHandler{
		cfg: cfg,
		bookUsecase: bookUsecase}
}

func (h *bookHttpHandler) CreateBook(c echo.Context) error{
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(book.CreateBookReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	
	res, err := h.bookUsecase.CreateBook(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}