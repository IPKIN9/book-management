package usecases

import (
	entity "api-gateway/entities"
	"context"
)

type BookUseCase interface {
	GetBook(ctx context.Context, id string) (*entity.Book, error)
}

type booUseCase struct {
}
