package bookUsecase

import bookRepository "tansan/modules/book/bookRepository"

type (
	BookUsecaseService interface {}

	bookUsecase struct {
		bookRepository bookRepository.BookRepositoryService
	}
)

func NewBookUsecase(bookRepository bookRepository.BookRepositoryService) BookUsecaseService {
	return &bookUsecase{bookRepository: bookRepository}
}