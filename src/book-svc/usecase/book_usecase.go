package usecase

import (
	"book-svc/entities"
	"book-svc/repositories"
	"context"
)

type BookUseCase interface {
	GetBook(ctx context.Context, bookID string) (*entities.Book, error)
	GetAllBooks(ctx context.Context) ([]*entities.Book, error)
}

type bookUsecase struct {
	bookRepo repositories.BookRepository
}

func NewBookUsecase(bookRepo repositories.BookRepository) BookUseCase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) GetBook(ctx context.Context, bookID string) (*entities.Book, error) {
	return u.bookRepo.GetBook(ctx, bookID)
}

func (u *bookUsecase) GetAllBooks(ctx context.Context) ([]*entities.Book, error) {
	return u.bookRepo.GetAllBooks(ctx)
}
