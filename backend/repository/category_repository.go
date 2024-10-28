package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pgvector/pgvector-go"
	"github.com/sammig6i/sydneys-sourdough-co/database"
	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type categoryRepository struct {
	db database.Database
}

func NewCategoryRepository(db database.Database) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	searchText := category.Name
	_, err := c.db.Exec(ctx, `
		INSERT INTO categories (name, embedding) 
		VALUES ($1, $2)
		RETURNING id
	`, category.Name, pgvector.NewVector(getEmbedding(searchText)))

	if err != nil {
		return fmt.Errorf("error creating category: %w", err)
	}

	return nil
}

func (c *categoryRepository) Fetch(ctx context.Context) ([]*domain.Category, error) {
	rows, err := c.db.Query(ctx, `
	SELECT id, name
	FROM categories
`)
	if err != nil {
		return nil, fmt.Errorf("error fetching category: %w", err)
	}
	defer rows.Close()

	var categoryItems []*domain.Category
	for rows.Next() {
		var ci domain.Category
		err := rows.Scan(
			&ci.ID, &ci.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning category: %w", err)
		}
		categoryItems = append(categoryItems, &ci)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating category: %w", err)
	}

	return categoryItems, nil

}

func (c *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	query := "UPDATE categories SET "
	var args []interface{}
	argCount := 1
	updatesMade := false

	if category.Name != "" {
		query += fmt.Sprintf("name = $%d, ", argCount)
		args = append(args, category.Name)
		argCount++
		updatesMade = true

		query += fmt.Sprintf("embedding = $%d, ", argCount)
		args = append(args, pgvector.NewVector(getEmbedding(category.Name)))
		argCount++
	}

	if !updatesMade {
		return nil
	}

	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf("WHERE id = $%d", argCount)
	args = append(args, category.ID)

	_, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	return nil

}

func (c *categoryRepository) Delete(ctx context.Context, id int) error {
	result, err := c.db.Exec(ctx, `
		DELETE FROM categories
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("error deleting menu item: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("category with id %d not found", id)
	}

	return nil

}

func (c *categoryRepository) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	row := c.db.QueryRow(ctx, `
		SELECT id, name
		FROM categories
		WHERE id = $1
	`, id)

	var ci domain.Category
	err := row.Scan(
		&ci.ID, &ci.Name,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("error scanning category: %w", err)
	}

	return &ci, nil
}
