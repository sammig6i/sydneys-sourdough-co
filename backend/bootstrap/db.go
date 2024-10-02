package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/database"
)

func NewPostgresDatabase(env *Env) database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	if dbUser == "" && dbPass == "" {
		connString = fmt.Sprintf("postgres://%s:%s/%s", dbHost, dbPort, dbName)
	} else {
		connString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	conn, err := database.NewPostgresDatabase(ctx, connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return conn

}

func closePostgresDBConnection(db database.Database) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
