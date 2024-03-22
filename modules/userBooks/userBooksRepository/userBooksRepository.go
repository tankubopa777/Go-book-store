package userBooksRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserBooksRepositoryService interface {}

	userBooksRepository struct {
		db *mongo.Client
	}
)

func NewUserBooksRepository(db *mongo.Client) UserBooksRepositoryService {
	return &userBooksRepository{db}
}

func (r *userBooksRepository) userBooksDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("userBooks_db")
}