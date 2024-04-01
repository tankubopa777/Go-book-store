package userbooksRepository

import (
	"context"
	"errors"
	"log"
	"tansan/pkg/grpccon"
	"tansan/pkg/jwtauth"
	"time"

	bookPb "tansan/modules/book/bookPb"
	"tansan/modules/userbooks"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserbooksRepositoryService interface {
		FindBooksInIds(pctx context.Context, grpcUrl string, req *bookPb.FindBooksInIdsReq) (*bookPb.FindBooksInIdsRes, error)
		FindUserBooks(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*userbooks.Userbooks, error)
		CountUserbooks(pctx context.Context, userId string) (int64, error)
	}

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

func (r *userbooksRepository) FindBooksInIds(pctx context.Context, grpcUrl string, req *bookPb.FindBooksInIdsReq) (*bookPb.FindBooksInIdsRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %v", err.Error())
		return nil, errors.New("error: gRPC connection failed in FindBooksInIds")
	}

	result, err := conn.Book().FindBooksInIds(ctx, req)
	if err != nil {
		log.Printf("Error: FindBooksInIds failed: %v", err.Error())
		return nil, errors.New("errors: email or password is incorrect")
	}

	if result == nil {
		log.Printf("Error: FindBooksInIds failed: %v", err.Error())
		return nil, errors.New("errors: book not found")
	}

	if len(result.Books) == 0 {
		log.Printf("Error: FindBooksInIds failed: %v", err.Error())
		return nil, errors.New("errors: book not found")
	}

	return result, nil
}

func (r *userbooksRepository) FindUserBooks(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*userbooks.Userbooks, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	db := r.userbooksDbConn(ctx)
	col := db.Collection("users_userbooks")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindUserBooks failed: %v", err.Error())
		return nil, errors.New("errors: user books not found")
	}

	results := make([]*userbooks.Userbooks, 0)
	for cursors.Next(ctx) {
		result := new(userbooks.Userbooks)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindUserBooks failed: %v", err.Error())
			return nil, errors.New("errors: user books not found")
		}

		results = append(results, result)
	}
	
	return results, nil
}

func (r *userbooksRepository) CountUserbooks(pctx context.Context, userId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.userbooksDbConn(ctx)
	col := db.Collection("users_userbooks")

	count, err := col.CountDocuments(ctx, bson.M{"user_id": userId})
	if err != nil {
		log.Printf("Error: CountuserBooks failed: %v", err.Error())
		return -1, errors.New("error: count Userbooks failed")
	}

	return count, nil
}