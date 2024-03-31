package bookUsecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"tansan/modules/book"
	bookRepository "tansan/modules/book/bookRepository"
	"tansan/modules/models"
	"tansan/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	BookUsecaseService interface {
		CreateBook(pctx context.Context, req *book.CreateBookReq) (any, error)
		FindOneBook(pctx context.Context, bookId string) (*book.BookShowCase, error)
		FindManyBooks(pctx context.Context, basePaginateUrl string, req *book.BookSearchReq) (*models.PaginateRes, error)
		EditBook(pctx context.Context, bookId string, req *book.BookUpdateReq) (*book.BookShowCase, error) 
	}

	bookUsecase struct {
		bookRepository bookRepository.BookRepositoryService
	}
)

func NewBookUsecase(bookRepository bookRepository.BookRepositoryService) BookUsecaseService {
	return &bookUsecase{bookRepository: bookRepository}
}

func (u *bookUsecase) CreateBook(pctx context.Context, req *book.CreateBookReq) (any, error){
	if !u.bookRepository.IsUniqueBook(pctx, req.Title){
		return nil, errors.New("error: book title already exists")
	}

	bookId, err := u.bookRepository.InsertOneBook(pctx, &book.Book{
		Title: req.Title,
		Price: req.Price,
		Damage: req.Damage,
		UsageStatus: true,
		ImageUrl: req.ImageUrl,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, errors.New("error: inserting book failed")
	}

	return u.FindOneBook(pctx, bookId.Hex())
}

func (u *bookUsecase) FindOneBook(pctx context.Context, bookId string) (*book.BookShowCase, error) {
	result, err := u.bookRepository.FindOneBook(pctx, bookId)
	if err != nil {
		return nil, err
	}
	return &book.BookShowCase{
		BookId: "book:" + result.Id.Hex(),
		Title: result.Title,
		Price: result.Price,
		Damage: result.Damage,
		ImageUrl: result.ImageUrl,
	}, nil
}

func (u *bookUsecase) FindManyBooks(pctx context.Context, basePaginateUrl string, req *book.BookSearchReq) (*models.PaginateRes, error) {
	findBooksFilter := bson.D{}
	findBooksOpts := make([]*options.FindOptions, 0)

	countBooksFilter := bson.D{}

	//Filter
	if req.Start != "" {
		req.Start = strings.TrimPrefix(req.Start, "book:")
		findBooksFilter = append(findBooksFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}
	if req.Title != "" {
		findBooksFilter = append(findBooksFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
		countBooksFilter = append(countBooksFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	}

	findBooksFilter  = append(findBooksFilter , bson.E{"usage_status", true})
	countBooksFilter = append(countBooksFilter, bson.E{"usage_status", true})

	//Options
	findBooksOpts = append(findBooksOpts, options.Find().SetSort(bson.D{{"_id", 1}}))
	findBooksOpts = append(findBooksOpts, options.Find().SetLimit(int64(req.Limit)))
	

	//Find
	results, err := u.bookRepository.FindManyBooks(pctx, findBooksFilter, findBooksOpts)
	if err != nil {
		return nil, err
	}

	//Count
	total, err := u.bookRepository.CountBooks(pctx, countBooksFilter)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &models.PaginateRes{
			Data: make([]*book.BookShowCase, 0),
			Total: int(total),
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
			},
			Next: models.NextPaginate{
				Start: "",
				Href: "",
			},
		}, nil
	}

	return &models.PaginateRes{
		Data: results,
		Total: int(total),
		Limit: req.Limit,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
		},
		Next: models.NextPaginate{
			Start: results[len(results)-1].BookId,
			Href: fmt.Sprintf("%s?limit=%d&title=%s&start=%s", basePaginateUrl, req.Limit, req.Title, results[len(results)-1].BookId),
		},
	}, nil
}

func (u *bookUsecase) EditBook(pctx context.Context, bookId string, req *book.BookUpdateReq) (*book.BookShowCase, error) {
	// Update Logical
	updateReq := bson.M{}
	if req.Title == "" {
		if !u.bookRepository.IsUniqueBook(pctx, req.Title){
			log.Printf("Error: EditBook failed: %v", "error: book title already exists")
			return nil, errors.New("error: book title already exists")
		}

		updateReq["title"] = req.Title
	}
	if req.ImageUrl != "" {
		updateReq["image_url"] = req.ImageUrl
	}
	if req.Price >= 0 {
		updateReq["price"] = req.Price
	}
	if req.Damage > 0 {
		updateReq["damage"] = req.Damage
	}
	updateReq["updated_at"] = utils.LocalTime()

	if err := u.bookRepository.UpdateOneBook(pctx, bookId, updateReq); err != nil {
		log.Printf("Error: EditBook failed: %v", err.Error())
		return nil, errors.New("error: updating book failed")
	}

	return u.FindOneBook(pctx, bookId)
}

