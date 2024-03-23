package payment


type (
	BookServiceReq struct {
		Books []*BookServiceReqDatum `json:"books" validate:"required"`
	}

	BookServiceReqDatum struct {
		BookId string `json:"book_id" validate:"required,max=64"`
		Price float64 `json:"price"`
	}
)