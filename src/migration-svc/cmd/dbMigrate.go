package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var dbMigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations()
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}

func runMigrations() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databases := map[string]string{
		"book":     os.Getenv("BOOK_DB_DSN"),
		"author":   os.Getenv("AUTHOR_DB_DSN"),
		"category": os.Getenv("CATEGORY_DB_DSN"),
		"user":     os.Getenv("USER_DB_DSN"),
	}

	for key, dsn := range databases {
		if dsn == "" {
			log.Fatalf("Missing environment variable for %s database", key)
		}
	}

	for service, dsn := range databases {
		fmt.Printf("Running migrations for %s service...\n", service)
		m, err := migrate.New(
			fmt.Sprintf("file://migrations/%s", service),
			dsn,
		)
		if err != nil {
			fmt.Printf("Migration failed for %s: %v\n", service, err)
			continue
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("Migration failed for %s: %v\n", service, err)
			continue
		}

		fmt.Printf("Migrations ran successfully for %s service\n", service)
	}
}
