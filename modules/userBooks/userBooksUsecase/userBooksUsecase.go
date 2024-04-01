package userbooksUsecase

import (
	"context"
	"fmt"
	"tansan/config"
	"tansan/modules/book"
	bookPb "tansan/modules/book/bookPb"
	"tansan/modules/models"
	"tansan/modules/userbooks"
	"tansan/modules/userbooks/userbooksRepository"
	"tansan/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserbooksUsecaseService interface {
		FindUserBooks(pcxt context.Context, cfg *config.Config, userId string, req *userbooks.UserbooksSearchReq) (*models.PaginateRes, error)
	}
	
	userbooksUsecase struct {
		userbooksRepository userbooksRepository.UserbooksRepositoryService
	}
)

func NewUserbooksUsecase(userbooksRepository userbooksRepository.UserbooksRepositoryService) UserbooksUsecaseService {
	return &userbooksUsecase{
		userbooksRepository: userbooksRepository,
	}
}

func (u *userbooksUsecase) FindUserBooks(pcxt context.Context, cfg *config.Config, userId string, req *userbooks.UserbooksSearchReq) (*models.PaginateRes, error) {
	// Filter
	filter := bson.D{}

	if req.Start != "" {
		filter = append(filter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}
	filter = append(filter, bson.E{"user_id", userId})
	
	// Options
	opts := make([]*options.FindOptions, 0)

	opts = append(opts, options.Find().SetSort(bson.D{{"_id", 1}}))
	opts = append(opts, options.Find().SetLimit(int64(req.Limit)))

	//Find
	userbooksData, err := u.userbooksRepository.FindUserBooks(pcxt, filter, opts)
	if err != nil {
		return nil, err
	}
	
	bookData, err := u.userbooksRepository.FindBooksInIds(pcxt, cfg.Grpc.BookUrl, &bookPb.FindBooksInIdsReq{
		Ids: func() []string {
			bookIds := make([]string, 0)
			for _, v := range userbooksData{
				bookIds = append(bookIds, v.BookId)
			}
			return bookIds
		}(),
	})

	bookMaps := make(map[string]*book.BookShowCase)
	for _, v := range bookData.Books {
		bookMaps[v.Id] = &book.BookShowCase{
			BookId: v.Id,
			Title: v.Title,
			Price: v.Price,
			ImageUrl: v.ImageUrl,
			Damage: int(v.Damage),
		}
	}

	results := make([]*userbooks.BookInUserbooks, 0)
	for _, v := range userbooksData {
		results = append(results, &userbooks.BookInUserbooks{
			UserbooksId: v.Id.Hex(),
			UserId: v.UserId,
			BookShowCase: &book.BookShowCase{
				BookId: v.BookId,
				Title: bookMaps[v.BookId].Title,
				Price: bookMaps[v.BookId].Price,				
				Damage: bookMaps[v.BookId].Damage,
				ImageUrl: bookMaps[v.BookId].ImageUrl,
			},
		})
	}


	//Count
	total, err := u.userbooksRepository.CountUserbooks(pcxt, userId)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &models.PaginateRes{
			Data: make([]*userbooks.BookInUserbooks, 0),
			Total: int(total),
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.BookNextPageUrl, userId, req.Limit),
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
			Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.BookNextPageUrl, userId, req.Limit),
		},
		Next: models.NextPaginate{
			Start: results[len(results)-1].UserbooksId,
			Href: fmt.Sprintf("%s/%s?limit=%d&start=%s",  cfg.Paginate.BookNextPageUrl, userId, req.Limit, results[len(results)-1].UserbooksId),
		},
			
	}, nil
}