package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pgvector/pgvector-go"
	"github.com/sammig6i/sydneys-sourdough-co/database"
	"github.com/sammig6i/sydneys-sourdough-co/domain"
)

type menuItemRepository struct {
	db database.Database
}

func NewMenuItemRepository(db database.Database) domain.MenuItemRepository {
	return &menuItemRepository{db: db}
}

func (m *menuItemRepository) Create(c context.Context, menuItem *domain.MenuItem) error {
	searchText := menuItem.Name
	if menuItem.Description != "" {
		searchText += " " + menuItem.Description
	}

	_, err := m.db.Exec(c, `
		INSERT INTO menu_items (
			name, description, price, category_id, 
			image_url, embedding, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, NOW(), NOW()
		)
	`, menuItem.Name, menuItem.Description, menuItem.Price,
		menuItem.CategoryID, menuItem.ImageURL,
		pgvector.NewVector(getEmbedding(searchText)))

	if err != nil {
		return fmt.Errorf("error creating menu item: %w", err)
	}

	return nil
}

func (m *menuItemRepository) Fetch(c context.Context) ([]*domain.MenuItem, error) {
	rows, err := m.db.Query(c, `
		SELECT id, name, description, price, category_id, image_url, created_at, updated_at
		FROM menu_items
	`)
	if err != nil {
		return nil, fmt.Errorf("error fetching menu items: %w", err)
	}
	defer rows.Close()

	var menuItems []*domain.MenuItem
	for rows.Next() {
		var mi domain.MenuItem
		err := rows.Scan(
			&mi.ID, &mi.Name, &mi.Description, &mi.Price, &mi.CategoryID, &mi.ImageURL, &mi.CreatedAt, &mi.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning menu item: %w", err)
		}
		menuItems = append(menuItems, &mi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu items: %w", err)
	}

	return menuItems, nil
}

func (m *menuItemRepository) Update(c context.Context, menuItem *domain.MenuItem) error {
	query := "UPDATE menu_items SET "
	var args []interface{}
	argCount := 1
	updatesMade := false

	needsEmbeddingUpdate := menuItem.Name != "" || menuItem.Description != ""

	if menuItem.Name != "" {
		query += fmt.Sprintf("name = $%d, ", argCount)
		args = append(args, menuItem.Name)
		argCount++
		updatesMade = true
	}
	if menuItem.Description != "" {
		query += fmt.Sprintf("description = $%d, ", argCount)
		args = append(args, menuItem.Description)
		argCount++
		updatesMade = true
	}

	if needsEmbeddingUpdate {
		currentItem, err := m.GetByID(c, menuItem.ID)
		if err != nil {
			return fmt.Errorf("error getting current menu item: %w", err)
		}

		searchText := menuItem.Name
		if searchText == "" {
			searchText = currentItem.Name
		}
		if menuItem.Description != "" {
			searchText += " " + menuItem.Description
		} else if currentItem.Description != "" {
			searchText += " " + currentItem.Description
		}

		query += fmt.Sprintf("embedding = $%d, ", argCount)
		args = append(args, pgvector.NewVector(getEmbedding(searchText)))
		argCount++
		updatesMade = true
	}

	if menuItem.Price != 0 {
		query += fmt.Sprintf("price = $%d, ", argCount)
		args = append(args, menuItem.Price)
		argCount++
		updatesMade = true
	}
	if menuItem.CategoryID != 0 {
		query += fmt.Sprintf("category_id = $%d, ", argCount)
		args = append(args, menuItem.CategoryID)
		argCount++
		updatesMade = true
	}
	if menuItem.ImageURL != "" {
		query += fmt.Sprintf("image_url = $%d, ", argCount)
		args = append(args, menuItem.ImageURL)
		argCount++
		updatesMade = true
	}

	if !updatesMade {
		return nil
	}

	query += fmt.Sprintf("updated_at = $%d ", argCount)
	args = append(args, time.Now())
	argCount++

	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf("WHERE id = $%d", argCount)
	args = append(args, menuItem.ID)

	_, err := m.db.Exec(c, query, args...)
	if err != nil {
		return fmt.Errorf("error updating menu item: %w", err)
	}

	return nil
}

func (m *menuItemRepository) Delete(c context.Context, id int) error {
	result, err := m.db.Exec(c, `
		DELETE FROM menu_items
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("error deleting menu item: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("menu item with id %d not found", id)
	}

	return nil
}

func (m *menuItemRepository) GetByID(c context.Context, id int) (*domain.MenuItem, error) {
	row := m.db.QueryRow(c, `
		SELECT id, name, description, price, category_id, image_url, created_at, updated_at
		FROM menu_items
		WHERE id = $1
	`, id)

	var mi domain.MenuItem
	err := row.Scan(
		&mi.ID, &mi.Name, &mi.Description, &mi.Price, &mi.CategoryID, &mi.ImageURL, &mi.CreatedAt, &mi.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("menu item not found")
		}
		return nil, fmt.Errorf("error scanning menu item: %w", err)
	}

	return &mi, nil
}

func (m *menuItemRepository) GetByCategory(c context.Context, categoryID int) ([]*domain.MenuItem, error) {
	rows, err := m.db.Query(c, `
		SELECT id, name, description, price, category_id, image_url, created_at, updated_at
		FROM menu_items
		WHERE category_id = $1
	`, categoryID)

	if err != nil {
		return nil, fmt.Errorf("error fetching menu items by category: %w", err)
	}

	defer rows.Close()

	var menuItems []*domain.MenuItem
	for rows.Next() {
		var mi domain.MenuItem
		err := rows.Scan(
			&mi.ID, &mi.Name, &mi.Description, &mi.Price, &mi.CategoryID, &mi.ImageURL, &mi.CreatedAt, &mi.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning menu item by category: %w", err)
		}
		menuItems = append(menuItems, &mi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu items by category: %w", err)
	}

	return menuItems, nil

}

func (m *menuItemRepository) GetByPriceRange(c context.Context, minPrice, maxPrice float64) ([]*domain.MenuItem, error) {
	rows, err := m.db.Query(c, `
		SELECT id, name, description, price, category_id, image_url, created_at, updated_at
		FROM menu_items
		WHERE price >= $1 AND price <= $2
		ORDER BY price ASC
	`, minPrice, maxPrice)
	if err != nil {
		return nil, fmt.Errorf("error fetching menu items by price range: %w", err)
	}
	defer rows.Close()

	var menuItems []*domain.MenuItem
	for rows.Next() {
		var mi domain.MenuItem
		err := rows.Scan(
			&mi.ID, &mi.Name, &mi.Description, &mi.Price, &mi.CategoryID, &mi.ImageURL, &mi.CreatedAt, &mi.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning menu item by price range: %w", err)
		}
		menuItems = append(menuItems, &mi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu items by price range: %w", err)
	}

	return menuItems, nil
}

func (m *menuItemRepository) Search(c context.Context, query string) ([]*domain.MenuItem, error) {
	rows, err := m.db.Query(c, `
		SELECT 
			id, name, description, price, category_id, 
			image_url, created_at, updated_at
		FROM menu_items
		WHERE embedding IS NOT NULL
		ORDER BY embedding <=> $1
	`, pgvector.NewVector(getEmbedding(query)))

	if err != nil {
		return nil, fmt.Errorf("error searching menu items: %w", err)
	}
	defer rows.Close()

	var menuItems []*domain.MenuItem
	for rows.Next() {
		var mi domain.MenuItem
		err := rows.Scan(
			&mi.ID, &mi.Name, &mi.Description, &mi.Price,
			&mi.CategoryID, &mi.ImageURL, &mi.CreatedAt, &mi.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning menu item: %w", err)
		}
		menuItems = append(menuItems, &mi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating menu items: %w", err)
	}

	return menuItems, nil
}
