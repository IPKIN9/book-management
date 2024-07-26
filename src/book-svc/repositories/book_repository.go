package repositories

import (
	"book-svc/entities"
	"context"
	"database/sql"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookRepository interface {
	GetBook(ctx context.Context, bookId string) (*ProtoBook, error)
	GetAllBooks(ctx context.Context) ([]*ProtoBook, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*ProtoBook, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT book_id, title, isbn, author_id, category_id, published_date, description, created_at, updated_at FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*ProtoBook
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(
			&book.BookID, &book.Title, &book.ISBN, &book.AuthorID, &book.CategoryID,
			&book.PublishedDate, &book.Description, &book.CreatedAt, &book.UpdatedAt,
		); err != nil {
			return nil, err
		}

		protoBook := &ProtoBook{
			BookID:        book.BookID,
			Title:         book.Title,
			ISBN:          book.ISBN,
			AuthorID:      book.AuthorID,
			CategoryID:    book.CategoryID,
			PublishedDate: timestamppb.New(*book.PublishedDate),
			Description:   book.Description,
			CreatedAt:     timestamppb.New(*book.CreatedAt),
			UpdatedAt:     timestamppb.New(*book.UpdatedAt),
		}

		books = append(books, protoBook)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) GetBook(ctx context.Context, bookId string) (*ProtoBook, error) {
	var book entities.Book
	query := `SELECT book_id, title, isbn, author_id, category_id, published_date, description, created_at, updated_at
				FROM books WHERE book_id = $1`
	err := r.db.QueryRowContext(ctx, query, bookId).Scan(
		&book.BookID, &book.Title, &book.ISBN, &book.AuthorID, &book.CategoryID,
		&book.PublishedDate, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	protoBook := &ProtoBook{
		BookID:        book.BookID,
		Title:         book.Title,
		ISBN:          book.ISBN,
		AuthorID:      book.AuthorID,
		CategoryID:    book.CategoryID,
		PublishedDate: timestamppb.New(*book.PublishedDate),
		Description:   book.Description,
		CreatedAt:     timestamppb.New(*book.CreatedAt),
		UpdatedAt:     timestamppb.New(*book.UpdatedAt),
	}

	log.Printf("Retrieved Book: %+v", book)
	return protoBook, nil
}
