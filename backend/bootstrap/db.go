package bootstrap

import (
	"context"
	"log"
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/database"
)

func NewPostgresDB(env *Env) database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connString := env.DatabaseURL

	log.Printf("Database URL: %s", connString)

	migrationPath := "/app/supabase/migrations"
	if err := database.RunMigrations(connString, migrationPath); err != nil {
		log.Printf("Warning: Migration error: %v", err)
	}

	conn, err := database.NewPostgresDatabase(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return conn
}

func ClosePostgresDBConnection(db database.Database) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}

	log.Println("Database connection pool closed successfully.")

}
