package domain

import (
	"context"
	"time"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]*Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*Category, error)
	GetByName(c context.Context, name string) (*Category, error)
}

type MenuItemRepository interface {
	Create(c context.Context, menuItem *MenuItem) error
	Fetch(c context.Context) ([]*MenuItem, error)
	Update(c context.Context, menuItem *MenuItem) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*MenuItem, error)
	GetByCategory(c context.Context, categoryID int) ([]*MenuItem, error)
	GetByPriceRange(c context.Context, minPrice, maxPrice float64) ([]*MenuItem, error)
	Search(c context.Context, query string) ([]*MenuItem, error)
}

type CategoryUsecase interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]*Category, error)
	Update(c context.Context, category *Category) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*Category, error)
	GetByName(c context.Context, name string) (*Category, error)
}

type MenuItemUsecase interface {
	Create(c context.Context, menuItem *MenuItem) error
	Fetch(c context.Context) ([]*MenuItem, error)
	Update(c context.Context, menuItem *MenuItem) error
	Delete(c context.Context, id int) error
	GetByID(c context.Context, id int) (*MenuItem, error)
	GetByCategory(c context.Context, categoryID int) ([]*MenuItem, error)
	GetByName(c context.Context, name string) (*MenuItem, error)
	GetByPriceRange(c context.Context, minPrice, maxPrice float64) ([]*MenuItem, error)
	Search(c context.Context, query string) ([]*MenuItem, error)
}

/*
Domain:
+----------------------------------------------------------+
|                                                          |
| Controller --> Usecase --> Repository --> DB             |
|                                                          |
+----------------------------------------------------------+
*/
