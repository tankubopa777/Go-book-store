package book

import (
	"tansan/modules/models"
)

type (
	CreateBookReq struct {
		Title string `json:"title" validate:"required,max=64"`
		Price float64 `json:"price" validate:"required"`
		Damage int `json:"damage" validate:"required"`
		ImageUrl string `json:"image_url" validate:"required,max=255"`
	}

	BookShowCase struct {
		BookId string `json:"book_id"`
		Title string `json:"title"`
		Price float64 `json:"price"`
		Damage int `json:"damage"`
		ImageUrl string `json:"image_url"`
	}

	BookSearchReq struct {
		Title string `json:"title"`
		models.PaginateReq
	}
	BookUpdateReq struct {
		Title string `json:"title" validate:"required,max=64"`
		Price float64 `json:"price" validate:"required"`
		ImageUrl string `json:"image_url" validate:"required,max=255"`
		Damage int `json:"damage" validate:"required"`
	}

	EnableOrDisableBookReq struct {
		UsageStatus bool `json:"usage_status"`
	}
)