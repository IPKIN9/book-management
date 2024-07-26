package usecases

import (
	entity "api-gateway/entities"
	pb "api-gateway/protos"
	"context"
)

type BookUsecase interface {
	GetBook(ctx context.Context, id string) (*entity.Book, error)
}

type bookUsecase struct {
	bookClient pb.BookServiceClient
}

func NewBookUsecase(bookClient pb.BookServiceClient) BookUsecase {
	return &bookUsecase{bookClient: bookClient}
}

func (u *bookUsecase) GetBook(ctx context.Context, id string) (*entity.Book, error) {
	req := &pb.GetBookRequest{BookId: id}
	res, err := u.bookClient.GetBook(ctx, req)
	if err != nil {
		return nil, err
	}

	return &entity.Book{
		BookID:        res.Book.BookId,
		ISBN:          res.Book.Isbn,
		Title:         res.Book.Title,
		AuthorID:      res.Book.AuthorId,
		CategoryID:    res.Book.CategoryId,
		PublishedDate: res.Book.PublishedDate,
		Description:   res.Book.Description,
		CreatedAt:     res.Book.CreatedAt,
		UpdatedAt:     res.Book.UpdatedAt,
	}, nil
}
