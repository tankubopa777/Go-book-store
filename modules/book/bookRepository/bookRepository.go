package bookRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BookRepositoryService interface {}

	bookRepository struct{
		db *mongo.Client
	}
)

func NewBookRepository(db *mongo.Client) BookRepositoryService {
	return &bookRepository{db: db}
}

func (b *bookRepository) itemDbConn(pctx context.Context) *mongo.Database {
	return b.db.Database("book_db")
}