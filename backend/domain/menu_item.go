package domain

import (
	"context"
	"time"
)

type MenuItem struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  int       `json:"category_id"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    *Category `json:"category,omitempty"`
	Embedding   []float32 `json:"-"`
}

type MenuItemRepository interface {
	Create(ctx context.Context, menuItem *MenuItem) error
	Fetch(ctx context.Context) ([]*MenuItem, error)
	Update(ctx context.Context, menuItem *MenuItem) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*MenuItem, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*MenuItem, error)
	GetByPriceRange(ctx context.Context, minPrice, maxPrice float64) ([]*MenuItem, error)
}

type MenuItemUsecase interface {
	Create(ctx context.Context, menuItem *MenuItem) error
	Fetch(ctx context.Context) ([]*MenuItem, error)
	Update(ctx context.Context, menuItem *MenuItem) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*MenuItem, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*MenuItem, error)
	GetByPriceRange(ctx context.Context, minPrice, maxPrice float64) ([]*MenuItem, error)
}


