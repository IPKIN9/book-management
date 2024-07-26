package entities

import (
	"time"
)

type Book struct {
	BookID        int64
	Title         string
	ISBN          string
	AuthorID      int64
	CategoryID    int64
	PublishedDate *time.Time
	Description   string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
