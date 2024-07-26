package cmd

import (
	"api-gateway/protos"
	"book-svc/adapters/db"
	"book-svc/repositories"
	"book-svc/usecase"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type server struct {
	protos.UnimplementedBookServiceServer
	bookUsecase usecase.BookUseCase
}

func (s *server) GetBook(ctx context.Context, req *protos.GetBookRequest) (*protos.GetBookResponse, error) {
	book, err := s.bookUsecase.GetBook(ctx, req.BookId)
	if err != nil {
		return nil, err
	}

	return &protos.GetBookResponse{
		Book: &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: book.PublishedDate,
			Description:   book.Description,
			CreatedAt:     book.CreatedAt,
			UpdatedAt:     book.UpdatedAt,
		},
	}, nil
}

func (s *server) GetAllBooks(ctx context.Context, req *protos.GetAllBooksRequest) (*protos.GetAllBooksResponse, error) {
	books, err := s.bookUsecase.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	var protoBooks []*protos.Book
	for _, book := range books {
		protoBooks = append(protoBooks, &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: book.PublishedDate,
			Description:   book.Description,
			CreatedAt:     book.CreatedAt,
			UpdatedAt:     book.UpdatedAt,
		})
	}
	fmt.Println(protoBooks)
	return &protos.GetAllBooksResponse{Books: protoBooks}, nil
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {

		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		postgresUser := os.Getenv("POSTGRES_USER")
		postgresPassword := os.Getenv("POSTGRES_PASSWORD")
		postgresDB := os.Getenv("POSTGRES_DB")
		postgresHost := os.Getenv("POSTGRES_HOST")
		postgresPort := os.Getenv("POSTGRES_PORT")

		dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)

		db, err := db.NewPostgresDB(dataSourceName)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer db.Close()

		bookRepo := repositories.NewBookRepository(db)
		bookUsecase := usecase.NewBookUsecase(bookRepo)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := grpc.NewServer()
		protos.RegisterBookServiceServer(s, &server{bookUsecase: bookUsecase})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	},
}

func Execute() {
	if err := serveCmd.Execute(); err != nil {
		log.Fatalf("Could not start command: %v", err)
	}
}
