package repositories

import (
	"book-svc/entities"
	"context"
	"database/sql"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookRepository interface {
	GetBook(ctx context.Context, bookId string) (*entities.Book, error)
	GetAllBooks(ctx context.Context) ([]*entities.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*entities.Book, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT book_id, title, isbn, author_id, category_id, published_date, description, created_at, updated_at FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entities.Book
	for rows.Next() {
		var book entities.Book
		var publishedDate, createdAt, updatedAt time.Time
		if err := rows.Scan(
			&book.BookID, &book.Title, &book.ISBN, &book.AuthorID, &book.CategoryID,
			&book.PublishedDate, &book.Description, &book.CreatedAt, &book.UpdatedAt,
		); err != nil {
			return nil, err
		}

		book.PublishedDate = timestamppb.New(publishedDate)
		book.CreatedAt = timestamppb.New(createdAt)
		book.UpdatedAt = timestamppb.New(updatedAt)

		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBook(ctx context.Context, bookId string) (*entities.Book, error) {
	var book entities.Book
	var publishedDate, createdAt, updatedAt time.Time
	query := `SELECT book_id, title, isbn, author_id, category_id, published_date, description, created_at, updated_at
				FROM books WHERE book_id = $1`
	err := r.db.QueryRowContext(ctx, query, bookId).Scan(
		&book.BookID, &book.Title, &book.ISBN, &book.AuthorID, &book.CategoryID,
		&book.PublishedDate, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	book.PublishedDate = timestamppb.New(publishedDate)
	book.CreatedAt = timestamppb.New(createdAt)
	book.UpdatedAt = timestamppb.New(updatedAt)

	log.Printf("Retrieved Book: %+v", book)
	return &book, nil
}
