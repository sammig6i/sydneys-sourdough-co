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
	log.Printf("Environment: %s", env.Environment)
	log.Printf("Database URL: %s", connString)

	conn, err := database.NewPostgresDatabase(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return conn
}

func ClosePostgresDBConnection(db database.Database) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}

	log.Println("Database connection pool closed successfully.")

}
