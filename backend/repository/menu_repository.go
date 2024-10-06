package repository

import (
	"context"

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
	_, err := m.db.Exec(c, `
		INSERT INTO menu_items (
			name, description, price, category_id, image_url
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`, menuItem.Name, menuItem.Description, menuItem.Price, menuItem.CategoryID, menuItem.ImageURL)

	return err
}

func (m *menuItemRepository) Fetch(c context.Context) ([]*domain.MenuItem, error)

func (m *menuItemRepository) GetByID(c context.Context, id int) (*domain.MenuItem, error)

func (m *menuItemRepository) Update(c context.Context, menuItem *domain.MenuItem) error

func (m *menuItemRepository) Delete(c context.Context, id int) error

func (m *menuItemRepository) GetByCategory(c context.Context, categoryID int) ([]*domain.MenuItem, error)
