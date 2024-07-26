package usecase

import (
	"book-svc/repositories"
	"context"
)

type BookUseCase interface {
	GetBook(ctx context.Context, bookID string) (*repositories.ProtoBook, error)
	GetAllBooks(ctx context.Context) ([]*repositories.ProtoBook, error)
}

type bookUsecase struct {
	bookRepo repositories.BookRepository
}

func NewBookUsecase(bookRepo repositories.BookRepository) BookUseCase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) GetBook(ctx context.Context, bookID string) (*repositories.ProtoBook, error) {
	return u.bookRepo.GetBook(ctx, bookID)
}

func (u *bookUsecase) GetAllBooks(ctx context.Context) ([]*repositories.ProtoBook, error) {
	return u.bookRepo.GetAllBooks(ctx)
}
