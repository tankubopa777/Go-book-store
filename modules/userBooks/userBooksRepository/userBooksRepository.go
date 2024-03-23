package userbooksRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserbooksRepositoryService interface {}

	userbooksRepository struct {
		db *mongo.Client
	}
)

func NewUserbooksRepository(db *mongo.Client) UserbooksRepositoryService {
	return &userbooksRepository{db}
}

func (r *userbooksRepository) userbooksDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("userbooks_db")
}
