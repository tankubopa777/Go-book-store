package bookRepository

import (
	"context"
	"errors"
	"log"
	"tansan/modules/book"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BookRepositoryService interface {
		InsertOneBook(pctx context.Context, req *book.Book) (primitive.ObjectID, error)
		IsUniqueBook(pctx context.Context, title string) bool
	}

	bookRepository struct{
		db *mongo.Client
	}
)

func NewBookRepository(db *mongo.Client) BookRepositoryService {
	return &bookRepository{db: db}
}

func (b *bookRepository) bookDbConn(pctx context.Context) *mongo.Database {
	return b.db.Database("book_db")
}

func (r *bookRepository) IsUniqueBook(pctx context.Context, title string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	result := new(book.Book)
	if err := col.FindOne(ctx, bson.M{"title": title},).Decode(result); err != nil {
		log.Printf("Error: IsUniqueBook: %v", err)
		return true
	}
	return false
}


func (r *bookRepository) InsertOneBook(pctx context.Context, req *book.Book) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	bookId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneBook: %v", err.Error())
		return primitive.NilObjectID, errors.New("error: inserting book failed")
	}
	return bookId.InsertedID.(primitive.ObjectID), nil
}