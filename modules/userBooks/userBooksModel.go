package userbooks

import (
	"tansan/modules/book"
	"tansan/modules/models"
)

type (
	UpdateUserbooksReq struct {
		UserId string `json:"user_id" validate:"required,max=64"`
		BookId string `json:"book_id" validate:"required,max=64"`
	}

	BookInUserbooks struct {
		UserbooksId string `json:"user_books_id"`
		*book.BookShowCase
	}

	UserInventory struct {
		UserId string `json:"user_id"`
		*models.PaginateRes `json:"data"`
	}
)