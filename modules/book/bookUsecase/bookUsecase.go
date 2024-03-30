package bookUsecase

import (
	"context"
	"errors"
	"tansan/modules/book"
	bookRepository "tansan/modules/book/bookRepository"
	"tansan/pkg/utils"
)

type (
	BookUsecaseService interface {
		CreateBook(pctx context.Context, req *book.CreateBookReq) (any, error)
	}

	bookUsecase struct {
		bookRepository bookRepository.BookRepositoryService
	}
)

func NewBookUsecase(bookRepository bookRepository.BookRepositoryService) BookUsecaseService {
	return &bookUsecase{bookRepository: bookRepository}
}

func (u *bookUsecase) CreateBook(pctx context.Context, req *book.CreateBookReq) (any, error){
	if !u.bookRepository.IsUniqueBook(pctx, req.Title){
		return nil, errors.New("error: book title already exists")
	}

	bookId, err := u.bookRepository.InsertOneBook(pctx, &book.Book{
		Title: req.Title,
		Price: req.Price,
		Damage: req.Damage,
		UsageStatus: true,
		ImageUrl: req.ImageUrl,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, errors.New("error: inserting book failed")
	}

	return bookId.Hex(), nil
}