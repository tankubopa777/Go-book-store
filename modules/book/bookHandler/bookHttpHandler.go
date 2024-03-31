package bookHandler

import (
	"context"
	"net/http"
	"strings"
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
		FindOneBook(c echo.Context) error
		FindManyBooks(c echo.Context) error
		EditBook(c echo.Context) error
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

func (h *bookHttpHandler) FindOneBook(c echo.Context) error{
	ctx := context.Background()
	
	bookId := strings.TrimPrefix(c.Param("book_id"), "book:")

	res, err := h.bookUsecase.FindOneBook(ctx, bookId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *bookHttpHandler) FindManyBooks(c echo.Context) error{
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(book.BookSearchReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}


	res, err := h.bookUsecase.FindManyBooks(ctx, h.cfg.Paginate.BookNextPageUrl, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *bookHttpHandler) EditBook(c echo.Context) error {
	ctx := context.Background()

	bookId := strings.TrimPrefix(c.Param("book_id"), "book:")

	wrapper := request.ContextWrapper(c)

	req := new(book.BookUpdateReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.bookUsecase.EditBook(ctx, bookId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}