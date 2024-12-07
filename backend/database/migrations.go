package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(connString string, migrationPath string) error {
	log.Printf("Starting migrations from path: %s", migrationPath)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Printf("Database connection error: %v", err)
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("Migration driver error: %v", err)
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	sourceURL := fmt.Sprintf("file://%s", migrationPath)
	log.Printf("Migration source URL: %s", sourceURL)

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		log.Printf("Migration instance creation error: %v", err)
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration execution error: %v", err)
		return fmt.Errorf("failed to run up migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

func RunDownMigrations(connString string, migrationPath string) error {
	log.Printf("Starting down migrations from path: %s", migrationPath)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Printf("Database connection error: %v", err)
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("Migration driver error: %v", err)
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	sourceURL := fmt.Sprintf("file://%s", migrationPath)
	log.Printf("Migration source URL: %s", sourceURL)

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		log.Printf("Migration instance creation error: %v", err)
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Down migration execution error: %v", err)
		return fmt.Errorf("failed to run down migrations: %w", err)
	}

	log.Println("Down migrations completed successfully")
	return nil
}

func RunSeeds(connString string, seedPath string) error {
	log.Printf("Running seeds from path: %s", seedPath)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	seedContent, err := os.ReadFile(seedPath)
	if err != nil {
		return fmt.Errorf("failed to read seed file: %w", err)
	}

	_, err = db.Exec(string(seedContent))
	if err != nil {
		return fmt.Errorf("failed to execute seed: %w", err)
	}

	log.Println("Seeds completed successfully")
	return nil
}
