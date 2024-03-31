package bookRepository

import (
	"context"
	"errors"
	"log"
	"tansan/modules/book"
	"tansan/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	BookRepositoryService interface {
		InsertOneBook(pctx context.Context, req *book.Book) (primitive.ObjectID, error)
		IsUniqueBook(pctx context.Context, title string) bool
		FindOneBook(pctx context.Context, bookId string) (*book.Book, error)
		FindManyBooks(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*book.BookShowCase, error)
		CountBooks(pctx context.Context, filter primitive.D) (int64, error)
		UpdateOneBook(pctx context.Context, bookId string, req primitive.M) error
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

func (r *bookRepository) FindOneBook(pctx context.Context, bookId string) (*book.Book, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	result := new(book.Book)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bookId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneBook: %v", err.Error())
		return nil, errors.New("error: item not found")
	}
	return result, nil
}

func (r *bookRepository) FindManyBooks(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*book.BookShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: FindManyBooks failed: %v", err.Error())
		return make([]*book.BookShowCase, 0), errors.New("error: find many books failed")
	}

	results := make([]*book.BookShowCase, 0)
	for cursors.Next(ctx) {
		var result *book.Book
		if err := cursors.Decode(&result); err != nil {
			log.Printf("Error: FindManyBooks decode failed: %v", err.Error())
			return make([]*book.BookShowCase, 0), errors.New("error: find many books failed")
		}
		results = append(results, &book.BookShowCase{
			BookId: "book:" + result.Id.Hex(),
			Title: result.Title,
			Price: result.Price,
			Damage: result.Damage,
			ImageUrl: result.ImageUrl,
		})
	}

	return results, nil
}

func (r *bookRepository) CountBooks(pctx context.Context, filter primitive.D) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error: CountBooks failed: %v", err.Error())
		return -1, errors.New("error: count books failed")
	}

	return count, nil
}

func (r *bookRepository) UpdateOneBook(pctx context.Context, bookId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.bookDbConn(ctx)
	col := db.Collection("books")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(bookId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: UpdateOneBook failed: %v", err.Error())
		return errors.New("error: update book failed")
	}
	log.Printf("UpdateOneBook: %v", result.ModifiedCount)


	return nil
}