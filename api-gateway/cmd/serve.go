package cmd

import (
	pb "api-gateway/protos"
	"api-gateway/routes"
	"api-gateway/usecases"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// Inisialisasi koneksi gRPC ke book-service
		conn, err := grpc.NewClient(
			"localhost:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		bookClient := pb.NewBookServiceClient(conn)
		bookUsecase := usecases.NewBookUsecase(bookClient)

		r := routes.NewRouter(bookUsecase)

		log.Println("Starting server on :8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	},
}

func Execute() {
	if err := serveCmd.Execute(); err != nil {
		log.Fatalf("Could not start command: %v", err)
	}
}
