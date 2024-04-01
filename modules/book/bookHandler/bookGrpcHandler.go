package bookHandler

import (
	"context"
	bookPb "tansan/modules/book/bookPb"
	"tansan/modules/book/bookUsecase"
)

type (
	bookGrpcHandler struct {
		bookUsecase bookUsecase.BookUsecaseService
		bookPb.UnimplementedBookGrpcServiceServer
	}
)

// mustEmbedUnimplementedBookGrpcServiceServer implements bookPb.BookGrpcServiceServer.
func (*bookGrpcHandler) mustEmbedUnimplementedBookGrpcServiceServer() {
	panic("unimplemented")
}

func NewBookGrpcHandler(bookUsecase bookUsecase.BookUsecaseService) *bookGrpcHandler {
	return &bookGrpcHandler{
		bookUsecase: bookUsecase}
}

func (g *bookGrpcHandler) FindBooksInIds(ctx context.Context, req *bookPb.FindBooksInIdsReq) (*bookPb.FindBooksInIdsRes, error) {
	return g.bookUsecase.FindBooksInIds(ctx, req)
}
